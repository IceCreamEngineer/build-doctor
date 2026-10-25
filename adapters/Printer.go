package adapters

import "fmt"

type Printer struct{}

func (p Printer) Print(message string) {
	fmt.Printf(message)
}
