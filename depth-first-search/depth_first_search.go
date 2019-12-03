package main

type Node struct {
	Name     string
	Children []*Node
}

// O(log(v + e)) time, O(log(v)) space // where v stands for nodes and e stands for relations
func (n *Node) DepthFirstSearch(arr []string) []string {
	arr = append(arr, n.Name)
	for _, children := range n.Children {
		arr = children.DepthFirstSearch(arr)
	}
	return arr
}

/* Example:
     A
    / \
   C   B
       /
			G

func main() {
	n := &Node{Name: "A"}
	n.Children = append(n.Children, &Node{Name: "C"}, &Node{Name: "B", Children: []*Node{{Name: "G"}}})
	arr := make([]string, 0)
	fmt.Println(n.DepthFirstSearch(arr))
}
*/
