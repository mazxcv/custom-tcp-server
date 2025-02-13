package context

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	val int
	err error
}

func con() {
	start := time.Now()
	ctx := context.Background()
	userId := 10
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	resch := make(chan Response, 1)

	go func() {
		time.Sleep(time.Microsecond * 2500)
		resch <- Response{val: 666 + userId, err: nil}
	}()

	for {
		select {
		case resp := <-resch:
			return resp.val, resp.err
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long: %v", ctx.Err())
		}
	}
}
