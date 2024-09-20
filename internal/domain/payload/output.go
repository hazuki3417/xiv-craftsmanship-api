package payload

type Craft struct {
	RecipeId   string
	ItemId     string
	Name       string
	Job        string
	Pieces     int
	ItemLevel  int
	CraftLevel int
}
type Recipe struct {
	RecipeID  string
	ItemID    string
	Materials []Material
}

type Material struct {
	RecipeId string
	ItemID   string
	Quantity int
	Type     string
	Recipes  []Recipe
}
