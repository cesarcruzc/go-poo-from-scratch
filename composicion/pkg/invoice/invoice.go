package invoice

import (
	"github.com/cesarcruzc/go-poo-from-scratch/composicion/pkg/customer"
	"github.com/cesarcruzc/go-poo-from-scratch/composicion/pkg/invoiceitem"
)

//Invoice is the structure of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer // 1-1
	// items   []invoiceitem.Item // 1-many
	items invoiceitem.Items
}

//New return a new Invoice
func New(
	country,
	city string,
	client customer.Customer,
	// items []invoiceitem.Item,
	items invoiceitem.Items,
) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

//SetTotal is the setter of Invoice.total
/*
	func (i *Invoice) SetTotal() {
		for _, item := range i.items {
			i.total += item.Value()
		}
	}
*/

//SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
