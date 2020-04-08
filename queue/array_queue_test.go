package queue_test

import (
	queue2 "datastures/queue"
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	queue := queue2.NewArrayQueue(20)

	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < 10; i++ {
		v, err := queue.Dequeue()
		if err != nil {
			t.Error(err)
		}
		if v != i {
			t.Errorf("the queue excepted is %d, but the actual is %d", i, v)
		}
	}
}

func TestArrayQueueOverflow(t *testing.T) {
	queue := queue2.NewArrayQueue(2)
	queue.Enqueue(1)
	queue.Enqueue(2)
	err := queue.Enqueue(3)
	if err == nil || err.Error() != "queue overflow" {
		t.Error("without underflow")
	}
	t.Log(err)
}

func TestArrayQueueUnderflow(t *testing.T) {
	queue := queue2.NewArrayQueue(2)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Dequeue()
	queue.Dequeue()
	_, err := queue.Dequeue()

	if err == nil || err.Error() != "queue underflow" {
		t.Error("without underflow")
	}
	t.Log(err)
}
