package express

// Express 快递柜相关接口
type Express struct {
	GetPackage *GetPackage
}

func NewExpress(getPackage *GetPackage) *Express {
	return &Express{GetPackage: getPackage}
}

const ExpressID = `1000000`
