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
	RecipeId  string
	Pieces    int
	ItemId    string
	Materials []Material
}

type Material struct {
	ItemId   string
	ItemName string
	Quantity int
	Type     string
	Recipes  []Recipe
}
