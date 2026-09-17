package main

type LinkedStack struct {
	list *LinkedList
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{list: NewLinkedList()}
}

func (s *LinkedStack) Push(e int) {
	s.list.PushTail(e)
}

func (s *LinkedStack) Pop() interface{} {
	el := s.list.GetTail()

	if el != nil {
		s.list.DeleteTail()
		return el.Value
	}
	return nil
}

func (s *LinkedStack) Peek() interface{} {
	el := s.list.GetTail()

	if el != nil {
		return el.Value
	}
	return nil
}

func (s *LinkedStack) Size() interface{} {
	return s.list.Size
}

func (s *LinkedStack) Print() {
	s.list.PrintFromHead()
}
