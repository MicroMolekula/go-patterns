package composite

type Component interface {
	GetName() string
	GetSize() int
	Print(string, string, string)
}
