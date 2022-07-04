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

	res = make([]user, 0, n)
	var mtx sync.Mutex
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, pool)

	for i := int64(0); i < n; i++ {

		wg.Add(1)
		semaphore <- struct{}{}

		go func(id int64) {
			user := getOne(id)
			mtx.Lock()
			res = append(res, user)
			mtx.Unlock()
			<-semaphore
			wg.Done()
		}(i)

	}

	wg.Wait()
	return
}
