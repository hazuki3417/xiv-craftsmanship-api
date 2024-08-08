package schema

type Craft struct {
	ItemId   string `db:"item_id"`
	ItemName string `db:"item_name"`
}

type MaterialTree struct {
	ParentItemId string `db:"parent_item_id"`
	ChildItemId  string `db:"child_item_id"`
	ParentName   string `db:"parent_name"`
	ChildName    string `db:"child_name"`
	Unit         int    `db:"unit"`
	Total        int    `db:"total"`
	X            int    `db:"x"`
	Y            int    `db:"y"`
	NodeType     string `db:"node_type"`
}
