package two_stacks_queue

type Stack interface {
	Push(i int)
	Pop() (int, bool)
	Len() int
}

type Queue struct {
	stackPush Stack
	stackPop  Stack
}

func (q *Queue) Push(i int) {
	q.stackPush.Push(i)
}

func (q *Queue) Poll() (int, bool) {
	if q.stackPop.Len() == 0 {
		v, ok := q.stackPush.Pop()
		for ok {
			q.stackPop.Push(v)
			v, ok = q.stackPush.Pop()
		}
	}

	return q.stackPop.Pop()
}
