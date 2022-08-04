package chance

import (
	"crypto/rand"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func RandomInt(max int) (int, error) {
	d := make([]byte, sizeOfInt(max))
	err := randomData(d)
	if err != nil {
		return 0, err
	}

	fmt.Println(d)
	fmt.Println(*byteSliceToIntPtr(d))

	return 0, nil

}

// scary shit; understand before use
func byteSliceToIntPtr(b []byte) *int {
	shdr := *(*reflect.SliceHeader)(unsafe.Pointer(&b))

	return unsafe.Pointer(shdr.Data)
}

// returns size (in bytes) required to hold i
func sizeOfInt(i int) int {
	n := 0
	for i != 0 {
		i >>= 8
		n++
	}

	return n
}

func randomData(b []byte) error {
	fs, err := os.Open("/dev/random")
	if err != nil {
		goto err
	}

	_, err = fs.Read(b)
	if err != nil {
		goto err
	}

	return nil

err:
	_, err = rand.Read(b)
	return err
}
