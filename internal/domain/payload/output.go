package payload

type Craft struct {
	ID   string
	Name string
}

type Material struct {
	ParentItemId string
	ChildItemId  string
	ParentName   string
	ChildName    string
	Unit         int
	Total        int
}
