package main

import "fmt"

type Printer interface {
	print()
}

type LaserPrinter struct {
	serialnumber  string
	cartridgeType uint8
}

func (lp LaserPrinter) print() {
	fmt.Printf("serialnumber: [%v], cartridgeType: [%v]\n", lp.serialnumber, lp.cartridgeType)
}

type TankPrinter struct {
	serialnumber  string
	cartridgeType uint8
	capacityLtrs  int
}

func (tp TankPrinter) print() {
	fmt.Printf("serialnumber: [%v], cartridgeType: [%v], capacityLtrs: [%v]\n", tp.serialnumber, tp.cartridgeType, tp.capacityLtrs)
}

func printSpecs(p Printer) {
	p.print()
}

func main() {
	hpLasterPrinter := LaserPrinter{
		serialnumber:  "9182712",
		cartridgeType: uint8(1),
	}

	nikonInkJetPrinter := TankPrinter{
		serialnumber:  "685598",
		cartridgeType: uint8(2),
		capacityLtrs:  2,
	}

	printSpecs(hpLasterPrinter)
	printSpecs(nikonInkJetPrinter)
}
