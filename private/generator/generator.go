// Package generator is used by generated code.
// It should not be used in other libraries as it
// does not guaratee backward compatibility
package generator

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/srikrsna/protojsoniter"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	secondsInNanos       = 999999999
	maxSecondsInDuration = 315576000000
)

func WriteMore(w *jsoniter.Stream, first bool) bool {
	if !first {
		w.WriteMore()
	}
	return false
}

func WriteString(w *jsoniter.Stream, field, value string, first bool) bool {
	if value == "" {
		return first
	}

	if !first {
		w.WriteMore()
	}

	WriteStringOneOf(w, field, value)

	return false
}

func WriteOptString(w *jsoniter.Stream, field string, value *string, first bool) bool {
	if value == nil {
		return first
	}

	if !first {
		w.WriteMore()
	}

	WriteStringOneOf(w, field, *value)

	return false
}

func WriteStringSlice(w *jsoniter.Stream, field string, value []string, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteStringSliceOneOf(w, field, value)
	return false
}

func WriteEnum(
	w *jsoniter.Stream,
	field string,
	value interface {
		fmt.Stringer
		protoreflect.Enum
	}, first bool,
) bool {
	if value.Number() == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteEnumOneOf(w, field, value)
	return false
}

func WriteOptEnum[P *E, E fmt.Stringer](w *jsoniter.Stream, field string, value P, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteEnumOneOf(w, field, *value)
	return false
}

func WriteEnumSlice[E fmt.Stringer](w *jsoniter.Stream, field string, value []E, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteEnumSliceOneOf(w, field, value)
	return false
}

func WriteBool(w *jsoniter.Stream, field string, value, first bool) bool {
	if !value {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBoolOneOf(w, field, value)
	return false
}

func WriteOptBool(w *jsoniter.Stream, field string, value *bool, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBoolOneOf(w, field, *value)
	return false
}

func WriteBoolSlice(w *jsoniter.Stream, field string, value []bool, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBoolSliceOneOf(w, field, value)
	return false
}

func WriteUint32(w *jsoniter.Stream, field string, value uint32, first bool) bool {
	if value == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUint32OneOf(w, field, value)
	return false
}

func WriteOptUint32(w *jsoniter.Stream, field string, value *uint32, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUint32OneOf(w, field, *value)
	return false
}

func WriteUint32Slice(w *jsoniter.Stream, field string, value []uint32, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUint32SliceOneOf(w, field, value)
	return false
}

func WriteUint64(w *jsoniter.Stream, field string, value uint64, first bool) bool {
	if value == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUint64OneOf(w, field, value)
	return false
}

func WriteOptUint64(w *jsoniter.Stream, field string, value *uint64, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUint64OneOf(w, field, *value)
	return false
}

func WriteUint64Slice(w *jsoniter.Stream, field string, value []uint64, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUint64SliceOneOf(w, field, value)
	return false
}

func WriteInt(w *jsoniter.Stream, field string, value int, first bool) bool {
	if value == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteIntOneOf(w, field, value)
	return false
}

func WriteIntSlice(w *jsoniter.Stream, field string, value []int, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteIntSliceOneOf(w, field, value)
	return false
}

func WriteFloat32(w *jsoniter.Stream, field string, value float32, first bool) bool {
	if value == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloat32OneOf(w, field, value)
	return false
}

func WriteOptFloat32(w *jsoniter.Stream, field string, value *float32, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloat32OneOf(w, field, *value)
	return false
}

func WriteFloat32Slice(w *jsoniter.Stream, field string, value []float32, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloat32SliceOneOf(w, field, value)
	return false
}

func WriteFloat64(w *jsoniter.Stream, field string, value float64, first bool) bool {
	if value == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloat64OneOf(w, field, value)
	return false
}

func WriteOptFloat64(w *jsoniter.Stream, field string, value *float64, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloat64OneOf(w, field, *value)
	return false
}

func WriteFloat64Slice(w *jsoniter.Stream, field string, value []float64, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloat64SliceOneOf(w, field, value)
	return false
}

func WriteInt32(w *jsoniter.Stream, field string, value int32, first bool) bool {
	if value == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt32OneOf(w, field, value)
	return false
}

func WriteOptInt32(w *jsoniter.Stream, field string, value *int32, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt32OneOf(w, field, *value)
	return false
}

func WriteInt32Slice(w *jsoniter.Stream, field string, value []int32, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt32SliceOneOf(w, field, value)
	return false
}

func WriteInt64(w *jsoniter.Stream, field string, value int64, first bool) bool {
	if value == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt64OneOf(w, field, value)
	return false
}

func WriteOptInt64(w *jsoniter.Stream, field string, value *int64, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt64OneOf(w, field, *value)
	return false
}

func WriteInt64Slice(w *jsoniter.Stream, field string, value []int64, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt64SliceOneOf(w, field, value)
	return false
}

func WriteFieldMask(w *jsoniter.Stream, field string, value *fieldmaskpb.FieldMask, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFieldMaskOneOf(w, field, value)
	return false
}

func WriteFieldMaskSlice(w *jsoniter.Stream, field string, value []*fieldmaskpb.FieldMask, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFieldMaskSliceOneOf(w, field, value)
	return false
}

func WriteFieldMaskSliceOneOf(w *jsoniter.Stream, field string, value []*fieldmaskpb.FieldMask) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	if value[0] == nil {
		w.WriteNil()
	} else {
		words := make([]string, len(value[0].GetPaths()))
		for i := range words {
			words[i] = Name(value[0].Paths[i]).LowerCamelCase().String()
		}
		w.WriteString(strings.Join(words, ","))
	}
	for _, v := range value[1:] {
		w.WriteMore()
		if v == nil {
			w.WriteNil()
		} else {
			words := make([]string, len(v.GetPaths()))
			for i := range words {
				words[i] = Name(v.Paths[i]).LowerCamelCase().String()
			}
			w.WriteString(strings.Join(words, ","))
		}
	}
	w.WriteArrayEnd()
}

