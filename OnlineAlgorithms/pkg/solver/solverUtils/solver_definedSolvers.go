package solverutils

type SolverTypeEnum int

const (
	Paging SolverTypeEnum = iota
	UpdateList
)

func (e SolverTypeEnum) String() string {
	switch e {
	case Paging:
		return "Paging"
	case UpdateList:
		return "UpdateList"
	default:
		return "NULL"
	}

}
