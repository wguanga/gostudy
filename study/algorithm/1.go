package algorithm

import "container/list"

type CQueue struct {
	s1, s2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.PushBack(value)
}

func (this *CQueue) DeleterHead() int {
	if this.s2.Len() == 0 {
		for this.s1.Len() != 0 {
			this.s2.PushBack(this.s1.Remove(this.s1.Back()))
		}
	}
	if this.s2.Len() != 0 {
		e := this.s2.Back()
		this.s2.Remove(e)
		return e.Value.(int)
	}
	return -1
}
