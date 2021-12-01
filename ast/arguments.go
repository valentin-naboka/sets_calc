package ast

type File struct {
	Name string
}

func (f *File) Print() {
	print(f.Name)
}

type Int uint64

func (i *Int) Print() {
	print(i)
}
