package teacherdao

type infoRepo struct {
	db *teacherDB
}

func NewInfoRepo(db *teacherDB) *infoRepo {
	return &infoRepo{db: db}
}
