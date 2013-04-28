package protos

import (
	"misc/packet"
	"reflect"
)

func pack(tbl interface{}, writer *packet.Packet) []byte {
	v := reflect.ValueOf(tbl)
	count := v.NumField()

	// code test
	code,ok := Code[reflect.TypeOf(tbl).Name()]
	if ok {
		// write code
		writer.WriteU16(code)
	}

	for i := 0; i < count; i++ {
		f := v.Field(i)
		if (_is_primitive(f)){
			_write_primitive(f, writer)
		} else {
			switch f.Type().Kind() {
			case reflect.Slice, reflect.Array:
				writer.WriteU16(uint16(f.Len()))
				for a:=0; a<f.Len();a++ {
					if (_is_primitive(f.Index(a))) {
						_write_primitive(f.Index(a), writer)
					} else {
						elem := f.Index(a).Interface()
						pack(elem, writer)
					}
				}
			}
		}
	}

	return nil
}

func _is_primitive(f reflect.Value) bool {
	switch f.Type().Kind() {
	case reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.String:
		return true
	}
	return false
}

func _write_primitive(f reflect.Value, writer *packet.Packet) {
	switch f.Type().Kind() {
	case reflect.Uint8:
		writer.WriteByte(f.Interface().(byte))
	case reflect.Uint16:
		writer.WriteU16(f.Interface().(uint16))
	case reflect.Uint32:
		writer.WriteU32(f.Interface().(uint32))
	case reflect.Uint64:
		writer.WriteU64(f.Interface().(uint64))

	case reflect.Int:
		writer.WriteU32(uint32(f.Interface().(int)))
	case reflect.Int8:
		writer.WriteByte(byte(f.Interface().(int8)))
	case reflect.Int16:
		writer.WriteU16(uint16(f.Interface().(int16)))
	case reflect.Int32:
		writer.WriteU32(uint32(f.Interface().(int32)))
	case reflect.Int64:
		writer.WriteU64(uint64(f.Interface().(int64)))

	case reflect.Float32:
		writer.WriteFloat32(f.Interface().(float32))

	case reflect.String:
		writer.WriteString(f.Interface().(string))
	}
}