package payload

type Level struct {
	Item  *int
	Craft int
}

type Craft struct {
	Id     string
	ItemId string
	Name   string
	Job    string
	Pieces int
	Level  Level
}
type Recipe struct {
	RecipeID  string
	ItemID    string
	Materials []Material
}

type Material struct {
	recipeId string
	ItemID   string
	Quantity int
	Type     string
	Recipes  []Recipe
}
