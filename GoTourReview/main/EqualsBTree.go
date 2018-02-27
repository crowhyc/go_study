package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) recursionInorderTraver(c chan int) {
	if t != nil {
		c <- t.Value
		t.Left.recursionInorderTraver(c)
		fmt.Printf(".%v", t.Value)
		t.Right.recursionInorderTraver(c)
	}
}
func (t *Tree) loopInorderTraver(c chan int) {
	var idx = 0
	stack := make([]*Tree, 10)
	pnode := t
	//循环的时候先找到最左的叶子节点,然后将这个叶子节点之前的父节点按
	//压栈的反顺序进行出栈,输出结果后再判断是否此父节点还有右子树,如果有右子树则继续循环右子树
	//进行与之前类似的压栈出栈操作
	for ; pnode != nil || idx > 0; {
		for ; pnode != nil; {
			stack[idx] = pnode
			idx += 1
			pnode = pnode.Left
		}
		pnode = stack[idx-1]
		idx -= 1
		fmt.Printf("%v.", pnode.Value)
		c <- pnode.Value
		pnode = pnode.Right
	}
	close(c)
}
func (t *Tree) appendLeft(left *Tree) {
	t.Left = left
}
func (t *Tree) appendRight(right *Tree) {
	t.Right = right
}
func (t *Tree) newTreeNode(value int) *Tree {
	t.Value = value
	return t
}
func initATree() *Tree {
	root := Tree{Value: 3}
	l1 := Tree{Value: 1}
	root.appendLeft(&l1)
	r1 := Tree{Value: 8}
	root.appendRight(&r1)
	l12 := Tree{Value: 1}
	l1.appendLeft(&l12)
	l1r2 := Tree{Value: 2}
	l1.appendRight(&l1r2)
	r1l2 := Tree{Value: 5}
	r1.appendLeft(&r1l2)
	r12 := Tree{Value: 13}
	r1.appendRight(&r12)
	return &root

}
func main() {
	aTree := initATree()
	aCh := make(chan int, 100)
	aTree.loopInorderTraver(aCh)
	fmt.Println()
	for i := range aCh {
		fmt.Println(i)
	}
}
