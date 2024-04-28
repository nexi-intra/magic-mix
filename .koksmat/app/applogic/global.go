package applogic

type Cell struct {
	Row    int
	Column int
	Value  string
}

type Row struct {
	Cells []Cell
}

type Sheet struct {
	Name string
	Rows []Row
}
