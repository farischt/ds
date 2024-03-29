package heap

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	heap := New[int](MinHeap)
	if heap == nil {
		t.Error("should create new heap, test new heap")
	}
}

func TestHeapSize(t *testing.T) {
	heap := New[int](MinHeap)
	if heap.size() != 0 {
		t.Error("should return correct size, test heap size")
	}
}

func TestIsEmpty(t *testing.T) {
	heap := New[int](MinHeap)
	if !heap.isEmpty() {
		t.Error("should be empty, test heap is empty")
	}

	heap.data = append(heap.data, *NewItem(0, nil))
	if heap.isEmpty() {
		t.Error("should not be empty, test heap is not empty")
	}
}

func TestHeapLess(t *testing.T) {
	heap := New[int](MinHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if !heap.less(0, 1) {
		t.Error("should be less, test heap less")
	}
}

func TestHeapGreater(t *testing.T) {
	heap := New[int](MaxHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if heap.greater(0, 1) {
		t.Error("should be greater, test heap greater")
	}
}

func TestHeapCompare(t *testing.T) {
	heap := New[int](MinHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if !heap.compare(0, 1) {
		t.Error("should compare correctly for min heap, test heap compare")
	}

	maxHeap := New[int](MaxHeap)
	maxHeap.data = append(heap.data, *NewItem(0, nil))
	maxHeap.data = append(heap.data, *NewItem(1, nil))

	if maxHeap.compare(0, 1) {
		t.Error("should compare correctly for max heap, test heap compare")
	}
}

func TestHeapSwap(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *secondItem)

	heap.swap(0, 1)

	if heap.data[0] != *secondItem {
		t.Error("should swap correctly, test heap swap")
	} else if heap.data[1] != *firstItem {
		t.Error("should swap correctly, test heap swap")
	}
}

func TestHeapParent(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *secondItem)

	parentIndex := heap.parent(1)

	if parentIndex != 0 {
		t.Error("should get parent index correctly, test heap parent")
	}
}

func TestHeapLeft(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *secondItem)

	leftIndex := heap.left(0)

	if leftIndex != 1 {
		t.Error("should get left index correctly, test heap left")
	}
}

func TestHeapRight(t *testing.T) {
	heap := New[int](MinHeap)

	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	rightIndex := heap.right(0)

	if rightIndex != 2 {
		t.Error("should get right index correctly, test heap right")
	}
}

func TestIsInBound(t *testing.T) {
	heap := New[int](MinHeap)
	heap.data = append(heap.data, *NewItem(0, nil))
	heap.data = append(heap.data, *NewItem(1, nil))

	if !heap.isInBound(1) {
		t.Error("should return true for valid child, test heap is in bound")
	}
}

func TestHeapUp(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(1, nil)

	heap.data = append(heap.data, *secondItem)
	heap.data = append(heap.data, *firstItem)

	heap.up(1)

	if heap.data[0] != *firstItem {
		t.Error("should bubble up correctly, test heap up")
	} else if heap.data[1] != *secondItem {
		t.Error("should bubble up correctly, test heap up")
	}
}

func TestHeapDown(t *testing.T) {
	heap := New[int](MinHeap)

	firstItem := NewItem(0, nil)
	secondItem := NewItem(5, nil)
	thirdItem := NewItem(3, nil)

	heap.data = append(heap.data, *secondItem)
	heap.data = append(heap.data, *firstItem)
	heap.data = append(heap.data, *thirdItem)

	heap.down(0)

	if heap.data[0] != *firstItem {
		t.Error("should bubble down correctly, test heap down")
	} else if heap.data[1] != *secondItem {
		t.Error("should bubble down correctly, test heap down")
	}
}

func TestHeapPush(t *testing.T) {
	heap := New[int](MinHeap)
	item := NewItem(1, nil)
	heap.Push(item.Value, item.Information)
	if heap.data[0] != *item {
		t.Error("should correctly insert a first element into heap, test heap push")
	}
}

func TestHeapPush2(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(1, nil)
	heap.Push(top.Value, top.Information)
	heap.Push(2, nil)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert 2 elements into heap, and maintain heap property. Top element should be %d, test heap push 2", top)
	}
}

func TestHeapPush3(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(top.Value, top.Information)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert 3 elements into heap, and maintain heap property. Top element should be %d, test heap push 3", top)
	}
}

func TestHeapPush3MaxHeap(t *testing.T) {
	heap := New[int](MaxHeap)
	top := NewItem(10, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(top.Value, top.Information)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert 3 elements into heap, and maintain heap property. Top element should be %d, test heap push 3 max heap", top)
	}
}

func TestHeapPushMany(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(12, nil)
	heap.Push(3, nil)
	heap.Push(4, nil)
	heap.Push(20, nil)
	heap.Push(5, nil)
	heap.Push(12, nil)
	heap.Push(6, nil)
	heap.Push(7, nil)
	heap.Push(8, nil)
	heap.Push(top.Value, top.Information)

	if heap.data[0] != *top {
		t.Errorf("should correctly insert many elements into heap, and maintain heap property. Top element should be %d, test heap push many", top)
	}
}

func TestHeapPop(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(top.Value, top.Information)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %v, test heap pop", top)
	}
}

func TestHeapPop2(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(12, nil)
	heap.Push(top.Value, top.Information)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d, test heap pop 2", top)
	}
}

func TestHeapPopEmpty(t *testing.T) {
	heap := New[int](MinHeap)

	_, err := heap.Pop()

	if err == nil {
		t.Error("should return error when pop from empty heap, test heap pop empty")
	}
}

func TestHeapPopOneElement(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(top.Value, top.Information)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d, test heap pop one element", top)
	}
}

func TestHeapPopManyElement(t *testing.T) {
	heap := New[int](MinHeap)
	top := NewItem(-1, nil)
	heap.Push(top.Value, top.Information)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(15, nil)
	heap.Push(1, nil)
	heap.Push(7, nil)
	heap.Push(9, nil)
	heap.Push(11, nil)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d, test heap pop many element", top)
	}
}

func TestHeapPopManyMaxHeap(t *testing.T) {
	heap := New[int](MaxHeap)
	top := NewItem(15, nil)
	heap.Push(top.Value, top.Information)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(12, nil)
	heap.Push(1, nil)
	heap.Push(2, nil)
	heap.Push(3, nil)
	heap.Push(-1, nil)
	heap.Push(1, nil)
	heap.Push(7, nil)
	heap.Push(9, nil)
	heap.Push(11, nil)

	pop, _ := heap.Pop()

	if *pop != *top {
		t.Errorf("should pop the top element %d, test heap pop many max heap", top)
	}
}
