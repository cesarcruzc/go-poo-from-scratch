package invoiceitem

//Item contains of information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

//New returns a new Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

//Value getter of Item.value
func (i Item) Value() float64 {
	return i.value
}

//Items new type of Item struct
type Items []Item

//NewItems constructor for Items type
func NewItems(items ...Item) Items {
	var is Items

	for _, item := range items {
		is = append(is, item)
	}

	return is
}

//Total it is a method of Items type
func (i Items) Total() float64 {
	var total float64
	for _, item := range i {
		total += item.value
	}

	return total
}
