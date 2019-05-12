package contract

type Transition struct {
	Name string
	Params []Field
}

type Field struct {
	Name string
	Type string
}
