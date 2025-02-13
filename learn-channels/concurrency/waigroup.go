package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func concurrency() {
	now := time.Now()
	respch := make(chan string, 3)
	userID := 10

	wg := &sync.WaitGroup{}

	go fetchUserData(userID, respch, wg)
	go fetchUserRecommendations(userID, respch, wg)
	go fetchUserLikes(userID, respch, wg)

	wg.Add(3)
	wg.Wait()

	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	respch <- fmt.Sprintf("user-data %d", userID)
	wg.Done()
}

func fetchUserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	respch <- fmt.Sprintf("user-recommendations %d", userID)
	wg.Done()
}

func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	respch <- fmt.Sprintf("user-likes %d", userID)
	wg.Done()
}
