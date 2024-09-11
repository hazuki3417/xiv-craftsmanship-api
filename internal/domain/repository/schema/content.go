package schema

type Craft struct {
	Id          string `db:"recipe_id"`
	Name        string `db:"name"`
	Pieces      int    `db:"pieces"`
	Job         string `db:"job"`
	ItemLevel   *int   `db:"item_level"`
	RecipeLevel int    `db:"recipe_level"`
}

type Material struct {
	TreeId           string `db:"tree_id"`
	ParentItemId     string `db:"parent_item_id"`
	ParentItemName   string `db:"parent_item_name"`
	ParentCraftLevel int    `db:"parent_craft_level"`
	ParentCraftJob   string `db:"parent_craft_job"`
	ChildItemId      string `db:"child_item_id"`
	ChildItemName    string `db:"child_item_name"`
	ChildItemType    string `db:"child_item_type"`
	Unit             int    `db:"unit"`
	Total            int    `db:"total"`
}
