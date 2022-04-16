package flatten_a_multilevel_doubly_linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func Flatten(root *Node) *Node {
	if root != nil {
		flatten(root)
	}
	return root
}

func flatten(p *Node) *Node {
	for p.Next != nil {
		tmp := p.Next
		if p.Child != nil {
			p.Next = p.Child
			p.Child.Prev = p

			child := flatten(p.Child)
			child.Next = tmp
			tmp.Prev = child

			p.Child = nil
		}
		p = tmp
	}

	if p.Child != nil {
		p.Next = p.Child
		p.Child.Prev = p

		child := flatten(p.Child)

		p.Child = nil

		return child
	}

	return p
}
