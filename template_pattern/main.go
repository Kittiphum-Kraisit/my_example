package main

import "fmt"

// Using Polymorphism
type DataExporter interface {
	OpenFile()
	FormatData()
	CloseFile()
}

func ExportData(d DataExporter) {
	d.OpenFile()
	d.FormatData()
	d.CloseFile()
}

type CSVExporter struct{}

func (ce *CSVExporter) OpenFile() {
	fmt.Println("Opening CSV file")
}

func (ce *CSVExporter) FormatData() {
	fmt.Println("Formatting data as CSV")
}

func (ce *CSVExporter) CloseFile() {
	fmt.Println("Closing CSV file")
}

type XMLExporter struct{}

func (xe *XMLExporter) OpenFile() {
	fmt.Println("Opening XML file")
}

func (xe *XMLExporter) FormatData() {
	fmt.Println("Formatting data as XML")
}

func (xe *XMLExporter) CloseFile() {
	fmt.Println("Closing XML file")
}

func main() {
	// Export data as CSV
	csvExporter := &CSVExporter{}
	ExportData(csvExporter)
	// Export data as XML
	xmlExporter := &XMLExporter{}
	ExportData(xmlExporter)
}
