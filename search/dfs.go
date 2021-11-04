package search

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//if node.Value == condition {
		//	do something
		//}

		left := node.Left
		right := node.Right
		if left != nil {
			stack = append(stack, left)
		}
		if right != nil {
			stack = append(stack, right)
		}
	}
}
