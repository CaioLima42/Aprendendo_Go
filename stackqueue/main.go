package main

import(
	"fmt"
)

type Stack struct{
	stack []int
}

func NewStack()Stack{
	return Stack{stack: []int{}}
}

func(s *Stack) AddToStack(val int){
	s.stack = append([]int{val}, s.stack...)
}
func(s *Stack) RemoveItemFromStack()(int, error){
	if len(s.stack) == 0 {
		return 0,fmt.Errorf("ta vazia essa porra fela da puta")
	}
	deletedValue := s.stack[0]
	if len(s.stack) == 1{
		s.stack = []int{}
	}else{
		s.stack = s.stack[1:len(s.stack)]
	}
	return deletedValue, nil

}


type Queue struct{
	queue []int
} 

func NewQueue()Queue{
	return Queue{queue: []int{}}
}

func (q *Queue) AddtoQueue(val int){
	q.queue = append(q.queue, val)
}

func (q *Queue) RemoveItemFromQueue()(int, error){
	if len(q.queue) == 0 {
		return 0,fmt.Errorf("ta vazia essa porra fela da puta")
	}
	deletedValue := q.queue[0]
	if len(q.queue) == 1{
		q.queue = []int{}
	}else{
		q.queue = q.queue[1:len(q.queue)]
	}
	return deletedValue, nil

}


func main(){
	
	stack := NewStack()
	stack.AddToStack(1)
	stack.AddToStack(2)
	stack.AddToStack(3)
	fmt.Println(stack.stack)
	fmt.Println(stack.RemoveItemFromStack())
	fmt.Println(stack.stack)
	
	q := NewQueue()
	q.AddtoQueue(1)
	q.AddtoQueue(2)
	q.AddtoQueue(3)
	fmt.Println(q.queue)
	fmt.Println(q.RemoveItemFromQueue())
	fmt.Println(q.queue)


}