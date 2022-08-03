package my_calendar_I

type MyCalendar struct {
	root *treeNode
}

type treeNode struct {
	start, end  int
	left, right *treeNode
}

func Constructor() MyCalendar {
	return MyCalendar{
		root: nil,
	}
}

func (c *MyCalendar) Book(start int, end int) bool {
	if end < start || start < 0 {
		return false
	}

	return insert(&c.root, start, end)
}

func insert(node **treeNode, start, end int) bool {
	if *node == nil {
		*node = &treeNode{
			start: start,
			end:   end,
			left:  nil,
			right: nil,
		}
		return true
	}

	if (*node).start >= end {
		return insert(&(*node).left, start, end)
	} else if (*node).end <= start {
		return insert(&(*node).right, start, end)
	}

	return false
}
