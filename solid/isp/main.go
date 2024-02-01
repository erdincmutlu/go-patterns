package main

// Interface Segregation Principle
// You should not putting too much into an interface

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {

}
func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

// Don't force OldFashionedPrinter to implement new functionality such as Fax and Scan
func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

// Deprecated: ...
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {}
func (p Photocopier) Scan(d Document)  {}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// Decorator design pattern
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

	// Here Scan function call be striked out in the IDE as it is deprecated
	// ofp := OldFashionedPrinter{}
	// ofp.Scan()

}
