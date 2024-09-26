package lessoncodebasic

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

//CopyParent(p *Parent) Parent, которая
//создает копию структуры Parent и возвращает ее:
func CopyParent(p *Parent) Parent {
	par := *p
	par.Children = make([]Child, len(p.Children))
	copy(par.Children, p.Children)
	return par
}
