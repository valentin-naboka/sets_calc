package ast

type File struct {
	Name string
}

//TODO: remove
func (f *File) Print() {
	print(f.Name)
}

type Int uint64

//TODO: remove
func (i *Int) Print() {
	print(i)
}
