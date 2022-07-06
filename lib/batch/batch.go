package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, pool)

	//ouput array
	var temp = make([]user, n, n)

	for i := 0; int64(i) < n; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int64) {
			temp[i] = getOne(i)
			<-sem
			wg.Done()
		}(int64(i))
	}
	wg.Wait()
	return temp
}
