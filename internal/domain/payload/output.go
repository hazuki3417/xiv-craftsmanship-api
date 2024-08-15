package payload

type Level struct {
	Item  *int
	Craft int
}

type Craft struct {
	ID     string
	Name   string
	Job    string
	Pieces int
	Level  Level
}

type Material struct {
	ParentItemId string
	ChildItemId  string
	ParentName   string
	ChildName    string
	Unit         int
	Total        int
}