func WriteEmpty(w *jsoniter.Stream, field string, value *emptypb.Empty, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteEmptyOneOf(w, field, value)
	return false
}

func WriteEmptySlice(w *jsoniter.Stream, field string, value []*emptypb.Empty, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteEmptySliceOneOf(w, field, value)
	return false
}

func WriteEmptySliceOneOf(w *jsoniter.Stream, field string, value []*emptypb.Empty) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteEmptyObject()
	for range value[1:] {
		w.WriteMore()
		w.WriteEmptyObject()
	}
	w.WriteArrayEnd()
}

func WriteTimestamp(w *jsoniter.Stream, field string, value *timestamppb.Timestamp, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteTimestampOneOf(w, field, value)
	return false
}

func WriteTimestampSlice(w *jsoniter.Stream, field string, value []*timestamppb.Timestamp, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteTimestampSliceOneOf(w, field, value)
	return false
}

func WriteDuration(w *jsoniter.Stream, field string, value *durationpb.Duration, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteDurationOneOf(w, field, value)
	return false
}

func WriteDurationSlice(w *jsoniter.Stream, field string, value []*durationpb.Duration, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteDurationSliceOneOf(w, field, value)
	return false
}

func WriteStringValue(w *jsoniter.Stream, field string, value *wrapperspb.StringValue, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteStringValueOneOf(w, field, value)
	return false
}

func WriteStringValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.StringValue, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteStringValueSliceOneOf(w, field, value)
	return false
}

func WriteStringValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.StringValue) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}

	w.WriteArrayStart()
	w.WriteString(value[0].GetValue())
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(v.GetValue())
	}
	w.WriteArrayEnd()
}

func WriteBytesValue(w *jsoniter.Stream, field string, value *wrapperspb.BytesValue, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBytesValueOneOf(w, field, value)
	return false
}

func WriteBytesValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.BytesValue, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBytesValueSliceOneOf(w, field, value)
	return false
}

func WriteBytesValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.BytesValue) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(base64.StdEncoding.EncodeToString(value[0].GetValue()))
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(base64.StdEncoding.EncodeToString(v.GetValue()))
	}
	w.WriteArrayEnd()
}

func WriteBoolValue(w *jsoniter.Stream, field string, value *wrapperspb.BoolValue, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBoolValueOneOf(w, field, value)
	return false
}

func WriteBoolValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.BoolValue, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBoolValueSliceOneOf(w, field, value)
	return false
}

func WriteBoolValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.BoolValue) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteBool(value[0].GetValue())
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteBool(v.GetValue())
	}
	w.WriteArrayEnd()
}

func WriteFloatValue(w *jsoniter.Stream, field string, value *wrapperspb.FloatValue, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloatValueOneOf(w, field, value)
	return false
}

func WriteFloatValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.FloatValue, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteFloatValueSliceOneOf(w, field, value)
	return false
}

func WriteFloatValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.FloatValue) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteFloat32(value[0].GetValue())
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteFloat32(v.GetValue())
	}
	w.WriteArrayEnd()
}

func WriteDoubleValue(w *jsoniter.Stream, field string, value *wrapperspb.DoubleValue, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteDoubleValueOneOf(w, field, value)
	return false
}

func WriteDoubleValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.DoubleValue, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteDoubleValueSliceOneOf(w, field, value)
	return false
}

func WriteDoubleValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.DoubleValue) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteFloat64(value[0].GetValue())
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteFloat64(v.GetValue())
	}
	w.WriteArrayEnd()
}

func WriteInt64Value(w *jsoniter.Stream, field string, value *wrapperspb.Int64Value, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt64ValueOneOf(w, field, value)
	return false
}

func WriteInt32Value(w *jsoniter.Stream, field string, value *wrapperspb.Int32Value, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt32ValueOneOf(w, field, value)
	return false
}

func WriteInt32ValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.Int32Value, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt32ValueSliceOneOf(w, field, value)
	return false
}

func WriteInt32ValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.Int32Value) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteInt32(value[0].GetValue())
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteInt32(v.GetValue())
	}
	w.WriteArrayEnd()
}

func WriteUInt32Value(w *jsoniter.Stream, field string, value *wrapperspb.UInt32Value, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUInt32ValueOneOf(w, field, value)
	return false
}

func WriteInt64ValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.Int64Value, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteInt64ValueSliceOneOf(w, field, value)
	return false
}

func WriteInt64ValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.Int64Value) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(strconv.FormatInt(value[0].GetValue(), 10))
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(strconv.FormatInt(v.GetValue(), 10))
	}
	w.WriteArrayEnd()
}

func WriteUInt64Value(w *jsoniter.Stream, field string, value *wrapperspb.UInt64Value, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUInt64ValueOneOf(w, field, value)
	return false
}

func WriteUInt64ValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.UInt64Value, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUInt64ValueSliceOneOf(w, field, value)
	return false
}

func WriteUInt64ValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.UInt64Value) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(strconv.FormatUint(value[0].GetValue(), 10))
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(strconv.FormatUint(v.GetValue(), 10))
	}
	w.WriteArrayEnd()
}

func WriteUInt32ValueSlice(w *jsoniter.Stream, field string, value []*wrapperspb.UInt32Value, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteUInt32ValueSliceOneOf(w, field, value)
	return false
}

