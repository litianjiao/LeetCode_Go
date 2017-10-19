package addtwonum

type ListNode struct {
	Val  int
	Next *ListNode
}

//模拟加法
//这个解法是另开一个链表用于和L1 L2对应，对应node相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	sentinel := new(ListNode)
	d := sentinel
	var sum int
	//每次一个node加完之后去下一个node，不同位隔离
	for c1 != nil || c2 != nil {
		//因为是int型 所以只会保存sum十位上的整数
		//如果前一个node有进位，则本次node相加起始值为进位的1
		sum /= 10
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}
		d.Next = new(ListNode)
		d.Next.Val = sum % 10
		d = d.Next
	}
	//如果还有进位，则加上
	if sum/10 == 1 {
		d.Next = new(ListNode)
		d.Next.Val = 1
	}
	return sentinel.Next
}
