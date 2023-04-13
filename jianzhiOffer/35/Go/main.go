package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	// 复制结点并插入原链表
	for cur != nil {
		curCopy := &Node{
			Val: cur.Val,
		}
		curCopy.Next = cur.Next
		cur.Next = curCopy
		cur = curCopy.Next
	}
	// 构造random
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 拆链表
	cur = head
	newHead := head.Next
	for cur.Next != nil {
		temp := cur.Next
		cur.Next = temp.Next
		cur = temp
	}
	return newHead
}
func main() {

}
