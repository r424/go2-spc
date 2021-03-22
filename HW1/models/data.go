package models

var DB map[int]Item

type Item struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}

func init() {
	DB = make(map[int]Item)
	DB[1] = Item{
		ID: 1,
		Title: "Сапоги",
		Amount: 3,
		Price: 10.90,
	}
}

func FindItemById(id int) (Item, bool) {
	item, found := DB[id]
	return item, found
}

func DeleteItemById(id int) bool {
	_, ok := DB[id]
	if ok {
		delete(DB, id)
		return true
	}
	return false
}

func SetItemById(id int, item Item) {
	DB[id] = item
}
