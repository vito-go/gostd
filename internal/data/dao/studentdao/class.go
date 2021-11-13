package studentdao

type classRepo struct {
	db *studentDB
}

func (c *classRepo) TableName() string {
	return "class"
}
func NewClassRepo(db *studentDB) *classRepo {
	return &classRepo{db: db}
}

type ClassModel struct {
	ID       int64
	Number   int64
	Name     string
	Province string
	City     string
}