func WriteUInt32ValueSliceOneOf(w *jsoniter.Stream, field string, value []*wrapperspb.UInt32Value) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteUint32(value[0].GetValue())
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteUint32(v.GetValue())
	}
	w.WriteArrayEnd()
}

func WriteStruct(w *jsoniter.Stream, field string, value *structpb.Struct, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteStructOneOf(w, field, value)
	return false
}

func WriteStructSlice(w *jsoniter.Stream, field string, value []*structpb.Struct, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteStructSliceOneOf(w, field, value)
	return false
}

func WriteAny(w *jsoniter.Stream, field string, value *anypb.Any, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteAnyOneOf(w, field, value)
	return false
}

func WriteAnySlice(w *jsoniter.Stream, field string, value []*anypb.Any, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteAnySliceOneOf(w, field, value)
	return false
}

func WriteJsoniter[E interface {
	protojsoniter.Writer
	*T
}, T any](w *jsoniter.Stream, field string, value E, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteJsoniterOneOf(w, field, value)
	return false
}

func WriteOptJsoniter[E interface {
	protojsoniter.Writer
	*T
}, T any](w *jsoniter.Stream, field string, value E, first bool) bool {
	return WriteJsoniter(w, field, value, first)
}

func WriteJsoniterSlice[E interface {
	protojsoniter.Writer
	*T
}, T any](w *jsoniter.Stream, field string, value []E, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteJsoniterSliceOneOf(w, field, value)
	return false
}

func WriteBytes(w *jsoniter.Stream, field string, value []byte, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBytesOneOf(w, field, value)
	return false
}

func WriteOptBytes(w *jsoniter.Stream, field string, value []byte, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBytesOneOf(w, field, value)
	return false
}

func WriteBytesSlice(w *jsoniter.Stream, field string, value [][]byte, first bool) bool {
	if len(value) == 0 {
		return first
	}
	if !first {
		w.WriteMore()
	}
	WriteBytesSliceOneOf(w, field, value)
	return false
}

func WriteBytesSliceOneOf(w *jsoniter.Stream, field string, value [][]byte) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(base64.StdEncoding.EncodeToString(value[0]))
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(base64.StdEncoding.EncodeToString(v))
	}
	w.WriteArrayEnd()
}

func WriteMap[K comparable, V any](w *jsoniter.Stream, field string, value map[K]V, first bool) bool {
	if value == nil {
		return first
	}
	if !first {
		w.WriteMore()
	}
	w.WriteObjectField(field)
	w.WriteObjectStart()
	var count int
	for k, v := range value {
		count++
		switch k := any(k).(type) {
		case string:
			w.WriteObjectField(k)
		case int32:
			w.WriteObjectField(strconv.FormatInt(int64(k), 10))
		case int64:
			w.WriteObjectField(strconv.FormatInt(int64(k), 10))
		case uint32:
			w.WriteObjectField(strconv.FormatUint(uint64(k), 10))
		case uint64:
			w.WriteObjectField(strconv.FormatUint(uint64(k), 10))
		case bool:
			w.WriteObjectField(strconv.FormatBool(k))
		default:
			w.Error = fmt.Errorf("protojsoniter: invalid proto map key type: %T", k)
			return false
		}
		switch v := any(v).(type) {
		case protojsoniter.Writer:
			v.WriteJSON(w)
		case protoreflect.Enum:
			enumEntry := v.Descriptor().Values().ByNumber(v.Number())
			if enumEntry == nil {
				continue
			}
			w.WriteString(string(enumEntry.Name()))
		case proto.Message:
			data, err := protojson.Marshal(v)
			if err != nil {
				w.Error = err
				return false
			}
			w.WriteVal(json.RawMessage(data))
		case int32:
			w.WriteInt32(v)
		case int64:
			w.WriteInt64(v)
		case uint32:
			w.WriteUint32(v)
		case uint64:
			w.WriteUint64(v)
		case float32:
			w.WriteFloat32(v)
		case float64:
			w.WriteFloat64(v)
		case bool:
			w.WriteBool(v)
		case string:
			w.WriteString(v)
		case []byte:
			w.WriteString(base64.StdEncoding.EncodeToString(v))
		default:
			w.Error = fmt.Errorf("protojsoniter: invalid proto map value type: %T", v)
			return false
		}
		if count < len(value) {
			w.WriteMore()
		}
	}
	w.WriteObjectEnd()
	return false
}

func WriteStringOneOf(w *jsoniter.Stream, field, value string) {
	w.WriteObjectField(field)
	w.WriteString(value)
}

func WriteStringSliceOneOf(w *jsoniter.Stream, field string, value []string) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(v)
	}
	w.WriteArrayEnd()
}

func WriteEnumOneOf(w *jsoniter.Stream, field string, value fmt.Stringer) {
	w.WriteObjectField(field)
	w.WriteString(value.String())
}

func WriteEnumSliceOneOf[E fmt.Stringer](w *jsoniter.Stream, field string, value []E) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(value[0].String())
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(v.String())
	}
	w.WriteArrayEnd()
}

func WriteBoolOneOf(w *jsoniter.Stream, field string, value bool) {
	w.WriteObjectField(field)
	w.WriteBool(value)
}

func WriteBoolSliceOneOf(w *jsoniter.Stream, field string, value []bool) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteBool(value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteBool(v)
	}
	w.WriteArrayEnd()
}

func WriteUint32OneOf(w *jsoniter.Stream, field string, value uint32) {
	w.WriteObjectField(field)
	w.WriteUint32(value)
}

func WriteUint32SliceOneOf(w *jsoniter.Stream, field string, value []uint32) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteUint32(value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteUint32(v)
	}
	w.WriteArrayEnd()
}

