package schema

type Craft struct {
	Id   string `db:"item_id"`
	Name string `db:"name"`
}

type Material struct {
	ParentItemId string `db:"parent_item_id"`
	ChildItemId  string `db:"child_item_id"`
	ParentName   string `db:"parent_name"`
	ChildName    string `db:"child_name"`
	Unit         int    `db:"unit"`
	Total        int    `db:"total"`
}
