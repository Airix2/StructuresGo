package main

import (
	"fmt"

	"github.com/Airix2/StructuresGo/pkg/customer"
	"github.com/Airix2/StructuresGo/pkg/invoice"
	"github.com/Airix2/StructuresGo/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Paris",
		customer.New("Alejandro", "Zona Rio", "665123423"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Curso Go", 12.5),
			invoiceitem.New(2, "Poo Go", 15.5),
			invoiceitem.New(3, "DB con Go", 20.0),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
