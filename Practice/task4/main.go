package main

import "fmt"

type Storage interface {
	Save(data string)
}

type FileStorage struct{}

type MemoryStorage struct{}

func (f *FileStorage) Save(data string) {
	fmt.Println("saving to file:", data)
}

func (m *MemoryStorage) Save(data string) {
	fmt.Println("saving to memory:", data)
}

func NewStorage(kind string) Storage {
	if kind == "file" {
		return &FileStorage{}
	}
	return &MemoryStorage{}
}

func main() {
	s1 := NewStorage("file")
	s2 := NewStorage("mem")

	s1.Save("hello")
	s2.Save("world")
}
