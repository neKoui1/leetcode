import java.util.ArrayList;

public class Solution {
    public class ListNode{
        int val;
        ListNode next;
        ListNode(){}
        ListNode(int val){this.val = val;}
        ListNode(int val, ListNode next){this.val = val; this.next = next;}
    }
    public ListNode removeNthFromEnd(ListNode head, int n) {
        int cnt = 0;
        ListNode tempHead = new ListNode(0,head);
        int cnt2 = 1;
        ListNode p = head;
        while(p != null){
            cnt ++;
            p = p.next;
        }

        ListNode q = tempHead;
        while(q != null){
            if(cnt2 == cnt - n + 1){
                q.next = q.next.next;
            }
            q = q.next;
            cnt2++;
        }
        return tempHead.next;
    }
}
