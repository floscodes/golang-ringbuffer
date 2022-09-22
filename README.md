# An easy to use RingBuffer in Go

First of all run:
`go get https://github.com/floscodes/golang-ringbuffer`

## Usage:

```go
func main() {
	rb := ringbuffer.New(6) // creates a new ringbuffer with capacity 6

	rb.Push(1)       // pushes value 1 to buffer
	rb.Push(2, 3)    // pushes values 2 and 3 to buffer
	rb.Push(4, 5, 6) // pushes values 4, 5 and 6 to buffer

	v, err := rb.Pop()  // returns value 1 from buffer
	fmt.Println(v, err) // Output: 1 <nil>

	v, err = rb.PopMany(2) // returns a slice with the values 2 and 3 from the buffer
	fmt.Println(v, err)    // Output: [2 3] <nil>

}
```
**read the docs for more ->** [![Go Reference](https://pkg.go.dev/badge/github.com/floscodes/golang-ringbuffer.svg)](https://pkg.go.dev/github.com/floscodes/golang-ringbuffer)
