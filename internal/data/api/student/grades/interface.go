package grades

type Interface interface {
	GetGradesByNameAndId(id int64, name string)
	GetTotalGradesByID(id int64) int64
}
