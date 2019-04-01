package impl

import (
	"github.com/golang/protobuf/v2/internal/encoding/wire"
	piface "github.com/golang/protobuf/v2/runtime/protoiface"
)

type extensionFieldInfo struct {
	wiretag uint64
	tagsize int
	funcs   ifaceCoderFuncs
}

func (mi *MessageType) extensionFieldInfo(desc *piface.ExtensionDescV1) *extensionFieldInfo {
	// As of this time (Go 1.12, linux/amd64), an RWMutex benchmarks as faster
	// than a sync.Map.
	mi.extensionFieldInfosMu.RLock()
	e, ok := mi.extensionFieldInfos[desc.Field]
	mi.extensionFieldInfosMu.RUnlock()
	if ok {
		return e
	}

	etype := extensionTypeFromDesc(desc)
	wiretag := wire.EncodeTag(etype.Number(), wireTypes[etype.Kind()])
	e = &extensionFieldInfo{
		wiretag: wiretag,
		tagsize: wire.SizeVarint(wiretag),
		funcs:   encoderFuncsForValue(etype, etype.GoType()),
	}

	mi.extensionFieldInfosMu.Lock()
	if mi.extensionFieldInfos == nil {
		mi.extensionFieldInfos = make(map[int32]*extensionFieldInfo)
	}
	mi.extensionFieldInfos[desc.Field] = e
	mi.extensionFieldInfosMu.Unlock()
	return e
}
