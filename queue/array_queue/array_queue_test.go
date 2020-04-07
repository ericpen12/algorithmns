package array_queue_test

import (
	"datastures/queue/array_queue"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := array_queue.NewQueue(20)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
	t.Log(queue.Dequeue())
}