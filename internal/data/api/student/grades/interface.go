package grades

type Interface interface {
	GetGradesByNameAndID(id int64, name string)
	GetTotalGradesByID(id int64) int64
}
