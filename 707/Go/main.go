package main

//type LinkNode struct {
//	Val  int
//	Next *LinkNode
//}
//type MyLinkedList struct {
//	Sentinel *LinkNode
//	Length   int
//}
//
//func Constructor() MyLinkedList {
//	return MyLinkedList{
//		&LinkNode{0, nil},
//		0,
//	}
//}
//
//func (this *MyLinkedList) Get(index int) int {
//	if index > this.Length {
//		return -1
//	}
//	cnt := 0
//	for p := this.Sentinel.Next; p != nil; p = p.Next {
//		if cnt == index-1 {
//			return p.Val
//		}
//		cnt++
//	}
//	return 0
//}
//
//func (this *MyLinkedList) AddAtHead(val int) {
//	node := &LinkNode{
//		val,
//		this.Sentinel.Next,
//	}
//	this.Sentinel.Next = node
//	this.Length++
//}
//
//func (this *MyLinkedList) AddAtTail(val int) {
//	for p := this.Sentinel.Next; p != nil; p = p.Next {
//		if p.Next == nil {
//			p.Next = &LinkNode{
//				val,
//				nil,
//			}
//			this.Length++
//			break
//		}
//	}
//}
//
//func (this *MyLinkedList) AddAtIndex(index int, val int) {
//	if index < 0 {
//		node := &LinkNode{
//			val,
//			this.Sentinel.Next,
//		}
//		this.Sentinel.Next = node
//		this.Length++
//		return
//	} else if index > this.Length {
//		return
//	} else {
//		cnt := 0
//		for p := this.Sentinel.Next; p != nil; p = p.Next {
//			if cnt == index-1 {
//				node := &LinkNode{
//					val,
//					p.Next,
//				}
//				p.Next = node
//				this.Length++
//				return
//			}
//			cnt++
//		}
//	}
//}
//
//func (this *MyLinkedList) DeleteAtIndex(index int) {
//	if index < 0 || index > this.Length {
//		return
//	}
//	cnt := 0
//	for p, q := this.Sentinel.Next, this.Sentinel; p != nil; p, q = p.Next, q.Next {
//		if cnt == index-1 {
//			q.Next = p.Next
//			this.Length--
//			return
//		}
//		cnt++
//	}
//}
type Node struct {
	Val  int
	Next *Node
}
type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&Node{}, 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(index, 0)
	l.size++
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &Node{val, pred.Next}
	pred.Next = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	l.size--
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {

}
