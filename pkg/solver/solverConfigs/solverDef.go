package solverconfigs

// SolverTypeEnum will hold enumerate for supported solvers.
type SolverTypeEnum int

// Defined in SolverTypeEnum solvers.
const (
	Paging SolverTypeEnum = iota
	UpdateList
)

// String creates string from SolverTypeEnum
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
