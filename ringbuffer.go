// Written by flopetautschnig (floscodes) (c) 2022

package ringbuffer

import (
	"errors"
)

type RingBuffer struct {
	reader  int // reads content, represents the tail of the RingBufer
	writer  int // represents the head of the RingBuffer where data is going to be written
	content []interface{}
}

// Returns a new RingBuffer with capacity s
func New(s uint) RingBuffer {
	return RingBuffer{0, 0, make([]interface{}, s+1, s+1)} // increasing the capacity by one in order tomake the ring flow possible
}

// Pushes one or more elements to the Ringbuffer
func (rb *RingBuffer) Push(elements ...interface{}) error {
	if rb.IsFull() {
		return errors.New("buffer overflow")
	} else {
		if len(elements) == 0 {
			return errors.New("no element has been pushed")
		}
		for _, e := range elements {
			rb.content[rb.writer] = e
			rb.moveWriter()
		}
		return nil
	}
}

// Returns the oldest element of the RingBuffer
func (rb *RingBuffer) Pop() (interface{}, error) {
	if rb.IsEmpty() {
		return nil, errors.New("buffer is empty")
	} else {
		defer rb.moveReader()
		return rb.content[rb.reader], nil
	}
}

// Returns a given number(n) of elements of the RingBuffer starting with the oldest one
func (rb *RingBuffer) PopMany(number uint) ([]interface{}, error) {
	output_slice := make([]interface{}, 0)
	for i := 0; i < int(number); i++ {
		element, err := rb.Pop()
		if err != nil {
			return output_slice, err
		}
		output_slice = append(output_slice, element)
	}
	return output_slice, nil
}

// Returns the capacity of the RingBuffer
func (rb *RingBuffer) Capacity() int {
	return cap(rb.content) - 1 // reducing output value by one because it has been increased by one in the making process (as explanied in the New-func)
}

// Returns the occupied capacity of the RingBuffer
func (rb *RingBuffer) Occupied() int {
	reader := rb.reader
	v := 0
	if rb.IsEmpty() {
		return v
	}
	for {
		reader++
		if reader > len(rb.content)-1 {
			reader = 0
		}
		v++
		if reader == rb.writer {
			break
		}
	}
	return v
}

// Returns the remaining capacity of the RingBuffer
func (rb *RingBuffer) Remaining() int {
	return rb.Capacity() - rb.Occupied()
}

// Checks if the Ringbuffer is full
func (rb *RingBuffer) IsFull() bool {
	if rb.writer+1 == rb.reader {
		return true
	} else if rb.writer == len(rb.content)-1 && rb.reader == 0 {
		return true
	} else {
		return false
	}
}

// Checks if the RingBuffer ist empty
func (rb *RingBuffer) IsEmpty() bool {
	return rb.reader == rb.writer
}

// Emptys the RingBuffer
func (rb *RingBuffer) Clear() {
	rb.reader = 0
	rb.writer = 0
	rb.content = make([]interface{}, len(rb.content), cap(rb.content))
}

// these funcs are not public - they are only an eneasement for moving reader and writer

func (rb *RingBuffer) moveReader() {
	rb.reader++
	if rb.reader > len(rb.content)-1 {
		rb.reader = 0
	}
}

func (rb *RingBuffer) moveWriter() {
	rb.writer++
	if rb.writer > len(rb.content)-1 {
		rb.writer = 0
	}
}
