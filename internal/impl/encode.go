// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"sort"
	"sync/atomic"

	"github.com/golang/protobuf/v2/internal/errors"
	proto "github.com/golang/protobuf/v2/proto"
	pref "github.com/golang/protobuf/v2/reflect/protoreflect"
	piface "github.com/golang/protobuf/v2/runtime/protoiface"
)

type marshalOptions uint

const (
	marshalAllowPartial marshalOptions = 1 << iota
	marshalDeterministic
	marshalUseCachedSize
)

func newMarshalOptions(opts piface.MarshalOptions) marshalOptions {
	var o marshalOptions
	if opts.AllowPartial {
		o |= marshalAllowPartial
	}
	if opts.Deterministic {
		o |= marshalDeterministic
	}
	if opts.UseCachedSize {
		o |= marshalUseCachedSize
	}
	return o
}

func (o marshalOptions) Options() proto.MarshalOptions {
	return proto.MarshalOptions{
		AllowPartial:  o.AllowPartial(),
		Deterministic: o.Deterministic(),
		UseCachedSize: o.UseCachedSize(),
	}
}

func (o marshalOptions) AllowPartial() bool  { return o&marshalAllowPartial != 0 }
func (o marshalOptions) Deterministic() bool { return o&marshalDeterministic != 0 }
func (o marshalOptions) UseCachedSize() bool { return o&marshalUseCachedSize != 0 }

// size is protoreflect.Methods.Size.
func (mi *MessageType) size(msg pref.ProtoMessage) (size int) {
	return mi.sizePointer(pointerOfIface(msg), 0)
}

func (mi *MessageType) sizePointer(p pointer, opts marshalOptions) (size int) {
	mi.init()
	if p.IsNil() {
		return 0
	}
	if opts.UseCachedSize() && mi.sizecacheOffset.IsValid() {
		return int(atomic.LoadInt32(p.Apply(mi.sizecacheOffset).Int32()))
	}
	return mi.sizePointerSlow(p, opts)
}

func (mi *MessageType) sizePointerSlow(p pointer, opts marshalOptions) (size int) {
	if mi.extensionOffset.IsValid() {
		e := p.Apply(mi.extensionOffset).Extensions()
		size += mi.sizeExtensions(Export{}.ExtensionFieldsOf(e), opts)
	}
	for _, f := range mi.fieldsOrdered {
		fptr := p.Apply(f.offset)
		if f.isPointer && fptr.Elem().IsNil() {
			continue
		}
		if f.funcs.size == nil {
			continue
		}
		size += f.funcs.size(fptr, f.tagsize, opts)
	}
	if mi.unknownOffset.IsValid() {
		u := *p.Apply(mi.unknownOffset).Bytes()
		size += len(u)
	}
	if mi.sizecacheOffset.IsValid() {
		atomic.StoreInt32(p.Apply(mi.sizecacheOffset).Int32(), int32(size))
	}
	return size
}

// marshalAppend is protoreflect.Methods.MarshalAppend.
func (mi *MessageType) marshalAppend(b []byte, msg pref.ProtoMessage, opts piface.MarshalOptions) ([]byte, error) {
	return mi.marshalAppendPointer(b, pointerOfIface(msg), newMarshalOptions(opts))
}

func (mi *MessageType) marshalAppendPointer(b []byte, p pointer, opts marshalOptions) ([]byte, error) {
	mi.init()
	if p.IsNil() {
		return b, nil
	}
	var err error
	var nerr errors.NonFatal
	// The old marshaler encodes extensions at beginning.
	if mi.extensionOffset.IsValid() {
		e := p.Apply(mi.extensionOffset).Extensions()
		// TODO: Special handling for MessageSet?
		b, err = mi.appendExtensions(b, Export{}.ExtensionFieldsOf(e), opts)
		if !nerr.Merge(err) {
			return b, err
		}
	}
	for _, f := range mi.fieldsOrdered {
		fptr := p.Apply(f.offset)
		if f.isPointer && fptr.Elem().IsNil() {
			continue
		}
		if f.funcs.marshal == nil {
			continue
		}
		b, err = f.funcs.marshal(b, fptr, f.wiretag, opts)
		if !nerr.Merge(err) {
			return b, err
		}
	}
	if mi.unknownOffset.IsValid() {
		u := *p.Apply(mi.unknownOffset).Bytes()
		b = append(b, u...)
	}
	return b, nerr.E
}

func (mi *MessageType) sizeExtensions(ext legacyExtensionFieldsIface, opts marshalOptions) (n int) {
	if !ext.HasInit() {
		return 0
	}
	ext.Lock()
	defer ext.Unlock()
	ext.Range(func(k pref.FieldNumber, e ExtensionFieldV1) bool {
		if e.Value == nil || e.Desc == nil {
			// Extension is only in its encoded form.
			n += len(e.Raw)
			return true
		}
		// Value takes precedence over Raw.
		ei := mi.extensionFieldInfo(e.Desc)
		if ei.funcs.size == nil {
			return true
		}
		n += ei.funcs.size(e.Value, ei.tagsize, opts)
		return true
	})
	return n
}

func (mi *MessageType) appendExtensions(b []byte, ext legacyExtensionFieldsIface, opts marshalOptions) ([]byte, error) {
	if !ext.HasInit() {
		return b, nil
	}
	ext.Lock()
	defer ext.Unlock()

	switch extl := ext.Len(); extl {
	case 0:
		return b, nil
	case 1:
		// Fast-path for one extension: Don't bother sorting the keys.
		var err error
		ext.Range(func(k pref.FieldNumber, e ExtensionFieldV1) bool {
			if e.Value == nil || e.Desc == nil {
				// Extension is only in its encoded form.
				b = append(b, e.Raw...)
				return true
			}
			// Value takes precedence over Raw.
			ei := mi.extensionFieldInfo(e.Desc)
			b, err = ei.funcs.marshal(b, e.Value, ei.wiretag, opts)
			return true
		})
		return b, err
	default:
		// Sort the keys to provide a deterministic encoding.
		// Not sure this is required, but the old code does it.
		keys := make([]int, 0, ext.Len())
		ext.Range(func(k pref.FieldNumber, _ ExtensionFieldV1) bool {
			keys = append(keys, int(k))
			return true
		})
		sort.Ints(keys)
		var err error
		var nerr errors.NonFatal
		for _, k := range keys {
			e := ext.Get(pref.FieldNumber(k))
			if e.Value == nil || e.Desc == nil {
				// Extension is only in its encoded form.
				b = append(b, e.Raw...)
				continue
			}
			// Value takes precedence over Raw.
			ei := mi.extensionFieldInfo(e.Desc)
			b, err = ei.funcs.marshal(b, e.Value, ei.wiretag, opts)
			if !nerr.Merge(err) {
				return b, err
			}
		}
		return b, nerr.E
	}
}
