package main

import "fmt"

type DataSource interface {
	ReadData() string
}

type FileSource struct {
	name string
}

type NetworkSource struct {
	url string
}

func (d FileSource) ReadData() string {
	return "reading from file: " + d.name
}

func (n NetworkSource) ReadData() string {
	return "reading from network: " + n.url
}

func process(ds DataSource) {
	fmt.Println(ds.ReadData())
}

func main() {
	f := FileSource{name: "config.txt"}
	n := NetworkSource{url: "https://example.com"}

	process(f)
	process(n)
}
