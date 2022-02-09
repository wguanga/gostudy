package algorithm

import (
	"container/list"
	"fmt"
)

func Constructorforarry(arry []int) *Btreenode {
	if arry == nil {
		return &Btreenode{}
	}
	i := 0
	root := &Btreenode{Value: arry[i]}
	Root := root
	stk := list.New()
	stk.PushBack(root)
	for 2*(i+1) <= len(arry) {

		btn1, btn2 := &Btreenode{Value: arry[2*(i+1)-1]}, &Btreenode{Value: arry[2*(i+1)]}
		if arry[2*(i+1)-1] == 0 {
			btn1 = nil
		}
		if arry[2*(i+1)] == 0 {
			btn2 = nil
		}
		i++
		root = stk.Remove(stk.Front()).(*Btreenode)
		root.Lchild = btn1
		root.Rchild = btn2

		if root.Lchild != nil {
			stk.PushBack(root.Lchild)
		}
		if root.Rchild != nil {
			stk.PushBack(root.Rchild)
		}
	}
	return Root
}

func (this *Btreenode) Preorder_traversal(arry *[]int) {
	if this == nil {
		return
	}
	fmt.Printf("this.Value: %v\n", this.Value)
	*arry = append(*arry, this.Value)
	this.Lchild.Preorder_traversal(arry)
	this.Rchild.Preorder_traversal(arry)
}
