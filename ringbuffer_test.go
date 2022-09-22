package ringbuffer

import (
	"testing"
)

func TestFunc(t *testing.T) {
	rb := New(4)
	if cap := rb.Capacity(); cap != 4 {
		t.Fatalf("Test failed. Capacity should be 4, but it is %v", cap)
	}
	if err := rb.Push(1, 2, 3); err != nil {
		t.Fatal(err)
	}
	if err := rb.Push(4); err != nil {
		t.Fatal(err)
	}
	occupied := rb.Occupied()
	remaining := rb.Remaining()
	is_full := rb.IsFull()
	is_empty := rb.IsEmpty()

	if occupied != 4 {
		t.Fatalf("Test failed. Occupied should be 4, but it is %v", occupied)
	}
	if remaining != 0 {
		t.Fatalf("Test failed. Remaining should be 0, but it is %v", remaining)
	}
	if is_full != true {
		t.Fatalf("Test failed. is_full should be true, but it is true")
	}
	if is_empty != false {
		t.Fatalf("Test failed. is_empty should be false, but it is true")
	}

	rb.Pop()
	rb.PopMany(2)

	occupied = rb.Occupied()
	remaining = rb.Remaining()

	if occupied != 1 {
		t.Fatalf("Test failed. Occupied should be 1, but it is %v", occupied)
	}
	if remaining != 3 {
		t.Fatalf("Test failed. Remaining should be 3, but it is %v", remaining)
	}

	rb.Clear()

	if rb.IsEmpty() != true {
		t.Fatalf("Test failed. RingBuffer should be empty after calling Clear()")
	}

	rb.Push(3)
	if v, err := rb.Pop(); v != 3 || err != nil {
		t.Errorf("Test failed. The popped value should be 3, but it is %v. Error is %v", v, err)
	}

	rb.Push(1, 2)
	v, err := rb.PopMany(2)
	if err != nil {
		t.Fatal(err)
	}
	if v[0] != 1 {
		t.Errorf("Test failed. Expected value 1, but value is %v", v[0])
	}

	if v[1] != 2 {
		t.Errorf("Test failed. Expected value 2, but value is %v", v[0])
	}
}
