package grades

type Interface interface {
	GetGradesByNameAndId(id int64, name string)
	GetTotalGradesById(id int64)int64
}
