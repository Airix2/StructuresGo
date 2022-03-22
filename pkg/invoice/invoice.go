package invoice

import (
	"github.com/Airix2/StructuresGo/pkg/customer"
	"github.com/Airix2/StructuresGo/pkg/invoiceitem"
)

// Inovice is the struct of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// Set Total is the Setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
