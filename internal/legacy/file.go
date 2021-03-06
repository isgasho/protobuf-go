// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package legacy

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"

	"github.com/golang/protobuf/v2/proto"
	pref "github.com/golang/protobuf/v2/reflect/protoreflect"
	descriptorpb "github.com/golang/protobuf/v2/types/descriptor"
)

// Every enum and message type generated by protoc-gen-go since commit 2fc053c5
// on February 25th, 2016 has had a method to get the raw descriptor.
// Types that were not generated by protoc-gen-go or were generated prior
// to that version are not supported.
//
// The []byte returned is the encoded form of a FileDescriptorProto message
// compressed using GZIP. The []int is the path from the top-level file
// to the specific message or enum declaration.
type (
	enumV1 interface {
		EnumDescriptor() ([]byte, []int)
	}
	messageV1 interface {
		Descriptor() ([]byte, []int)
	}
)

var fileDescCache sync.Map // map[*byte]*descriptorpb.FileDescriptorProto

// LoadFileDesc unmarshals b as a compressed FileDescriptorProto message.
//
// This assumes that b is immutable and that b does not refer to part of a
// concatenated series of GZIP files (which would require shenanigans that
// rely on the concatenation properties of both protobufs and GZIP).
// File descriptors generated by protoc-gen-go do not rely on that property.
//
// This is exported for testing purposes.
func LoadFileDesc(b []byte) *descriptorpb.FileDescriptorProto {
	// Fast-path: check whether we already have a cached file descriptor.
	if fd, ok := fileDescCache.Load(&b[0]); ok {
		return fd.(*descriptorpb.FileDescriptorProto)
	}

	// Slow-path: decompress and unmarshal the file descriptor proto.
	fd := new(descriptorpb.FileDescriptorProto)
	zr, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	b, err = ioutil.ReadAll(zr)
	if err != nil {
		panic(err)
	}
	err = proto.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(b, fd)
	if err != nil {
		panic(err)
	}
	if fd, ok := fileDescCache.LoadOrStore(&b[0], fd); ok {
		return fd.(*descriptorpb.FileDescriptorProto)
	}
	return fd
}

// parentFileDescriptor returns the parent protoreflect.FileDescriptor for the
// provide descriptor. It returns nil if there is no parent.
func parentFileDescriptor(d pref.Descriptor) pref.FileDescriptor {
	for ok := true; ok; d, ok = d.Parent() {
		if fd, _ := d.(pref.FileDescriptor); fd != nil {
			return fd
		}
	}
	return nil
}
