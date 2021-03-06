// Code generated by protoc-gen-jsoniter. DO NOT EDIT.
//
// Source: test/v1/test.proto

package testv1

import (
	_go "github.com/json-iterator/go"
	generator "github.com/srikrsna/protojsoniter/private/generator"
)

// WriteJSON writes the JSON representation to stream
func (x *All) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteJsoniter(w, "r", x.R, first)
	first = generator.WriteJsoniter(w, "s", x.S, first)
	first = generator.WriteJsoniter(w, "oF", x.OF, first)
	first = generator.WriteEnum(w, "e", x.E, first)
	first = generator.WriteJsoniter(w, "oW", x.OW, first)
	first = generator.WriteJsoniter(w, "w", x.W, first)
	first = generator.WriteJsoniter(w, "o", x.O, first)
	first = generator.WriteJsoniter(w, "rW", x.RW, first)
	first = generator.WriteJsoniter(w, "n", x.N, first)
	first = generator.WriteJsoniter(w, "m", x.M, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *All) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "r":
			generator.ReadJsoniter(r, &x.R)
		case "s":
			generator.ReadJsoniter(r, &x.S)
		case "oF":
			generator.ReadJsoniter(r, &x.OF)
		case "e":
			generator.ReadEnum(r, &x.E)
		case "oW":
			generator.ReadJsoniter(r, &x.OW)
		case "w":
			generator.ReadJsoniter(r, &x.W)
		case "o":
			generator.ReadJsoniter(r, &x.O)
		case "rW":
			generator.ReadJsoniter(r, &x.RW)
		case "n":
			generator.ReadJsoniter(r, &x.N)
		case "m":
			generator.ReadJsoniter(r, &x.M)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *Repeated) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteStringSlice(w, "s", x.S, first)
	first = generator.WriteInt32Slice(w, "i32", x.I32, first)
	first = generator.WriteInt64Slice(w, "i64", x.I64, first)
	first = generator.WriteUint32Slice(w, "u32", x.U32, first)
	first = generator.WriteUint64Slice(w, "u64", x.U64, first)
	first = generator.WriteFloat32Slice(w, "f32", x.F32, first)
	first = generator.WriteFloat64Slice(w, "f64", x.F64, first)
	first = generator.WriteInt32Slice(w, "si32", x.Si32, first)
	first = generator.WriteInt64Slice(w, "si64", x.Si64, first)
	first = generator.WriteUint32Slice(w, "fi32", x.Fi32, first)
	first = generator.WriteUint64Slice(w, "fi64", x.Fi64, first)
	first = generator.WriteInt32Slice(w, "sfi32", x.Sfi32, first)
	first = generator.WriteInt64Slice(w, "sfi64", x.Sfi64, first)
	first = generator.WriteBoolSlice(w, "bl", x.Bl, first)
	first = generator.WriteBytesSlice(w, "by", x.By, first)
	first = generator.WriteEnumSlice(w, "e", x.E, first)
	first = generator.WriteJsoniterSlice(w, "msg", x.Msg, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *Repeated) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "s":
			generator.ReadStringSlice(r, &x.S)
		case "i32":
			generator.ReadInt32Slice(r, &x.I32)
		case "i64":
			generator.ReadInt64Slice(r, &x.I64)
		case "u32":
			generator.ReadUint32Slice(r, &x.U32)
		case "u64":
			generator.ReadUint64Slice(r, &x.U64)
		case "f32":
			generator.ReadFloat32Slice(r, &x.F32)
		case "f64":
			generator.ReadFloat64Slice(r, &x.F64)
		case "si32":
			generator.ReadInt32Slice(r, &x.Si32)
		case "si64":
			generator.ReadInt64Slice(r, &x.Si64)
		case "fi32":
			generator.ReadUint32Slice(r, &x.Fi32)
		case "fi64":
			generator.ReadUint64Slice(r, &x.Fi64)
		case "sfi32":
			generator.ReadInt32Slice(r, &x.Sfi32)
		case "sfi64":
			generator.ReadInt64Slice(r, &x.Sfi64)
		case "bl":
			generator.ReadBoolSlice(r, &x.Bl)
		case "by":
			generator.ReadBytesSlice(r, &x.By)
		case "e":
			generator.ReadEnumSlice(r, &x.E)
		case "msg":
			generator.ReadJsoniterSlice(r, &x.Msg)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *Optionals) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteOptString(w, "id", x.Id, first)
	first = generator.WriteOptInt32(w, "i32", x.I32, first)
	first = generator.WriteOptInt64(w, "i64", x.I64, first)
	first = generator.WriteOptUint32(w, "u32", x.U32, first)
	first = generator.WriteOptUint64(w, "u64", x.U64, first)
	first = generator.WriteOptFloat32(w, "f32", x.F32, first)
	first = generator.WriteOptFloat64(w, "f64", x.F64, first)
	first = generator.WriteOptInt32(w, "si32", x.Si32, first)
	first = generator.WriteOptInt64(w, "si64", x.Si64, first)
	first = generator.WriteOptUint32(w, "fi32", x.Fi32, first)
	first = generator.WriteOptUint64(w, "fi64", x.Fi64, first)
	first = generator.WriteOptInt32(w, "sfi32", x.Sfi32, first)
	first = generator.WriteOptInt64(w, "sfi64", x.Sfi64, first)
	first = generator.WriteOptBool(w, "bl", x.Bl, first)
	first = generator.WriteOptBytes(w, "by", x.By, first)
	first = generator.WriteOptJsoniter(w, "s", x.S, first)
	first = generator.WriteOptEnum(w, "e", x.E, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *Optionals) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "id":
			generator.ReadOptString(r, &x.Id)
		case "i32":
			generator.ReadOptInt32(r, &x.I32)
		case "i64":
			generator.ReadOptInt64(r, &x.I64)
		case "u32":
			generator.ReadOptUint32(r, &x.U32)
		case "u64":
			generator.ReadOptUint64(r, &x.U64)
		case "f32":
			generator.ReadOptFloat32(r, &x.F32)
		case "f64":
			generator.ReadOptFloat64(r, &x.F64)
		case "si32":
			generator.ReadOptInt32(r, &x.Si32)
		case "si64":
			generator.ReadOptInt64(r, &x.Si64)
		case "fi32":
			generator.ReadOptUint32(r, &x.Fi32)
		case "fi64":
			generator.ReadOptUint64(r, &x.Fi64)
		case "sfi32":
			generator.ReadOptInt32(r, &x.Sfi32)
		case "sfi64":
			generator.ReadOptInt64(r, &x.Sfi64)
		case "bl":
			generator.ReadOptBool(r, &x.Bl)
		case "by":
			generator.ReadOptBytes(r, &x.By)
		case "s":
			generator.ReadOptJsoniter(r, &x.S)
		case "e":
			generator.ReadOptEnum(r, &x.E)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *Message) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteString(w, "id", x.Id, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *Message) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "id":
			generator.ReadString(r, &x.Id)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *WKTs) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteAny(w, "a", x.A, first)
	first = generator.WriteDuration(w, "d", x.D, first)
	first = generator.WriteTimestamp(w, "t", x.T, first)
	first = generator.WriteStruct(w, "st", x.St, first)
	first = generator.WriteInt32Value(w, "i32", x.I32, first)
	first = generator.WriteUInt32Value(w, "ui32", x.Ui32, first)
	first = generator.WriteInt64Value(w, "i64", x.I64, first)
	first = generator.WriteUInt64Value(w, "u64", x.U64, first)
	first = generator.WriteFloatValue(w, "f32", x.F32, first)
	first = generator.WriteDoubleValue(w, "f64", x.F64, first)
	first = generator.WriteBoolValue(w, "b", x.B, first)
	first = generator.WriteStringValue(w, "s", x.S, first)
	first = generator.WriteBytesValue(w, "by", x.By, first)
	first = generator.WriteFieldMask(w, "fm", x.Fm, first)
	first = generator.WriteEmpty(w, "em", x.Em, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *WKTs) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "a":
			generator.ReadAny(r, &x.A)
		case "d":
			generator.ReadDuration(r, &x.D)
		case "t":
			generator.ReadTimestamp(r, &x.T)
		case "st":
			generator.ReadStruct(r, &x.St)
		case "i32":
			generator.ReadInt32Value(r, &x.I32)
		case "ui32":
			generator.ReadUInt32Value(r, &x.Ui32)
		case "i64":
			generator.ReadInt64Value(r, &x.I64)
		case "u64":
			generator.ReadUInt64Value(r, &x.U64)
		case "f32":
			generator.ReadFloatValue(r, &x.F32)
		case "f64":
			generator.ReadDoubleValue(r, &x.F64)
		case "b":
			generator.ReadBoolValue(r, &x.B)
		case "s":
			generator.ReadStringValue(r, &x.S)
		case "by":
			generator.ReadBytesValue(r, &x.By)
		case "fm":
			generator.ReadFieldMask(r, &x.Fm)
		case "em":
			generator.ReadEmpty(r, &x.Em)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *RepeatedWKTs) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteAnySlice(w, "a", x.A, first)
	first = generator.WriteDurationSlice(w, "d", x.D, first)
	first = generator.WriteTimestampSlice(w, "t", x.T, first)
	first = generator.WriteStructSlice(w, "st", x.St, first)
	first = generator.WriteInt32ValueSlice(w, "i32", x.I32, first)
	first = generator.WriteUInt32ValueSlice(w, "ui32", x.Ui32, first)
	first = generator.WriteInt64ValueSlice(w, "i64", x.I64, first)
	first = generator.WriteUInt64ValueSlice(w, "u64", x.U64, first)
	first = generator.WriteFloatValueSlice(w, "f32", x.F32, first)
	first = generator.WriteDoubleValueSlice(w, "f64", x.F64, first)
	first = generator.WriteBoolValueSlice(w, "b", x.B, first)
	first = generator.WriteStringValueSlice(w, "s", x.S, first)
	first = generator.WriteBytesValueSlice(w, "by", x.By, first)
	first = generator.WriteFieldMaskSlice(w, "fm", x.Fm, first)
	first = generator.WriteEmptySlice(w, "em", x.Em, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *RepeatedWKTs) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "a":
			generator.ReadAnySlice(r, &x.A)
		case "d":
			generator.ReadDurationSlice(r, &x.D)
		case "t":
			generator.ReadTimestampSlice(r, &x.T)
		case "st":
			generator.ReadStructSlice(r, &x.St)
		case "i32":
			generator.ReadInt32ValueSlice(r, &x.I32)
		case "ui32":
			generator.ReadUInt32ValueSlice(r, &x.Ui32)
		case "i64":
			generator.ReadInt64ValueSlice(r, &x.I64)
		case "u64":
			generator.ReadUInt64ValueSlice(r, &x.U64)
		case "f32":
			generator.ReadFloatValueSlice(r, &x.F32)
		case "f64":
			generator.ReadDoubleValueSlice(r, &x.F64)
		case "b":
			generator.ReadBoolValueSlice(r, &x.B)
		case "s":
			generator.ReadStringValueSlice(r, &x.S)
		case "by":
			generator.ReadBytesValueSlice(r, &x.By)
		case "fm":
			generator.ReadFieldMaskSlice(r, &x.Fm)
		case "em":
			generator.ReadEmptySlice(r, &x.Em)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *OneOf) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	switch v := x.OneOf.(type) {
	case *OneOf_E:
		first = generator.WriteMore(w, first)
		generator.WriteEnumOneOf(w, "e", v.E)
	case *OneOf_S:
		first = generator.WriteMore(w, first)
		generator.WriteStringOneOf(w, "s", v.S)
	case *OneOf_I32:
		first = generator.WriteMore(w, first)
		generator.WriteInt32OneOf(w, "i32", v.I32)
	case *OneOf_I64:
		first = generator.WriteMore(w, first)
		generator.WriteInt64OneOf(w, "i64", v.I64)
	case *OneOf_U32:
		first = generator.WriteMore(w, first)
		generator.WriteUint32OneOf(w, "u32", v.U32)
	case *OneOf_U64:
		first = generator.WriteMore(w, first)
		generator.WriteUint64OneOf(w, "u64", v.U64)
	case *OneOf_F32:
		first = generator.WriteMore(w, first)
		generator.WriteFloat32OneOf(w, "f32", v.F32)
	case *OneOf_F64:
		first = generator.WriteMore(w, first)
		generator.WriteFloat64OneOf(w, "f64", v.F64)
	case *OneOf_Si32:
		first = generator.WriteMore(w, first)
		generator.WriteInt32OneOf(w, "si32", v.Si32)
	case *OneOf_Si64:
		first = generator.WriteMore(w, first)
		generator.WriteInt64OneOf(w, "si64", v.Si64)
	case *OneOf_Fi32:
		first = generator.WriteMore(w, first)
		generator.WriteUint32OneOf(w, "fi32", v.Fi32)
	case *OneOf_Fi64:
		first = generator.WriteMore(w, first)
		generator.WriteUint64OneOf(w, "fi64", v.Fi64)
	case *OneOf_Sfi32:
		first = generator.WriteMore(w, first)
		generator.WriteInt32OneOf(w, "sfi32", v.Sfi32)
	case *OneOf_Sfi64:
		first = generator.WriteMore(w, first)
		generator.WriteInt64OneOf(w, "sfi64", v.Sfi64)
	case *OneOf_Bl:
		first = generator.WriteMore(w, first)
		generator.WriteBoolOneOf(w, "bl", v.Bl)
	case *OneOf_By:
		first = generator.WriteMore(w, first)
		generator.WriteBytesOneOf(w, "by", v.By)
	case *OneOf_Msg:
		first = generator.WriteMore(w, first)
		generator.WriteJsoniterOneOf(w, "msg", v.Msg)
	}
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *OneOf) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "e":
			var of OneOf_E
			generator.ReadEnum(r, &of.E)
			x.OneOf = &of
		case "s":
			var of OneOf_S
			generator.ReadString(r, &of.S)
			x.OneOf = &of
		case "i32":
			var of OneOf_I32
			generator.ReadInt32(r, &of.I32)
			x.OneOf = &of
		case "i64":
			var of OneOf_I64
			generator.ReadInt64(r, &of.I64)
			x.OneOf = &of
		case "u32":
			var of OneOf_U32
			generator.ReadUint32(r, &of.U32)
			x.OneOf = &of
		case "u64":
			var of OneOf_U64
			generator.ReadUint64(r, &of.U64)
			x.OneOf = &of
		case "f32":
			var of OneOf_F32
			generator.ReadFloat32(r, &of.F32)
			x.OneOf = &of
		case "f64":
			var of OneOf_F64
			generator.ReadFloat64(r, &of.F64)
			x.OneOf = &of
		case "si32":
			var of OneOf_Si32
			generator.ReadInt32(r, &of.Si32)
			x.OneOf = &of
		case "si64":
			var of OneOf_Si64
			generator.ReadInt64(r, &of.Si64)
			x.OneOf = &of
		case "fi32":
			var of OneOf_Fi32
			generator.ReadUint32(r, &of.Fi32)
			x.OneOf = &of
		case "fi64":
			var of OneOf_Fi64
			generator.ReadUint64(r, &of.Fi64)
			x.OneOf = &of
		case "sfi32":
			var of OneOf_Sfi32
			generator.ReadInt32(r, &of.Sfi32)
			x.OneOf = &of
		case "sfi64":
			var of OneOf_Sfi64
			generator.ReadInt64(r, &of.Sfi64)
			x.OneOf = &of
		case "bl":
			var of OneOf_Bl
			generator.ReadBool(r, &of.Bl)
			x.OneOf = &of
		case "by":
			var of OneOf_By
			generator.ReadBytes(r, &of.By)
			x.OneOf = &of
		case "msg":
			var of OneOf_Msg
			generator.ReadJsoniter(r, &of.Msg)
			x.OneOf = &of
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *OneOfWKT) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	switch v := x.OneOf.(type) {
	case *OneOfWKT_A:
		first = generator.WriteMore(w, first)
		generator.WriteAnyOneOf(w, "a", v.A)
	case *OneOfWKT_D:
		first = generator.WriteMore(w, first)
		generator.WriteDurationOneOf(w, "d", v.D)
	case *OneOfWKT_T:
		first = generator.WriteMore(w, first)
		generator.WriteTimestampOneOf(w, "t", v.T)
	case *OneOfWKT_St:
		first = generator.WriteMore(w, first)
		generator.WriteStructOneOf(w, "st", v.St)
	case *OneOfWKT_I32:
		first = generator.WriteMore(w, first)
		generator.WriteInt32ValueOneOf(w, "i32", v.I32)
	case *OneOfWKT_Ui32:
		first = generator.WriteMore(w, first)
		generator.WriteUInt32ValueOneOf(w, "ui32", v.Ui32)
	case *OneOfWKT_I64:
		first = generator.WriteMore(w, first)
		generator.WriteInt64ValueOneOf(w, "i64", v.I64)
	case *OneOfWKT_U64:
		first = generator.WriteMore(w, first)
		generator.WriteUInt64ValueOneOf(w, "u64", v.U64)
	case *OneOfWKT_F32:
		first = generator.WriteMore(w, first)
		generator.WriteFloatValueOneOf(w, "f32", v.F32)
	case *OneOfWKT_F64:
		first = generator.WriteMore(w, first)
		generator.WriteDoubleValueOneOf(w, "f64", v.F64)
	case *OneOfWKT_B:
		first = generator.WriteMore(w, first)
		generator.WriteBoolValueOneOf(w, "b", v.B)
	case *OneOfWKT_S:
		first = generator.WriteMore(w, first)
		generator.WriteStringValueOneOf(w, "s", v.S)
	case *OneOfWKT_By:
		first = generator.WriteMore(w, first)
		generator.WriteBytesValueOneOf(w, "by", v.By)
	case *OneOfWKT_Fm:
		first = generator.WriteMore(w, first)
		generator.WriteFieldMaskOneOf(w, "fm", v.Fm)
	case *OneOfWKT_Em:
		first = generator.WriteMore(w, first)
		generator.WriteEmptyOneOf(w, "em", v.Em)
	}
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *OneOfWKT) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "a":
			var of OneOfWKT_A
			generator.ReadAny(r, &of.A)
			x.OneOf = &of
		case "d":
			var of OneOfWKT_D
			generator.ReadDuration(r, &of.D)
			x.OneOf = &of
		case "t":
			var of OneOfWKT_T
			generator.ReadTimestamp(r, &of.T)
			x.OneOf = &of
		case "st":
			var of OneOfWKT_St
			generator.ReadStruct(r, &of.St)
			x.OneOf = &of
		case "i32":
			var of OneOfWKT_I32
			generator.ReadInt32Value(r, &of.I32)
			x.OneOf = &of
		case "ui32":
			var of OneOfWKT_Ui32
			generator.ReadUInt32Value(r, &of.Ui32)
			x.OneOf = &of
		case "i64":
			var of OneOfWKT_I64
			generator.ReadInt64Value(r, &of.I64)
			x.OneOf = &of
		case "u64":
			var of OneOfWKT_U64
			generator.ReadUInt64Value(r, &of.U64)
			x.OneOf = &of
		case "f32":
			var of OneOfWKT_F32
			generator.ReadFloatValue(r, &of.F32)
			x.OneOf = &of
		case "f64":
			var of OneOfWKT_F64
			generator.ReadDoubleValue(r, &of.F64)
			x.OneOf = &of
		case "b":
			var of OneOfWKT_B
			generator.ReadBoolValue(r, &of.B)
			x.OneOf = &of
		case "s":
			var of OneOfWKT_S
			generator.ReadStringValue(r, &of.S)
			x.OneOf = &of
		case "by":
			var of OneOfWKT_By
			generator.ReadBytesValue(r, &of.By)
			x.OneOf = &of
		case "fm":
			var of OneOfWKT_Fm
			generator.ReadFieldMask(r, &of.Fm)
			x.OneOf = &of
		case "em":
			var of OneOfWKT_Em
			generator.ReadEmpty(r, &of.Em)
			x.OneOf = &of
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *Singular) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteEnum(w, "e", x.E, first)
	first = generator.WriteString(w, "s", x.S, first)
	first = generator.WriteInt32(w, "i32", x.I32, first)
	first = generator.WriteInt64(w, "i64", x.I64, first)
	first = generator.WriteUint32(w, "u32", x.U32, first)
	first = generator.WriteUint64(w, "u64", x.U64, first)
	first = generator.WriteFloat32(w, "f32", x.F32, first)
	first = generator.WriteFloat64(w, "f64", x.F64, first)
	first = generator.WriteInt32(w, "si32", x.Si32, first)
	first = generator.WriteInt64(w, "si64", x.Si64, first)
	first = generator.WriteUint32(w, "fi32", x.Fi32, first)
	first = generator.WriteUint64(w, "fi64", x.Fi64, first)
	first = generator.WriteInt32(w, "sfi32", x.Sfi32, first)
	first = generator.WriteInt64(w, "sfi64", x.Sfi64, first)
	first = generator.WriteBool(w, "bl", x.Bl, first)
	first = generator.WriteBytes(w, "by", x.By, first)
	first = generator.WriteJsoniter(w, "msg", x.Msg, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *Singular) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "e":
			generator.ReadEnum(r, &x.E)
		case "s":
			generator.ReadString(r, &x.S)
		case "i32":
			generator.ReadInt32(r, &x.I32)
		case "i64":
			generator.ReadInt64(r, &x.I64)
		case "u32":
			generator.ReadUint32(r, &x.U32)
		case "u64":
			generator.ReadUint64(r, &x.U64)
		case "f32":
			generator.ReadFloat32(r, &x.F32)
		case "f64":
			generator.ReadFloat64(r, &x.F64)
		case "si32":
			generator.ReadInt32(r, &x.Si32)
		case "si64":
			generator.ReadInt64(r, &x.Si64)
		case "fi32":
			generator.ReadUint32(r, &x.Fi32)
		case "fi64":
			generator.ReadUint64(r, &x.Fi64)
		case "sfi32":
			generator.ReadInt32(r, &x.Sfi32)
		case "sfi64":
			generator.ReadInt64(r, &x.Sfi64)
		case "bl":
			generator.ReadBool(r, &x.Bl)
		case "by":
			generator.ReadBytes(r, &x.By)
		case "msg":
			generator.ReadJsoniter(r, &x.Msg)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *Map) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteMap(w, "en", x.En, first)
	first = generator.WriteMap(w, "msg", x.Msg, first)
	first = generator.WriteMap(w, "str", x.Str, first)
	first = generator.WriteMap(w, "by", x.By, first)
	first = generator.WriteMap(w, "bo", x.Bo, first)
	first = generator.WriteMap(w, "an", x.An, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *Map) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "en":
			generator.ReadMap(r, &x.En)
		case "msg":
			generator.ReadMap(r, &x.Msg)
		case "str":
			generator.ReadMap(r, &x.Str)
		case "by":
			generator.ReadMap(r, &x.By)
		case "bo":
			generator.ReadMap(r, &x.Bo)
		case "an":
			generator.ReadMap(r, &x.An)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *Nested) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteJsoniter(w, "n", x.N, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *Nested) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "n":
			generator.ReadJsoniter(r, &x.N)
		default:
			r.Skip()
		}
	}
}

// WriteJSON writes the JSON representation to stream
func (x *Nested_NestedMessage) WriteJSON(w *_go.Stream) {
	if x == nil {
		w.WriteNil()
		return
	}
	first := true
	w.WriteObjectStart()
	first = generator.WriteEnum(w, "e", x.E, first)
	w.WriteObjectEnd()
	_ = first
}

// ReadJSON reads the message from the iterator
func (x *Nested_NestedMessage) ReadJSON(r *_go.Iterator) {
	if x == nil {
		return
	}
	for f := r.ReadObject(); f != ""; f = r.ReadObject() {
		switch f {
		case "e":
			generator.ReadEnum(r, &x.E)
		default:
			r.Skip()
		}
	}
}
