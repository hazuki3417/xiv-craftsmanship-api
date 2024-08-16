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

type Parent struct {
	ItemId     string
	ItemName   string
	CraftJob   string
	CraftLevel int
}

type Child struct {
	ItemId    string
	ItemName  string
	ItemType  string
	ItemUnit  int
	ItemTotal int
}

type Material struct {
	TreeId string
	Parent Parent
	Child  Child
}
