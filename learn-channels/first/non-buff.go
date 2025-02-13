package first

import (
	"fmt"
	"time"
)

func nonBuff() {
	userch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		userch <- "Bob"
	}()
	user := <-userch

	fmt.Println(user)

}
