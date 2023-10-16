package model

type InfoFruit struct {
	Fruit         string
	Calories      string
	Fats          string
	Sugar         string
	Carbohydrates string
	Protein       string
}

func Fruits() map[string]string {
	fruits := make(map[string]string)
	fruits["Яблоко"] = "apiFruit apple яблок"
	fruits["Банан"] = "apiFruit Banana бананов"
	fruits["Апульсин"] = "apiFruit Orange апельсинов"
	fruits["Дыня"] = "apiFruit Melon дыни"
	fruits["Лимон"] = "apiFruit Lemon лимона"
	fruits["Лайм"] = "apiFruit Lime лайма"
	fruits["Гранат"] = "apiFruit Pomegranate граната"
	fruits["Виноград"] = "apiFruit Grape винограда"
	fruits["Авокадо"] = "apiFruit Avocado авокадо"
	return fruits
}
