package populating_next_right_pointers_in_each_node_II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}

	for len(queue) != 0 {
		tmpQueue := []*Node{}
		for i, val := range queue {
			if i != len(queue)-1 {
				val.Next = queue[i+1]
			}
			if val.Left != nil {
				tmpQueue = append(tmpQueue, val.Left)
			}
			if val.Right != nil {
				tmpQueue = append(tmpQueue, val.Right)
			}
		}
		queue = tmpQueue
	}

	return root
}
