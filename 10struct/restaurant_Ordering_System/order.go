package main

type Order struct {
	orderID  int
	items    []string
	total    float64
	customer string
}

func (o *Order) Additem(item string) {
	o.items = append(o.items, item)
}
func (o *Order) RemoveItem(item string) {
	for i, value := range o.items {
		if value == item {
			o.items = append(o.items)
		}
	}
}

func main() {

}
