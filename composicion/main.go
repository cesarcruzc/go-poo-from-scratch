package main

import (
	"fmt"

	"github.com/cesarcruzc/go-poo-from-scratch/composicion/pkg/customer"
	"github.com/cesarcruzc/go-poo-from-scratch/composicion/pkg/invoice"
	"github.com/cesarcruzc/go-poo-from-scratch/composicion/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Bogotá",
		customer.New("César", "Calle 77B # 129 - 70", "3007871452"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(2, "Curso de POO con Go", 54.32),
			invoiceitem.New(3, "Curso de testing con Go", 90.00),
		),
	)

	i.SetTotal()

	fmt.Printf("%+v", i)

}
