package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var currentL1, currentL2, firstNode, prevNode *ListNode
	currentL1 = l1
	currentL2 = l2
	var carry, i int
	for {
		if shouldFinish(currentL1, currentL2, carry) {
			break
		}

		var l1Val, l2Val int
		if currentL1 != nil {
			l1Val = currentL1.Val
		}
		if currentL2 != nil {
			l2Val = currentL2.Val
		}

		var newNode *ListNode
		newNode, carry = createNode(l1Val, l2Val, carry)

		// 呼び出し元に返すのは最初のNode
		if i == 0 {
			firstNode = newNode
		} else {
			prevNode.Next = newNode
		}

		if currentL1 != nil {
			currentL1 = currentL1.Next
		}
		if currentL2 != nil {
			currentL2 = currentL2.Next
		}
		prevNode = newNode
		i++
	}

	return firstNode
}

func shouldFinish(l1, l2 *ListNode, carry int) bool {
	return l1 == nil && l2 == nil && carry == 0
}

func createNode(val1, val2, carry int) (*ListNode, int) {
	val, newCarry := calcVal(val1, val2, carry)
	return &ListNode{
		Val:  val,
		Next: nil,
	}, newCarry
}

func calcVal(val1, val2, carry int) (int, int) {
	sum := val1 + val2 + carry
	var newCarry int
	var newVal int
	// 繰り上がりがあるならならキャリーと値を分割
	if sum > 9 {
		newVal = sum % 10              // あまりがその桁値
		newCarry = (sum - newVal) / 10 // 次の桁に持っていく値
	} else {
		newVal = sum // 繰り上がりがないなら、そのまま
	}
	return newVal, newCarry
}
