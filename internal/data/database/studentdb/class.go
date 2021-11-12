package studentdb

type classRepo struct {
	db *studentDB
}

func NewClassRepo(db *studentDB) *classRepo {
	return &classRepo{db: db}
}

const ClassTableName = "class"

type ClassModel struct {
	ID       int64
	Number   int64
	Name     string
	Province string
	City     string
}
