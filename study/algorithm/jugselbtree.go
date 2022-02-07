package algorithm

type Btreenode struct {
	Value          int
	Lchild, Rchild *Btreenode
}

type Btree struct {
	Root *Btreenode
}

func (this *Btree) Breeinsert(a int) {
	if this.Root == nil {
		//at := &Btreenode{Value: a}
		this.Root = &Btreenode{Value: a}
		return
	}
	bt := this.Root
	for {
		if bt.Value >= a {
			if bt.Lchild != nil {
				bt = bt.Lchild
			} else {
				bt.Lchild = &Btreenode{Value: a}
				return
			}
		} else {
			if bt.Rchild != nil {
				bt = bt.Rchild
			} else {
				bt.Rchild = &Btreenode{Value: a}
				return
			}
		}
	}
}

func Jugselbtree(root *Btreenode, arry []int) bool {
	if root == nil {
		if len(arry) == 0 || len(arry) == 1 {
			return true
		} else {
			return false
		}
	} else {
		var right int = len(arry) - 1
		end := len(arry) - 1
		for i, v := range arry {
			if v > arry[len(arry)-1] {
				right = i
				break
			}
		}
		if arry[end] == root.Value {
			if Jugselbtree(root.Lchild, arry[:right]) && Jugselbtree(root.Rchild, arry[right:end]) {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}

}

func Postorder_traversal(bt *Btreenode, arry *[]int) {
	if bt == nil {
		return
	} else {
		Postorder_traversal(bt.Lchild, arry)
		Postorder_traversal(bt.Rchild, arry)
		*arry = append(*arry, bt.Value)
	}
}
