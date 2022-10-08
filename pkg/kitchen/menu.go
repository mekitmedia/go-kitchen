package kitchen

type Menu struct {
	recipes []Recipe
}

func (m *Menu) AddRecipe(recipe Recipe) {
	m.recipes = append(m.recipes, recipe)
}
