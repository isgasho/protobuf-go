// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal_gengo

import (
	"fmt"
	"math"
	"strings"

	"github.com/golang/protobuf/v2/proto"
	"github.com/golang/protobuf/v2/protogen"
	"github.com/golang/protobuf/v2/reflect/protoreflect"

	descriptorpb "github.com/golang/protobuf/v2/types/descriptor"
)

// TODO: Add support for proto options.

func genReflectFileDescriptor(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo) {
	g.P("var ", f.GoDescriptorIdent, " ", protoreflectPackage.Ident("FileDescriptor"))
	g.P()

	genFileDescriptor(gen, g, f)
	if len(f.allEnums) > 0 {
		g.P("var ", enumTypesVarName(f), " = make([]", protoreflectPackage.Ident("EnumType"), ",", len(f.allEnums), ")")
	}
	if len(f.allMessages) > 0 {
		g.P("var ", messageTypesVarName(f), " = make([]", protoimplPackage.Ident("MessageType"), ",", len(f.allMessages), ")")
	}

	// Generate a unique list of Go types for all declarations and dependencies,
	// and the associated index into the type list for all dependencies.
	var goTypes []string
	var depIdxs []string
	seen := map[protoreflect.FullName]int{}
	genDep := func(name protoreflect.FullName, depSource string) {
		if depSource != "" {
			line := fmt.Sprintf("%d, // %s -> %s", seen[name], depSource, name)
			depIdxs = append(depIdxs, line)
		}
	}
	genEnum := func(e *protogen.Enum, depSource string) {
		if e != nil {
			name := e.Desc.FullName()
			if _, ok := seen[name]; !ok {
				line := fmt.Sprintf("(%s)(0), // %d: %s", g.QualifiedGoIdent(e.GoIdent), len(goTypes), name)
				goTypes = append(goTypes, line)
				seen[name] = len(seen)
			}
			if depSource != "" {
				genDep(name, depSource)
			}
		}
	}
	genMessage := func(m *protogen.Message, depSource string) {
		if m != nil {
			name := m.Desc.FullName()
			if _, ok := seen[name]; !ok {
				line := fmt.Sprintf("(*%s)(nil), // %d: %s", g.QualifiedGoIdent(m.GoIdent), len(goTypes), name)
				if m.Desc.IsMapEntry() {
					// Map entry messages have no associated Go type.
					line = fmt.Sprintf("nil, // %d: %s", len(goTypes), name)
				}
				goTypes = append(goTypes, line)
				seen[name] = len(seen)
			}
			if depSource != "" {
				genDep(name, depSource)
			}
		}
	}

	// This ordering is significant. See protoimpl.FileBuilder.GoTypes.
	for _, enum := range f.allEnums {
		genEnum(enum, "")
	}
	for _, message := range f.allMessages {
		genMessage(message, "")
	}
	for _, extension := range f.allExtensions {
		source := string(extension.Desc.FullName())
		genMessage(extension.ExtendedType, source+":extendee")
	}
	for _, message := range f.allMessages {
		for _, field := range message.Fields {
			if field.Desc.IsWeak() {
				continue
			}
			source := string(field.Desc.FullName())
			genEnum(field.EnumType, source+":type_name")
			genMessage(field.MessageType, source+":type_name")
		}
	}
	for _, extension := range f.allExtensions {
		source := string(extension.Desc.FullName())
		genEnum(extension.EnumType, source+":type_name")
		genMessage(extension.MessageType, source+":type_name")
	}
	for _, service := range f.Services {
		for _, method := range service.Methods {
			source := string(method.Desc.FullName())
			genMessage(method.InputType, source+":input_type")
			genMessage(method.OutputType, source+":output_type")
		}
	}
	if len(depIdxs) > math.MaxInt32 {
		panic("too many dependencies") // sanity check
	}

	g.P("var ", goTypesVarName(f), " = []interface{}{")
	for _, s := range goTypes {
		g.P(s)
	}
	g.P("}")

	g.P("var ", depIdxsVarName(f), " = []int32{")
	for _, s := range depIdxs {
		g.P(s)
	}
	g.P("}")

	g.P("func init() { ", initFuncName(f.File), "() }")

	g.P("func ", initFuncName(f.File), "() {")
	g.P("if ", f.GoDescriptorIdent, " != nil {")
	g.P("return")
	g.P("}")

	// Ensure that initialization functions for different files in the same Go
	// package run in the correct order: Call the init funcs for every .proto file
	// imported by this one that is in the same Go package.
	for i, imps := 0, f.Desc.Imports(); i < imps.Len(); i++ {
		impFile, _ := gen.FileByName(imps.Get(i).Path())
		if impFile.GoImportPath != f.GoImportPath {
			continue
		}
		g.P(initFuncName(impFile), "()")
	}

	if len(f.allExtensions) > 0 {
		g.P("extensionTypes := make([]", protoreflectPackage.Ident("ExtensionType"), ",", len(f.allExtensions), ")")
	}

	g.P(f.GoDescriptorIdent, " = ", protoimplPackage.Ident("FileBuilder"), "{")
	g.P("RawDescriptor: ", rawDescVarName(f), ",")
	g.P("GoTypes: ", goTypesVarName(f), ",")
	g.P("DependencyIndexes: ", depIdxsVarName(f), ",")
	if len(f.allExtensions) > 0 {
		g.P("LegacyExtensions: ", extDecsVarName(f), ",")
	}
	if len(f.allEnums) > 0 {
		g.P("EnumOutputTypes: ", enumTypesVarName(f), ",")
	}
	if len(f.allMessages) > 0 {
		g.P("MessageOutputTypes: ", messageTypesVarName(f), ",")
	}
	if len(f.allExtensions) > 0 {
		g.P("ExtensionOutputTypes: extensionTypes,")
	}
	g.P("FilesRegistry: ", protoregistryPackage.Ident("GlobalFiles"), ",")
	g.P("TypesRegistry: ", protoregistryPackage.Ident("GlobalTypes"), ",")
	g.P("}.Init()")

	// The descriptor proto needs to register the option types with the
	// prototype so that the package can properly handle those option types.
	//
	// TODO: Should this be handled by fileinit at runtime?
	if f.Desc.Path() == "google/protobuf/descriptor.proto" && f.Desc.Package() == "google.protobuf" {
		for _, m := range f.allMessages {
			name := m.GoIdent.GoName
			if strings.HasSuffix(name, "Options") {
				g.P(prototypePackage.Ident("X"), ".Register", name, "((*", name, ")(nil))")
			}
		}
	}

	// Set inputs to nil to allow GC to reclaim resources.
	g.P(rawDescVarName(f), " = nil")
	g.P(goTypesVarName(f), " = nil")
	g.P(depIdxsVarName(f), " = nil")
	g.P("}")
}

