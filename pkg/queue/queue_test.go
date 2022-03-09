package queue_test

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/pkg/queue"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue(10)
	defer q.Close()
	for i := 0; i < 10; i++ {
		q.Push([]byte(strconv.Itoa(i)))
	}

	for i := 0; i < 10; i++ {

		data, err := q.WaitPop()
		if err != nil {
			t.Error(err)
			return
		}

		if string(data) != strconv.Itoa(i) {
			t.Errorf("expected %v but got %v", i, string(data))
		}

	}
}

func TestQueueMXLen(t *testing.T) {
	q := queue.NewQueue(5)
	defer q.Close()
	for i := 0; i < 10; i++ {
		q.Push([]byte(strconv.Itoa(i)))
	}

	for i := 5; i < 10; i++ {

		data, err := q.WaitPop()
		if err != nil {
			t.Error(err)
			return
		}

		if string(data) != strconv.Itoa(i) {
			t.Errorf("expected %v but got %v", i, string(data))
		}

	}
}

func TestQueueRace(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)
	q := queue.NewQueue(10)

	go func(wg *sync.WaitGroup, q queue.Queue) {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		<-ctx.Done()
		q.Close()

	}(&wg, q)

	go func(wg *sync.WaitGroup, q queue.Queue) {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				q.Push([]byte("test"))
			}
		}

	}(&wg, q)

	go func(wg *sync.WaitGroup, q queue.Queue) {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				q.WaitPop()
			}
		}

	}(&wg, q)

	wg.Wait()
}
