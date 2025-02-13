package first

import (
	"fmt"
	"time"
)

func main() {
	userch := make(chan string, 1000)

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-userch:
		}
	}

}
