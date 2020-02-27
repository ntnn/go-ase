package transport

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
)

var (
	endianness = binary.BigEndian
)

// StructToBytes converts a passed struct to a byte slice.
func StructToBytes(a interface{}) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	v := reflect.ValueOf(a)

	for i := 0; i < v.NumField(); i++ {
		field := v.FieldByIndex([]int{i})
		fieldif := field.Interface()

		switch fieldif.(type) {
		case string:
			buf.WriteString(fieldif.(string))
		case byte:
			buf.WriteByte(fieldif.(byte))
		case []byte:
			buf.Write(fieldif.([]byte))
		// TODO redo
		case int, int8, int16, int32, int64, uint, uint16, uint32, uint64:
			binary.Write(buf, endianness, fieldif)
		default:
			return nil, fmt.Errorf("Invalid field type: %T", fieldif)
		}
	}

	return buf, nil
}

func BytesToStruct(bs []byte, a interface{}) (interface{}, error) {
	orig := reflect.ValueOf(a)
	copied := reflect.New(orig.Type()).Elem()

	log.Printf("Fields: %d", copied.NumField())
	for i := 0; i < copied.NumField(); i++ {
		origField := orig.FieldByIndex([]int{i})
		origFieldIf := origField.Interface()

		copiedField := copied.FieldByIndex([]int{i})
		// copiedFieldIf := copiedField.Interface()

		switch origField.Kind() {
		case reflect.Uint8:
			copiedField.Set(reflect.ValueOf(bs[0]))
			bs = bs[1:]
		case reflect.Array:
			length := origField.Len()
			for j := 0; j < length; j++ {
				elem := copiedField.Index(j)
				elem.SetUint(uint64(bs[j]))
			}
			bs = bs[length:]
		default:
			return nil, fmt.Errorf("Invalid field type: %T", origFieldIf)
		}
	}

	log.Printf("%#v", copied)
	log.Printf("%#v", copied.Interface())

	return copied.Interface(), nil
}
