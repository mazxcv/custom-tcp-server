package first

import "fmt"

func buff() {
	userch := make(chan string, 2)

	userch <- "Bob"
	user := <-userch

	fmt.Println(user)

}