func WriteUint64OneOf(w *jsoniter.Stream, field string, value uint64) {
	w.WriteObjectField(field)
	w.WriteString(strconv.FormatUint(value, 10))
}

func WriteUint64SliceOneOf(w *jsoniter.Stream, field string, value []uint64) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(strconv.FormatUint(value[0], 10))
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(strconv.FormatUint(v, 10))
	}
	w.WriteArrayEnd()
}

func WriteIntOneOf(w *jsoniter.Stream, field string, value int) {
	w.WriteObjectField(field)
	w.WriteInt(value)
}

func WriteIntSliceOneOf(w *jsoniter.Stream, field string, value []int) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteInt(value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteInt(v)
	}
	w.WriteArrayEnd()
}

func WriteFloat32OneOf(w *jsoniter.Stream, field string, value float32) {
	w.WriteObjectField(field)
	w.WriteFloat32(value)
}

func WriteFloat32SliceOneOf(w *jsoniter.Stream, field string, value []float32) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteFloat32(value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteFloat32(v)
	}
	w.WriteArrayEnd()
}

func WriteFloat64OneOf(w *jsoniter.Stream, field string, value float64) {
	w.WriteObjectField(field)
	w.WriteFloat64(value)
}

func WriteFloat64SliceOneOf(w *jsoniter.Stream, field string, value []float64) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteFloat64(value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteFloat64(v)
	}
	w.WriteArrayEnd()
}

func WriteInt32OneOf(w *jsoniter.Stream, field string, value int32) {
	w.WriteObjectField(field)
	w.WriteInt32(value)
}

func WriteInt32SliceOneOf(w *jsoniter.Stream, field string, value []int32) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteInt32(value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteInt32(v)
	}
	w.WriteArrayEnd()
}

func WriteInt64OneOf(w *jsoniter.Stream, field string, value int64) {
	w.WriteObjectField(field)
	w.WriteString(strconv.FormatInt(value, 10))
}

func WriteInt64SliceOneOf(w *jsoniter.Stream, field string, value []int64) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(strconv.FormatInt(value[0], 10))
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(strconv.FormatInt(v, 10))
	}
	w.WriteArrayEnd()
}

func WriteFieldMaskOneOf(w *jsoniter.Stream, field string, value *fieldmaskpb.FieldMask) {
	w.WriteObjectField(field)
	words := make([]string, len(value.GetPaths()))
	for i := range words {
		words[i] = Name(value.Paths[i]).LowerCamelCase().String()
	}
	w.WriteString(strings.Join(words, ","))
}

func WriteEmptyOneOf(w *jsoniter.Stream, field string, value *emptypb.Empty) {
	w.WriteObjectField(field)
	w.WriteEmptyObject()
}

func WriteTimestampOneOf(w *jsoniter.Stream, field string, value *timestamppb.Timestamp) {
	w.WriteObjectField(field)
	w.WriteString(value.AsTime().Format(time.RFC3339Nano))
}

func WriteTimestampSliceOneOf(w *jsoniter.Stream, field string, value []*timestamppb.Timestamp) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	w.WriteString(value[0].AsTime().Format(time.RFC3339Nano))
	for _, v := range value[1:] {
		w.WriteMore()
		w.WriteString(v.AsTime().Format(time.RFC3339Nano))
	}
	w.WriteArrayEnd()
}

func WriteDurationOneOf(w *jsoniter.Stream, field string, value *durationpb.Duration) {
	w.WriteObjectField(field)
	writeDuration(w, value)
}

func writeDuration(w *jsoniter.Stream, d *durationpb.Duration) {
	if v, err := formatDuration(d); err != nil {
		w.Error = err
	} else {
		w.WriteString(v)
	}
}

func formatDuration(d *durationpb.Duration) (string, error) {
	secs, nanos := d.GetSeconds(), d.GetNanos()
	if secs < -maxSecondsInDuration || secs > maxSecondsInDuration {
		return "", fmt.Errorf("duration: seconds out of range %d", secs)
	}
	if nanos < -secondsInNanos || nanos > secondsInNanos {
		return "", fmt.Errorf("duration: nanos out of range %d", nanos)
	}
	if (secs > 0 && nanos < 0) || (secs < 0 && nanos > 0) {
		return "", errors.New("duration: signs of seconds and nanos do not match")
	}
	// Generated output always contains 0, 3, 6, or 9 fractional digits,
	// depending on required precision, followed by the suffix "s".
	var sign string
	if secs < 0 || nanos < 0 {
		sign, secs, nanos = "-", -1*secs, -1*nanos
	}
	x := fmt.Sprintf("%s%d.%09d", sign, secs, nanos)
	x = strings.TrimSuffix(x, "000")
	x = strings.TrimSuffix(x, "000")
	x = strings.TrimSuffix(x, ".000")
	return x + "s", nil
}

func WriteDurationSliceOneOf(w *jsoniter.Stream, field string, value []*durationpb.Duration) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	writeDuration(w, value[0])
	for _, v := range value[1:] {
		w.WriteMore()
		writeDuration(w, v)
	}
	w.WriteArrayEnd()
}

func WriteStringValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.StringValue) {
	w.WriteObjectField(field)
	w.WriteString(value.GetValue())
}

func WriteBytesValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.BytesValue) {
	w.WriteObjectField(field)
	w.WriteString(base64.StdEncoding.EncodeToString(value.GetValue()))
}

func WriteBoolValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.BoolValue) {
	w.WriteObjectField(field)
	w.WriteBool(value.GetValue())
}

func WriteFloatValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.FloatValue) {
	w.WriteObjectField(field)
	w.WriteFloat32(value.GetValue())
}

func WriteDoubleValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.DoubleValue) {
	w.WriteObjectField(field)
	w.WriteFloat64(value.GetValue())
}

func WriteInt64ValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.Int64Value) {
	w.WriteObjectField(field)
	w.WriteString(strconv.FormatInt(value.GetValue(), 10))
}

func WriteInt32ValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.Int32Value) {
	w.WriteObjectField(field)
	w.WriteInt32(value.GetValue())
}

func WriteUInt32ValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.UInt32Value) {
	w.WriteObjectField(field)
	w.WriteUint32(value.GetValue())
}

func WriteUInt64ValueOneOf(w *jsoniter.Stream, field string, value *wrapperspb.UInt64Value) {
	w.WriteObjectField(field)
	w.WriteString(strconv.FormatUint(value.GetValue(), 10))
}

func WriteStructOneOf(w *jsoniter.Stream, field string, value *structpb.Struct) {
	w.WriteObjectField(field)
	buf, err := protojson.Marshal(value)
	if err != nil {
		w.Error = err
		return
	}
	w.WriteVal(json.RawMessage(buf))
}

func WriteStructSliceOneOf(w *jsoniter.Stream, field string, value []*structpb.Struct) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	buf, err := protojson.Marshal(value[0])
	if err != nil {
		w.Error = err
		return
	}
	w.WriteVal(json.RawMessage(buf))
	for _, v := range value[1:] {
		w.WriteMore()
		buf, err := protojson.Marshal(v)
		if err != nil {
			w.Error = err
			return
		}
		w.WriteVal(json.RawMessage(buf))
	}
	w.WriteArrayEnd()
}

func WriteAnyOneOf(w *jsoniter.Stream, field string, value *anypb.Any) {
	w.WriteObjectField(field)
	buf, err := protojson.Marshal(value)
	if err != nil {
		w.Error = err
		return
	}
	w.WriteVal(json.RawMessage(buf))
}

func WriteAnySliceOneOf(w *jsoniter.Stream, field string, value []*anypb.Any) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}

	w.WriteArrayStart()
	buf, err := protojson.Marshal(value[0])
	if err != nil {
		w.Error = err
		return
	}
	w.WriteVal(json.RawMessage(buf))
	for _, v := range value[1:] {
		w.WriteMore()
		buf, err := protojson.Marshal(v)
		if err != nil {
			w.Error = err
			return
		}
		w.WriteVal(json.RawMessage(buf))
	}
	w.WriteArrayEnd()
}

func WriteJsoniterOneOf[E interface {
	protojsoniter.Writer
	*T
}, T any](w *jsoniter.Stream, field string, value E) {
	w.WriteObjectField(field)
	if value == nil {
		w.WriteNil()
		return
	}
	value.WriteJSON(w)
}

func WriteJsoniterSliceOneOf[E interface {
	protojsoniter.Writer
	*T
}, T any](w *jsoniter.Stream, field string, value []E) {
	w.WriteObjectField(field)
	if len(value) == 0 {
		w.WriteEmptyArray()
		return
	}
	w.WriteArrayStart()
	if value[0] == nil {
		w.WriteEmptyObject()
	} else {
		value[0].WriteJSON(w)
	}
	for _, v := range value[1:] {
		w.WriteMore()
		if v == nil {
			w.WriteEmptyObject()
		} else {
			v.WriteJSON(w)
		}
	}
	w.WriteArrayEnd()
}

func WriteBytesOneOf(w *jsoniter.Stream, field string, value []byte) {
	w.WriteObjectField(field)
	w.WriteString(base64.StdEncoding.EncodeToString(value))
}

func ReadMap[K comparable, V any](r *jsoniter.Iterator, m *map[K]V) {
	*m = make(map[K]V)
	r.ReadMapCB(func(r *jsoniter.Iterator, f string) bool {
		var k K
		switch any(k).(type) {
		case string:
			k = any(f).(K)
		case int32:
			i, err := strconv.ParseInt(f, 10, 32)
			if err != nil {
				r.ReportError("unable decode map", err.Error())
				return false
			}
			k = any(int32(i)).(K)
		case int64:
			i, err := strconv.ParseInt(f, 10, 64)
			if err != nil {
				r.ReportError("unable decode map", err.Error())
				return false
			}
			k = any(i).(K)
		case uint32:
			i, err := strconv.ParseUint(f, 10, 32)
			if err != nil {
				r.ReportError("unable decode map", err.Error())
				return false
			}
			k = any(uint32(i)).(K)
		case uint64:
			i, err := strconv.ParseUint(f, 10, 64)
			if err != nil {
				r.ReportError("unable decode map", err.Error())
				return false
			}
			k = any(i).(K)
		case bool:
			i, err := strconv.ParseBool(f)
			if err != nil {
				r.ReportError("unable decode map", err.Error())
				return false
			}
			k = any(i).(K)
		default:
			r.ReportError("unable decode map", fmt.Errorf("protojsoniter: invalid proto map key type: %T", k).Error())
			return false
		}
		var v V
		switch any(v).(type) {
		case protojsoniter.Reader:
			v = reflect.New(reflect.TypeOf(v).Elem()).Interface().(V)
			any(v).(protojsoniter.Reader).ReadJSON(r)
		case protoreflect.Enum:
			e := any(v).(protoreflect.Enum)
			ReadEnum(r, &e)
			v = e.(V)
		case proto.Message:
			var buf json.RawMessage
			r.ReadVal(&buf)
			v = reflect.New(reflect.TypeOf(v).Elem()).Interface().(V)
			if err := protojson.Unmarshal([]byte(buf), any(v).(proto.Message)); err != nil {
				r.ReportError("map decode error", err.Error())
				return false
			}
		case int32:
			v = any(r.ReadInt32()).(V)
		case int64:
			var i int64
			ReadInt64(r, &i)
			v = any(i).(V)
		case uint32:
			v = any(r.ReadUint32()).(V)
		case uint64:
			v = any(r.ReadUint64()).(V)
		case float32:
			v = any(r.ReadFloat32()).(V)
		case float64:
			v = any(r.ReadFloat64()).(V)
		case bool:
			v = any(r.ReadBool()).(V)
		case string:
			v = any(r.ReadString()).(V)
		case []byte:
			var b []byte
			ReadBytes(r, &b)
			v = any(b).(V)
		default:
			r.ReportError("map decode error", fmt.Errorf("protojsoniter: invalid proto map value type: %T", m).Error())
			return false
		}
		(*m)[k] = v
		return true
	})
}

