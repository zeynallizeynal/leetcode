package main

type CircularArray struct {
	Size, Count int    // size of the array and item count
	start, end  int    // indexes
	array       []*int // actual buffer
}

func NewCircularArray(size int) *CircularArray {
	s := size
	if size <= 0 {
		s = 1
	}
	return &CircularArray{
		Size:  s,
		array: make([]*int, s),
	}
}

func (c *CircularArray) AddFront(v int) {
	c.start = (c.start - 1 + c.Size) % c.Size
	c.array[c.start] = &v
	c.Count++
}

func (c *CircularArray) RemoveFront() *int {
	v := c.array[c.start]
	c.array[c.start] = nil
	c.start = (c.start + 1) % c.Size
	c.Count--
	return v
}

func (c *CircularArray) AddBack(v int) {
	c.array[c.end] = &v
	c.end = (c.end + 1) % c.Size
	c.Count++
}

func (c *CircularArray) RemoveBack() *int {
	c.end = (c.end - 1 + c.Size) % c.Size
	v := c.array[c.end]
	c.array[c.end] = nil
	c.Count--
	return v
}

func (c *CircularArray) GetFront() *int {
	return c.array[c.start]
}

func (c *CircularArray) GetBack() *int {
	return c.array[(c.end-1+c.Size)%c.Size]
}
