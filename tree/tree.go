package tree

import "log"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
}

type seqStack struct {
	data [100]*TreeNode
	top  int
}

func (t *Tree) Inert(value int) {
	if t.root == nil {
		t.root = &TreeNode{
			data:  value,
			left:  nil,
			right: nil,
		}
	} else {
		t.root.Insert(value)
	}
}

func (b *TreeNode) Insert(value int) {
	if b == nil {
		return
	} else if value <= b.data {
		if b.left == nil {
			b.left = &TreeNode{
				data:  value,
				left:  nil,
				right: nil,
			}
		} else {
			b.left.Insert(value)
		}
	} else {
		if b.right == nil {
			b.right = &TreeNode{
				data:  value,
				left:  nil,
				right: nil,
			}
		} else {
			b.right.Insert(value)
		}
	}
}

func (t *Tree) PreOrder() (result []int) {
	var (
		s seqStack
	)

	s.top = -1

	if t.root == nil {
		panic("no data")
	} else {
		for t.root != nil || s.top != -1 {
			for t.root != nil {
				result = append(result, t.root.data)
				s.top++
				s.data[s.top] = t.root
				t.root = t.root.left
			}
			s.top--
			t.root = s.data[s.top+1].right
		}
	}
	return
}

func (t *Tree) InOrder() (result []int) {
	var (
		s seqStack
	)

	s.top = -1

	if t.root == nil {
		panic("no data")
	} else {
		for t.root != nil || s.top != -1 {
			log.Println("enter")
			for t.root != nil {
				s.top++
				s.data[s.top] = t.root
				t.root = t.root.left
			}

			s.top--
			t.root = s.data[s.top+1]
			result = append(result, t.root.data)
			log.Print(result)
			t.root = t.root.right
			if t.root != nil{
				log.Println("right",t.root.data)
			}
		}
	}
	return
}

type seqStack2 struct {
	data [100]*TreeNode
	top int
}
func (t *Tree)Inorder2() (result []int) {
	var(
		s seqStack2
	)

	s.top = -1

	if t.root == nil{
		panic("no data")
	}else{
		for s.top != -1 && t.root != nil{
			for t.root != nil{
				s.top++
				s.data[s.top] = t.root
				t.root = t.root.left
			}

			s.top--
			t.root = s.data[s.top+1]
			result = append(result,t.root.data)
			t.root = t.root.right
		}
	}
	return
}


func (t *Tree)Preorder2() (result []int) {
	var(
		s seqStack2
	)

	s.top = -1

	if t.root == nil {
		panic("no data")
	}else{
		for t.root != nil && s.top != -1 {
			for t.root != nil {
				result = append(result,t.root.data)
				s.top++
				s.data[s.top] = t.root
				t.root = t.root.left
			}

			s.top--
			t.root = s.data[s.top+1]
			t.root = t.root.right
		}
	}

}