func ReadJsoniter[P interface {
	*E
	protojsoniter.Reader
}, E any](r *jsoniter.Iterator, v *P) {
	if r.ReadNil() {
		return
	}
	var x E
	*v = P(&x)
	(*v).ReadJSON(r)
}

func ReadEnum[E protoreflect.Enum](r *jsoniter.Iterator, v *E) {
	*v = ((*v).Type().New((*v).Descriptor().Values().ByName(protoreflect.Name(r.ReadString())).Number())).(E)
}

func ReadStringSlice(r *jsoniter.Iterator, v *[]string) {
	for r.ReadArray() {
		*v = append(*v, r.ReadString())
	}
}

func ReadInt32Slice(r *jsoniter.Iterator, v *[]int32) {
	for r.ReadArray() {
		*v = append(*v, r.ReadInt32())
	}
}

func ReadInt64Slice(r *jsoniter.Iterator, v *[]int64) {
	for r.ReadArray() {
		if i, err := strconv.ParseInt(r.ReadString(), 10, 64); err != nil {
			r.ReportError("expected int", err.Error())
			return
		} else {
			*v = append(*v, i)
		}
	}
}

func ReadUint32Slice(r *jsoniter.Iterator, v *[]uint32) {
	for r.ReadArray() {
		*v = append(*v, r.ReadUint32())
	}
}

func ReadUint64Slice(r *jsoniter.Iterator, v *[]uint64) {
	for r.ReadArray() {
		if i, err := strconv.ParseUint(r.ReadString(), 10, 64); err != nil {
			r.ReportError("expected int", err.Error())
			return
		} else {
			*v = append(*v, i)
		}
	}
}

func ReadFloat32Slice(r *jsoniter.Iterator, v *[]float32) {
	for r.ReadArray() {
		*v = append(*v, r.ReadFloat32())
	}
}

func ReadFloat64Slice(r *jsoniter.Iterator, v *[]float64) {
	for r.ReadArray() {
		*v = append(*v, r.ReadFloat64())
	}
}

func ReadBoolSlice(r *jsoniter.Iterator, v *[]bool) {
	for r.ReadArray() {
		*v = append(*v, r.ReadBool())
	}
}

func ReadBytesSlice(r *jsoniter.Iterator, v *[][]byte) {
	for r.ReadArray() {
		buf, err := base64.StdEncoding.DecodeString(r.ReadString())
		if err != nil {
			r.Error = err
			return
		}
		*v = append(*v, buf)
	}
}

func ReadEnumSlice[E protoreflect.Enum](r *jsoniter.Iterator, v *[]E) {
	for r.ReadArray() {
		var e E
		ReadEnum(r, &e)
		*v = append(*v, e)
	}
}

func ReadJsoniterSlice[R interface {
	*E
	protojsoniter.Reader
}, E any](r *jsoniter.Iterator, v *[]R) {
	for r.ReadArray() {
		var e R
		ReadJsoniter(r, &e)
		*v = append(*v, e)
	}
}

func ReadOptString(r *jsoniter.Iterator, v **string) {
	if r.ReadNil() {
		return
	}
	s := r.ReadString()
	*v = &s
}

func ReadOptInt32(r *jsoniter.Iterator, v **int32) {
	if r.ReadNil() {
		return
	}
	s := r.ReadInt32()
	*v = &s
}

func ReadOptInt64(r *jsoniter.Iterator, v **int64) {
	if r.ReadNil() {
		return
	}
	var i int64
	ReadInt64(r, &i)
	*v = &i
}

func ReadOptUint32(r *jsoniter.Iterator, v **uint32) {
	if r.ReadNil() {
		return
	}
	s := r.ReadUint32()
	*v = &s
}

func ReadOptUint64(r *jsoniter.Iterator, v **uint64) {
	if r.ReadNil() {
		return
	}
	var s uint64
	ReadUint64(r, &s)
	*v = &s
}

func ReadOptFloat32(r *jsoniter.Iterator, v **float32) {
	if r.ReadNil() {
		return
	}
	s := r.ReadFloat32()
	*v = &s
}

func ReadOptFloat64(r *jsoniter.Iterator, v **float64) {
	if r.ReadNil() {
		return
	}
	s := r.ReadFloat64()
	*v = &s
}

func ReadOptBool(r *jsoniter.Iterator, v **bool) {
	if r.ReadNil() {
		return
	}
	s := r.ReadBool()
	*v = &s
}

func ReadOptBytes(r *jsoniter.Iterator, v *[]byte) {
	if r.ReadNil() {
		return
	}
	ReadBytes(r, v)
}

func ReadOptEnum[E protoreflect.Enum](r *jsoniter.Iterator, v **E) {
	if r.ReadNil() {
		return
	}
	var e E
	ReadEnum(r, &e)
	*v = &e
}

func ReadOptJsoniter[P interface {
	*E
	protojsoniter.Reader
}, E any](r *jsoniter.Iterator, v *P) {
	ReadJsoniter(r, v)
}

