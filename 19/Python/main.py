class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        sentinel = ListNode(0)
        sentinel.next = head
        slow = sentinel
        fast = sentinel.next
        for i in range(n):
            fast = fast.next
        while fast != None:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
        return sentinel.next