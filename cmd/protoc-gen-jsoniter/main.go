package main

import (
	"os"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	generatedFilenameExtension = ".jsoniter.go"

	jsoniterPkg  = protogen.GoImportPath("github.com/json-iterator/go")
	generatorPkg = protogen.GoImportPath("github.com/srikrsna/protojsoniter/private/generator")
)

var (
	wktSet = map[protoreflect.FullName]bool{
		(new(structpb.NullValue)).Descriptor().FullName():                  true,
		(&structpb.Struct{}).ProtoReflect().Descriptor().FullName():        true,
		(&structpb.ListValue{}).ProtoReflect().Descriptor().FullName():     true,
		(&structpb.Value{}).ProtoReflect().Descriptor().FullName():         true,
		(&fieldmaskpb.FieldMask{}).ProtoReflect().Descriptor().FullName():  true,
		(&timestamppb.Timestamp{}).ProtoReflect().Descriptor().FullName():  true,
		(&durationpb.Duration{}).ProtoReflect().Descriptor().FullName():    true,
		(&anypb.Any{}).ProtoReflect().Descriptor().FullName():              true,
		(&emptypb.Empty{}).ProtoReflect().Descriptor().FullName():          true,
		(&wrapperspb.BoolValue{}).ProtoReflect().Descriptor().FullName():   true,
		(&wrapperspb.StringValue{}).ProtoReflect().Descriptor().FullName(): true,
		(&wrapperspb.BytesValue{}).ProtoReflect().Descriptor().FullName():  true,
		(&wrapperspb.Int32Value{}).ProtoReflect().Descriptor().FullName():  true,
		(&wrapperspb.Int64Value{}).ProtoReflect().Descriptor().FullName():  true,
		(&wrapperspb.UInt32Value{}).ProtoReflect().Descriptor().FullName(): true,
		(&wrapperspb.UInt64Value{}).ProtoReflect().Descriptor().FullName(): true,
		(&wrapperspb.FloatValue{}).ProtoReflect().Descriptor().FullName():  true,
		(&wrapperspb.DoubleValue{}).ProtoReflect().Descriptor().FullName(): true,
	}
)

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.Files {
			if file.Generate {
				gen(plugin, file)
			}
		}
		return nil
	})
}

func gen(plugin *protogen.Plugin, file *protogen.File) {
	if len(file.Messages) == 0 {
		return
	}
	generatedFile := plugin.NewGeneratedFile(
		file.GeneratedFilenamePrefix+generatedFilenameExtension,
		file.GoImportPath,
	)
	genFileHeader(generatedFile, file)
	for _, message := range file.Messages {
		genMessage(generatedFile, file, message)
	}
}

func genFileHeader(g *protogen.GeneratedFile, file *protogen.File) {
	g.P("// Code generated by ", filepath.Base(os.Args[0]), ". DO NOT EDIT.")
	g.P("//")
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("//", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// Source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
}

func genMessage(g *protogen.GeneratedFile, file *protogen.File, message *protogen.Message) {
	genWriter(g, file, message)
	genReader(g, file, message)
	for _, message := range message.Messages {
		if message.Desc.IsMapEntry() {
			continue
		}
		genMessage(g, file, message)
	}
}

func genWriter(g *protogen.GeneratedFile, file *protogen.File, message *protogen.Message) {
	g.P("// WriteJSON writes the JSON representation to stream")
	g.P("func (x *", message.GoIdent, ") WriteJSON(w *", jsoniterPkg.Ident("Stream"), ") {")
	g.P("if x == nil {")
	g.P("w.WriteNil()")
	g.P("return")
	g.P("}")
	g.P("first := true")
	g.P("w.WriteObjectStart()")
	for _, field := range message.Fields {
		if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
			continue
		}
		genFieldWrite(g, field)
	}
	for _, oneOf := range message.Oneofs {
		if oneOf.Desc.IsSynthetic() {
			continue
		}
		g.P("switch v := x.", oneOf.GoName, ".(type) {")
		for _, field := range oneOf.Fields {
			genWriteOneOfField(g, field)
		}
		g.P("}")
	}
	g.P("w.WriteObjectEnd()")
	g.P("_ = first")
	g.P("}")
}

func genReader(g *protogen.GeneratedFile, file *protogen.File, message *protogen.Message) {
	g.P("// ReadJSON reads the message from the iterator")
	g.P("func (x *", message.GoIdent, ") ReadJSON(r *", jsoniterPkg.Ident("Iterator"), ") {")
	g.P("if x == nil {")
	g.P("return ")
	g.P("}")
	g.P("for f := r.ReadObject(); f != \"\"; f = r.ReadObject() {")
	g.P("switch f {")
	for _, field := range message.Fields {
		g.P("case \"", field.Desc.JSONName(), "\":")
		if field.Oneof == nil || field.Oneof.Desc.IsSynthetic() {
			g.P(generatorPkg.Ident("Read"+lkpFn(field)), "(r, &x.", field.GoName, ")")
		} else {
			g.P("var of ", field.GoIdent)
			g.P(generatorPkg.Ident("Read"+lkpFn(field)), "(r, &of.", field.GoName, ")")
			g.P("x.", field.Oneof.GoName, " = ", "&of")
		}
	}
	g.P("default:")
	g.P("r.Skip()")
	g.P("}")
	g.P("}")
	g.P("}")
}

func genFieldWrite(g *protogen.GeneratedFile, field *protogen.Field) {
	g.P("first = ", generatorPkg.Ident("Write"+lkpFn(field)), "(w, \"", field.Desc.JSONName(), "\", x.", field.GoName, ", first)")
}

func genWriteOneOfField(g *protogen.GeneratedFile, field *protogen.Field) {
	g.P("case *", field.GoIdent, ":")
	g.P("first = ", generatorPkg.Ident("WriteMore"), "(w, first)")
	g.P(generatorPkg.Ident("Write"+lkpFn(field)+"OneOf"), "(w, \"", field.Desc.JSONName(), "\", v.", field.GoName, ")")
}

func lkpFn(field *protogen.Field) (fn string) {
	defer func() {
		if field.Desc.HasOptionalKeyword() {
			fn = "Opt" + fn
		}
		if field.Desc.IsList() {
			fn += "Slice"
		}
	}()
	switch {
	case field.Desc.IsMap():
		return "Map"
	case field.Message != nil:
		if wktSet[field.Message.Desc.FullName()] {
			return field.Message.GoIdent.GoName
		} else {
			return "Jsoniter"
		}
	case field.Enum != nil:
		return "Enum"
	default:
		// Only scalar fields
		return lkpScalarFn(field.Desc.Kind())
	}
}

func lkpScalarFn(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.StringKind:
		return "String"
	case protoreflect.BytesKind:
		return "Bytes"
	case protoreflect.BoolKind:
		return "Bool"
	case protoreflect.DoubleKind:
		return "Float64"
	case protoreflect.FloatKind:
		return "Float32"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "Int32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "Int64"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "Uint32"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "Uint64"
	default:
		// This should never happen
		return "Unknown"
	}
}