package schema

type Craft struct {
	Id          string `db:"recipe_id"`
	ItemId      string `db:"item_id"`
	Name        string `db:"name"`
	Pieces      int    `db:"pieces"`
	Job         string `db:"job"`
	ItemLevel   *int   `db:"item_level"`
	RecipeLevel int    `db:"recipe_level"`
}

type ParentItem struct {
	Id   string `db:"parent_item_id"`
	Type string `db:"parent_item_type"`
	Name string `db:"parent_item_name"`
}

type Material struct {
	Id            string `db:"id"`
	RecipeId      string `db:"recipe_id"`
	ParentItemId  string `db:"parent_item_id"`
	ChildItemId   string `db:"child_item_id"`
	ChildItemName string `db:"child_item_name"`
	Quantity      int    `db:"quantity"`
	Type          string `db:"type"`
}
