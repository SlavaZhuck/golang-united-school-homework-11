package batch

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	ctx := context.TODO()
	sem := semaphore.NewWeighted(pool)
	var temp = make([]user, n, n)

	for i := 0; int64(i) < n; i++ {
		// When maxWorkers goroutines are in flight, Acquire blocks until one of the
		// workers finishes.
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(i int64) {
			defer sem.Release(1)
			temp[i] = getOne(int64(i))
		}(int64(i))
	}

	if err := sem.Acquire(ctx, pool); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}

	return temp
}
