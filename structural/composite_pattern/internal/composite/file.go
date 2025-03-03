package composite

import "fmt"

type File struct {
	name string
	size int
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Print(indication, symbol, childPrefix string) {
	fmt.Printf("%s%s%s (%d bytes)\n", indication, symbol, f.name, f.size)
}
