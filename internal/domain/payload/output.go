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
	ID    string
	Name  string
	Unit  int
	Total int
	Depth int
}

type Edge struct {
	Source string
	Target string
}
