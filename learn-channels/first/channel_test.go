package first

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAddUsers(t *testing.T) {

	b := false
	fmt.Println(unsafe.Sizeof(b))

	// server := NewServer()
	// server.Start()

	// for i := 0; i < 10; i++ {
	// 	// go func(i int) {
	// 	server.userch <- fmt.Sprintf("user-%d", i)
	// 	// }(i)

	// }
	// fmt.Println("loop is done")
}
