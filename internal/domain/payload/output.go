package payload

type Craft struct {
	ID   string
	Name string
}

type Recipe struct {
	Nodes []Node
	Edges []Edge
}

type Node struct {
	ID       string
	Name     string
	Unit     int
	Total    int
	X        int
	Y        int
	NodeType string
}

type Edge struct {
	Source string
	Target string
}
