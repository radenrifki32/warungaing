package web

type ErrorNotFounded struct {
	Error string
}

func ResponseErrorNotFound(error string) ErrorNotFounded {
	return ErrorNotFounded{Error: error}
}
