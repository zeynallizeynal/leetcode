package main

type LinkedQueue struct {
	list *LinkedList
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{list: NewLinkedList()}
}

func (s *LinkedQueue) Push(e int) {
	s.list.PushHead(e)
}

func (s *LinkedQueue) Pop() interface{} {
	el := s.list.GetHead()

	if el != nil {
		s.list.DeleteHead()
		return el.Value
	}
	return nil
}

func (s *LinkedQueue) Peek() interface{} {
	el := s.list.GetHead()

	if el != nil {
		return el.Value
	}
	return nil
}

func (s *LinkedQueue) Size() interface{} {
	return s.list.Size
}

func (s *LinkedQueue) Print() {
	s.list.PrintFromTail()
}
