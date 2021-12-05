package main

type Document struct {
}

type Machine interface {	
	Print(d Document)	
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {
	println("Printing")	
}

func (m *MultiFunctionPrinter) Fax(d Document) {
	println("Faxing")	
}

func (m *MultiFunctionPrinter) Scan(d Document) {
	println("Scanning")	
}

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d Document) {
	println("Printing")	
}

func (o *OldFashionedPrinter) Fax(d Document) {
	panic("Faxing is not supported")
}

func (o *OldFashionedPrinter) Scan(d Document) {
	panic("Scanning is not supported")
}

// ISP - Interface Segregation Principle
// So the interface segregation principle basically states 
// that try to break up an interface into separate
// parts that people will definitely need

type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}
func (m *MyPrinter) Print(d Document) {
	println("Printing")	
}
type Photocopier struct {
}

func (p *Photocopier) Scan(d Document) {
	panic("implement me")
}

func (p *Photocopier) Print(d Document) {
	panic("implement me")
}

type MultiFunctionDevice struct {
	Printer
	Scanner
}

// decorator 
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}
func (m *MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m *MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	ofp := OldFashionedPrinter{}
	ofp.Print(Document{})
	ofp.Fax(Document{})	
}