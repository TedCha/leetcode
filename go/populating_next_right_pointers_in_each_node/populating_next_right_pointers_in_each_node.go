package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make(chan *Node, 4096)

	// initalize queue
	push(root, queue)
	for {
		node := pop(queue)
		if node == nil {
			break
		}

		if node.Left != nil && node.Right != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}

			push(node.Left, queue)
			push(node.Right, queue)
		}
	}

	return root
}

func push(node *Node, q chan *Node) {
	q <- node
}

func pop(q chan *Node) *Node {
	var node *Node
	select {
	case node = <-q:
		return node
	default:
		close(q)
		return node
	}
}
