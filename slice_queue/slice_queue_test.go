package slice_queue

import "testing"

func TestQueueSinglePutPeekAndPop(t *testing.T) {
	queue := New(100)
	queue.Put(100)
	if queue.Peek() != 100 {
		t.Errorf("Peek didn't return expected 100. Slice: %v", queue.slice)
	}

	if *queue.Pop() != 100 {
		t.Errorf("Pop didn't return expected 100. Slice: %v", queue.slice)
	}

	if queue.Pop() != nil {
		t.Errorf("Pop didn't return expected nil. Slice: %v", queue.slice)
	}

}
