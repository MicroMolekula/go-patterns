package composite

import "fmt"

type Folder struct {
	children []Component
	name     string
}

func NewFolder(name string, children []Component) *Folder {
	if children == nil {
		children = make([]Component, 0)
	}
	return &Folder{
		children: children,
		name:     name,
	}
}

func (f *Folder) Add(c Component) {
	f.children = append(f.children, c)
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) GetSize() int {
	size := 0
	for _, child := range f.children {
		size += child.GetSize()
	}
	return size
}

func (f *Folder) Print(indication, symbol, childPrefix string) {
	fmt.Printf("%s%s%s (%d bytes)\n", indication, symbol, f.name, f.GetSize())
	for i, child := range f.children {
		switch i {
		case len(f.children) - 1:
			child.Print(indication+childPrefix+"    ", "└───", "")
		default:
			child.Print(indication+childPrefix+"    ", "├───", "|")
		}
	}
}
