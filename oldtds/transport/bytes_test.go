package transport

import (
	"reflect"
	"testing"
)

type testStructA struct {
	A, B byte
	C, D [5]byte
}

var cases = map[string]struct {
	struc testStructA
	bytes []byte
}{
	"simple": {
		struc: testStructA{
			A: 0x0,
			B: 0x0,
			C: [5]byte{},
			D: [5]byte{},
		},
		bytes: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
	"1 2 3 4": {
		struc: testStructA{
			A: 0x1,
			B: 0x2,
			C: [5]byte{0x3},
			D: [5]byte{0x4},
		},
		bytes: []byte{0x1, 0x2, 0x3, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0},
	},
	"different lengths": {
		struc: testStructA{
			A: 0x3,
			B: 0x4,
			C: [5]byte{0x1, 0x1, 0x2, 0x2, 0x9},
			D: [5]byte{0x5, 0x6, 0x7},
		},
		bytes: []byte{0x3, 0x4, 0x1, 0x1, 0x2, 0x2, 0x9, 0x5, 0x6, 0x7, 0x0, 0x0},
	},
}

func TestStructToBytes(t *testing.T) {
	for name, cas := range cases {
		t.Run(name,
			func(t *testing.T) {
				recv := StructToBytes(cas.struc)
				if !reflect.DeepEqual(cas.bytes, recv) {
					t.Errorf("Received unexpected result:")
					t.Errorf("Expected: %#v", cas.bytes)
					t.Errorf("Received: %#v", recv)
				}
			},
		)
	}
}

func TestBytesToStruct(t *testing.T) {
	for name, cas := range cases {
		t.Run(name,
			func(t *testing.T) {
				recv, err := BytesToStruct(cas.bytes, testStructA{})
				if err != nil {
					t.Errorf("error received converting bytes to struct: %v", err)
					return
				}

				if !reflect.DeepEqual(cas.struc, recv) {
					t.Errorf("Received unexpected result:")
					t.Errorf("Expected: %#v", cas.struc)
					t.Errorf("Received: %#v", recv)
				}
			},
		)
	}
}
