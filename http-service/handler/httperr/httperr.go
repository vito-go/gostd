package httperr

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ErrInternal = Err("system error")
)
