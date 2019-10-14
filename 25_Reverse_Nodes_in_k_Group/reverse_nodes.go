package lc_0025

/**
 * Definition for singly-linked list.
 * */
type ListNode struct {
      Val int
      Next *ListNode
  }
 
  func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy:=&ListNode{}
    dummy.Next=head
    prev,end:=dummy,dummy

    for end.Next!=nil{
      for i := 0; i < k&&end!=nil; i++ {
        end=end.Next
      }
      if end ==nil {
        break
      }
      s,n:=prev.Next,end.Next
      end.Next=nil
      prev.Next=reverse(s)
      s.Next=n
      prev,end=s,s
    }
    return dummy.Next
  }

  func reverse(head *ListNode) *ListNode {
    pre := &ListNode{}
    cur := head
    for cur != nil {
      next := cur.Next
      cur.Next = pre
      pre = cur
      cur = next
    }
  
    return pre
  }