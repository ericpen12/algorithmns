package search

func bfs(root *Node) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]

			left := node.Left
			right := node.Right

			//if node.Value == someCondition {
			//
			//}

			if left != nil {
				queue = append(queue, left)
			}
			if right != nil {
				queue = append(queue, right)
			}
		}
	}
}