func ReadString(r *jsoniter.Iterator, v *string) {
	*v = r.ReadString()
}

func ReadInt32(r *jsoniter.Iterator, v *int32) {
	*v = r.ReadInt32()
}

func ReadInt64(r *jsoniter.Iterator, v *int64) {
	if i, err := strconv.ParseInt(r.ReadString(), 10, 64); err != nil {
		r.ReportError("expected int", err.Error())
		return
	} else {
		*v = i
	}
}

func ReadUint32(r *jsoniter.Iterator, v *uint32) {
	*v = r.ReadUint32()
}

func ReadUint64(r *jsoniter.Iterator, v *uint64) {
	if i, err := strconv.ParseUint(r.ReadString(), 10, 64); err != nil {
		r.ReportError("expected int", err.Error())
		return
	} else {
		*v = i
	}
}

func ReadFloat32(r *jsoniter.Iterator, v *float32) {
	*v = r.ReadFloat32()
}

func ReadFloat64(r *jsoniter.Iterator, v *float64) {
	*v = r.ReadFloat64()
}

func ReadBool(r *jsoniter.Iterator, v *bool) {
	*v = r.ReadBool()
}

func ReadBytes(r *jsoniter.Iterator, v *[]byte) {
	b, err := base64.StdEncoding.DecodeString(r.ReadString())
	if err != nil {
		r.ReportError("decoding byte slice", err.Error())
	}
	*v = b
}

func ReadTimestamp(r *jsoniter.Iterator, v **timestamppb.Timestamp) {
	t, err := time.Parse(time.RFC3339Nano, r.ReadString())
	if err != nil {
		r.ReportError("decoding timestamp", err.Error())
	}
	*v = timestamppb.New(t)
}

func ReadDuration(r *jsoniter.Iterator, v **durationpb.Duration) {
	sec, nanos, ok := parseDuration(r.ReadString())
	if !ok {
		r.ReportError("decode", "invalid duration")
	}
	*v = &durationpb.Duration{
		Seconds: sec,
		Nanos:   nanos,
	}
}

// Copied from protojson
//
// parseDuration parses the given input string for seconds and nanoseconds value
// for the Duration JSON format. The format is a decimal number with a suffix
// 's'. It can have optional plus/minus sign. There needs to be at least an
// integer or fractional part. Fractional part is limited to 9 digits only for
// nanoseconds precision, regardless of whether there are trailing zero digits.
// Example values are 1s, 0.1s, 1.s, .1s, +1s, -1s, -.1s.
func parseDuration(input string) (int64, int32, bool) {
	b := []byte(input)
	size := len(b)
	if size < 2 {
		return 0, 0, false
	}
	if b[size-1] != 's' {
		return 0, 0, false
	}
	b = b[:size-1]

	// Read optional plus/minus symbol.
	var neg bool
	switch b[0] {
	case '-':
		neg = true
		b = b[1:]
	case '+':
		b = b[1:]
	}
	if len(b) == 0 {
		return 0, 0, false
	}

	// Read the integer part.
	var intp []byte
	switch {
	case b[0] == '0':
		b = b[1:]

	case '1' <= b[0] && b[0] <= '9':
		intp = b[0:]
		b = b[1:]
		n := 1
		for len(b) > 0 && '0' <= b[0] && b[0] <= '9' {
			n++
			b = b[1:]
		}
		intp = intp[:n]

	case b[0] == '.':
		// Continue below.

	default:
		return 0, 0, false
	}

	hasFrac := false
	var frac [9]byte
	if len(b) > 0 {
		if b[0] != '.' {
			return 0, 0, false
		}
		// Read the fractional part.
		b = b[1:]
		n := 0
		for len(b) > 0 && n < 9 && '0' <= b[0] && b[0] <= '9' {
			frac[n] = b[0]
			n++
			b = b[1:]
		}
		// It is not valid if there are more bytes left.
		if len(b) > 0 {
			return 0, 0, false
		}
		// Pad fractional part with 0s.
		for i := n; i < 9; i++ {
			frac[i] = '0'
		}
		hasFrac = true
	}

	var secs int64
	if len(intp) > 0 {
		var err error
		secs, err = strconv.ParseInt(string(intp), 10, 64)
		if err != nil {
			return 0, 0, false
		}
	}

	var nanos int64
	if hasFrac {
		nanob := bytes.TrimLeft(frac[:], "0")
		if len(nanob) > 0 {
			var err error
			nanos, err = strconv.ParseInt(string(nanob), 10, 32)
			if err != nil {
				return 0, 0, false
			}
		}
	}

	if neg {
		if secs > 0 {
			secs = -secs
		}
		if nanos > 0 {
			nanos = -nanos
		}
	}
	return secs, int32(nanos), true
}

func ReadInt64Value(r *jsoniter.Iterator, v **wrapperspb.Int64Value) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.Int64Value{}
	ReadInt64(r, &(*v).Value)
}

func ReadInt32Value(r *jsoniter.Iterator, v **wrapperspb.Int32Value) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.Int32Value{Value: r.ReadInt32()}
}

func ReadUInt32Value(r *jsoniter.Iterator, v **wrapperspb.UInt32Value) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.UInt32Value{Value: r.ReadUint32()}
}

func ReadUInt64Value(r *jsoniter.Iterator, v **wrapperspb.UInt64Value) {
	if r.ReadNil() {
		return
	}
	if i, err := strconv.ParseUint(r.ReadString(), 10, 64); err != nil {
		r.ReportError("expected int", err.Error())
		return
	} else {
		*v = &wrapperspb.UInt64Value{Value: i}
	}
}

