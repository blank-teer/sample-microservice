package cmn

import "reflect"

const (
	KindInvalid       = "invalid"
	KindBool          = "bool"
	KindInt           = "int"
	KindInt8          = "int8"
	KindInt16         = "int16"
	KindInt32         = "int32"
	KindInt64         = "int64"
	KindUint          = "uint"
	KindUint8         = "uint8"
	KindUint16        = "uint16"
	KindUint32        = "uint32"
	KindUint64        = "uint64"
	KindUintptr       = "uintptr"
	KindFloat32       = "float32"
	KindFloat64       = "float64"
	KindComplex64     = "complex64"
	KindComplex128    = "complex128"
	KindArray         = "array"
	KindChan          = "chan"
	KindFunc          = "func"
	KindInterface     = "interface"
	KindMap           = "map"
	KindPtr           = "ptr"
	KindSlice         = "slice"
	KindString        = "string"
	KindStruct        = "struct"
	KindUnsafePointer = "unsafe.Pointer"
)

var kinds = map[string]struct{}{
	KindInvalid:       {},
	KindBool:          {},
	KindInt:           {},
	KindInt8:          {},
	KindInt16:         {},
	KindInt32:         {},
	KindInt64:         {},
	KindUint:          {},
	KindUint8:         {},
	KindUint16:        {},
	KindUint32:        {},
	KindUint64:        {},
	KindUintptr:       {},
	KindFloat32:       {},
	KindFloat64:       {},
	KindComplex64:     {},
	KindComplex128:    {},
	KindArray:         {},
	KindChan:          {},
	KindFunc:          {},
	KindInterface:     {},
	KindMap:           {},
	KindPtr:           {},
	KindSlice:         {},
	KindString:        {},
	KindStruct:        {},
	KindUnsafePointer: {},
}

var kindsPrimitive = map[string]struct{}{
	KindBool:       {},
	KindInt:        {},
	KindInt8:       {},
	KindInt16:      {},
	KindInt32:      {},
	KindInt64:      {},
	KindUint:       {},
	KindUint8:      {},
	KindUint16:     {},
	KindUint32:     {},
	KindUint64:     {},
	KindFloat32:    {},
	KindFloat64:    {},
	KindComplex64:  {},
	KindComplex128: {},
	KindString:     {},
}

var kindsComposite = map[string]struct{}{
	KindInvalid:   {},
	KindUintptr:   {},
	KindArray:     {},
	KindChan:      {},
	KindFunc:      {},
	KindMap:       {},
	KindSlice:     {},
	KindStruct:    {},
	KindInterface: {},
}

var kindsToDereference = map[string]struct{}{
	KindPtr:           {},
	KindUnsafePointer: {},
}

func IsPrimitiveKind(tp reflect.Type) bool {
	_, ok := kindsPrimitive[tp.Kind().String()]
	return ok
}

func IsCompositeKind(tp reflect.Type) bool {
	_, ok := kindsComposite[tp.Kind().String()]
	return ok
}

func IsToDereferenceKind(tp reflect.Type) bool {
	_, ok := kindsToDereference[tp.Kind().String()]
	return ok
}

func IsEqualTypes(value1, value2 interface{}) bool {
	return DeriveType(value1).Name() == DeriveType(value2).Name()
}

func DeriveType(value interface{}) reflect.Type {
	var tp reflect.Type

	if tp2, ok := value.(reflect.Type); ok {
		tp = tp2
	} else {
		tp = reflect.TypeOf(value)
	}

	if tp == nil {
		return tp
	}

	if IsToDereferenceKind(tp) {
		return dereferenceType(tp.Elem())
	}
	return tp
}

func dereferenceType(tp reflect.Type) reflect.Type {
	if IsToDereferenceKind(tp) {
		return dereferenceType(tp.Elem())
	}
	return tp
}