func genFileDescriptor(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo) {
	// TODO: Replace this with v2 Clone.
	descProto := new(descriptorpb.FileDescriptorProto)
	b, err := proto.Marshal(f.Proto)
	if err != nil {
		gen.Error(err)
		return
	}
	if err := proto.Unmarshal(b, descProto); err != nil {
		gen.Error(err)
		return
	}

	// Trim the source_code_info from the descriptor.
	descProto.SourceCodeInfo = nil
	b, err = proto.MarshalOptions{Deterministic: true}.Marshal(descProto)
	if err != nil {
		gen.Error(err)
		return
	}

	g.P("var ", rawDescVarName(f), " = []byte{")
	for len(b) > 0 {
		n := 16
		if n > len(b) {
			n = len(b)
		}

		s := ""
		for _, c := range b[:n] {
			s += fmt.Sprintf("0x%02x,", c)
		}
		g.P(s)

		b = b[n:]
	}
	g.P("}")
	g.P()

	onceVar := rawDescVarName(f) + "_once"
	dataVar := rawDescVarName(f) + "_data"
	g.P("var (")
	g.P(onceVar, " ", syncPackage.Ident("Once"))
	g.P(dataVar, " = ", rawDescVarName(f))
	g.P(")")
	g.P()

	g.P("func ", rawDescVarName(f), "GZIP() []byte {")
	g.P(onceVar, ".Do(func() {")
	g.P(dataVar, " = ", protoimplPackage.Ident("X"), ".CompressGZIP(", dataVar, ")")
	g.P("})")
	g.P("return ", dataVar)
	g.P("}")
	g.P()
}

func genReflectEnum(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, enum *protogen.Enum) {
	idx := f.allEnumsByPtr[enum]
	typesVar := enumTypesVarName(f)

	// Type method.
	g.P("func (", enum.GoIdent, ") Type() ", protoreflectPackage.Ident("EnumType"), " {")
	g.P("return ", typesVar, "[", idx, "]")
	g.P("}")
	g.P()

	// Number method.
	g.P("func (x ", enum.GoIdent, ") Number() ", protoreflectPackage.Ident("EnumNumber"), " {")
	g.P("return ", protoreflectPackage.Ident("EnumNumber"), "(x)")
	g.P("}")
	g.P()
}

func genReflectMessage(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, message *protogen.Message) {
	idx := f.allMessagesByPtr[message]
	typesVar := messageTypesVarName(f)

	// ProtoReflect method.
	g.P("func (x *", message.GoIdent, ") ProtoReflect() ", protoreflectPackage.Ident("Message"), " {")
	g.P("return ", typesVar, "[", idx, "].MessageOf(x)")
	g.P("}")
	g.P()
	g.P("func (m *", message.GoIdent, ") XXX_Methods() *", protoifacePackage.Ident("Methods"), " {")
	g.P("return ", typesVar, "[", idx, "].Methods()")
	g.P("}")
}

func rawDescVarName(f *fileInfo) string {
	return "xxx_" + f.GoDescriptorIdent.GoName + "_rawDesc"
}
func goTypesVarName(f *fileInfo) string {
	return "xxx_" + f.GoDescriptorIdent.GoName + "_goTypes"
}
func depIdxsVarName(f *fileInfo) string {
	return "xxx_" + f.GoDescriptorIdent.GoName + "_depIdxs"
}
func enumTypesVarName(f *fileInfo) string {
	return "xxx_" + f.GoDescriptorIdent.GoName + "_enumTypes"
}
func messageTypesVarName(f *fileInfo) string {
	return "xxx_" + f.GoDescriptorIdent.GoName + "_messageTypes"
}
func extDecsVarName(f *fileInfo) string {
	return "xxx_" + f.GoDescriptorIdent.GoName + "_extDescs"
}
func initFuncName(f *protogen.File) string {
	return "xxx_" + f.GoDescriptorIdent.GoName + "_init"
}