func ReadFloatValue(r *jsoniter.Iterator, v **wrapperspb.FloatValue) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.FloatValue{Value: r.ReadFloat32()}
}

func ReadDoubleValue(r *jsoniter.Iterator, v **wrapperspb.DoubleValue) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.DoubleValue{Value: r.ReadFloat64()}
}

func ReadBoolValue(r *jsoniter.Iterator, v **wrapperspb.BoolValue) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.BoolValue{Value: r.ReadBool()}
}

func ReadStringValue(r *jsoniter.Iterator, v **wrapperspb.StringValue) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.StringValue{Value: r.ReadString()}
}

func ReadBytesValue(r *jsoniter.Iterator, v **wrapperspb.BytesValue) {
	if r.ReadNil() {
		return
	}
	*v = &wrapperspb.BytesValue{}
	ReadBytes(r, &(*v).Value)
}

func ReadEmpty(r *jsoniter.Iterator, v **emptypb.Empty) {
	if r.ReadNil() {
		return
	}
	for r.ReadObject() != "" {
	}
	*v = &emptypb.Empty{}
}

func ReadAny(r *jsoniter.Iterator, v **anypb.Any) {
	if r.ReadNil() {
		return
	}
	var buf json.RawMessage
	r.ReadVal(&buf)
	var any anypb.Any
	if err := protojson.Unmarshal([]byte(buf), &any); err != nil {
		r.ReportError("decoding any", err.Error())
	}
	*v = &any
}

func ReadStruct(r *jsoniter.Iterator, v **structpb.Struct) {
	if r.ReadNil() {
		return
	}
	var buf json.RawMessage
	r.ReadVal(&buf)
	var st structpb.Struct
	if err := protojson.Unmarshal([]byte(buf), &st); err != nil {
		r.ReportError("decoding struct", err.Error())
	}
	*v = &st
}

func ReadFieldMask(r *jsoniter.Iterator, v **fieldmaskpb.FieldMask) {
	if r.ReadNil() {
		return
	}
	rv := r.ReadString()
	if rv == "" {
		*v = &fieldmaskpb.FieldMask{}
		return
	}
	words := strings.Split(rv, ",")
	for i := range words {
		words[i] = Name(words[i]).LowerSnakeCase().String()
	}
	*v = &fieldmaskpb.FieldMask{
		Paths: words,
	}
}

func ReadAnySlice(r *jsoniter.Iterator, v *[]*anypb.Any) {
	for r.ReadArray() {
		var e *anypb.Any
		ReadAny(r, &e)
		*v = append(*v, e)
	}
}

func ReadDurationSlice(r *jsoniter.Iterator, v *[]*durationpb.Duration) {
	for r.ReadArray() {
		var e *durationpb.Duration
		ReadDuration(r, &e)
		*v = append(*v, e)
	}
}

func ReadTimestampSlice(r *jsoniter.Iterator, v *[]*timestamppb.Timestamp) {
	for r.ReadArray() {
		var e *timestamppb.Timestamp
		ReadTimestamp(r, &e)
		*v = append(*v, e)
	}
}

func ReadStructSlice(r *jsoniter.Iterator, v *[]*structpb.Struct) {
	for r.ReadArray() {
		var e *structpb.Struct
		ReadStruct(r, &e)
		*v = append(*v, e)
	}
}

func ReadInt32ValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.Int32Value) {
	for r.ReadArray() {
		var e *wrapperspb.Int32Value
		ReadInt32Value(r, &e)
		*v = append(*v, e)
	}
}

func ReadInt64ValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.Int64Value) {
	for r.ReadArray() {
		var e *wrapperspb.Int64Value
		ReadInt64Value(r, &e)
		*v = append(*v, e)
	}
}

func ReadUInt32ValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.UInt32Value) {
	for r.ReadArray() {
		var e *wrapperspb.UInt32Value
		ReadUInt32Value(r, &e)
		*v = append(*v, e)
	}
}

func ReadUInt64ValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.UInt64Value) {
	for r.ReadArray() {
		var e *wrapperspb.UInt64Value
		ReadUInt64Value(r, &e)
		*v = append(*v, e)
	}
}

func ReadFloatValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.FloatValue) {
	for r.ReadArray() {
		var e *wrapperspb.FloatValue
		ReadFloatValue(r, &e)
		*v = append(*v, e)
	}
}

func ReadDoubleValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.DoubleValue) {
	for r.ReadArray() {
		var e *wrapperspb.DoubleValue
		ReadDoubleValue(r, &e)
		*v = append(*v, e)
	}
}

func ReadBoolValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.BoolValue) {
	for r.ReadArray() {
		var e *wrapperspb.BoolValue
		ReadBoolValue(r, &e)
		*v = append(*v, e)
	}
}

func ReadStringValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.StringValue) {
	for r.ReadArray() {
		var e *wrapperspb.StringValue
		ReadStringValue(r, &e)
		*v = append(*v, e)
	}
}

func ReadBytesValueSlice(r *jsoniter.Iterator, v *[]*wrapperspb.BytesValue) {
	for r.ReadArray() {
		var e *wrapperspb.BytesValue
		ReadBytesValue(r, &e)
		*v = append(*v, e)
	}
}

func ReadFieldMaskSlice(r *jsoniter.Iterator, v *[]*fieldmaskpb.FieldMask) {
	for r.ReadArray() {
		var e *fieldmaskpb.FieldMask
		ReadFieldMask(r, &e)
		*v = append(*v, e)
	}
}

func ReadEmptySlice(r *jsoniter.Iterator, v *[]*emptypb.Empty) {
	for r.ReadArray() {
		var e *emptypb.Empty
		ReadEmpty(r, &e)
		*v = append(*v, e)
	}
}
