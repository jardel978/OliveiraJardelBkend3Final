package errs

import "fmt"

type ErrRecordNotFound struct {
	Message string
}

func (e *ErrRecordNotFound) Error() string {
	return e.Message
}

type ErrInvalidMapping struct {
	Err error
}

func (i *ErrInvalidMapping) Error() string {
	return fmt.Sprintf("falha ao mapear valores. %v", i.Err.Error())
}

type ErrInvalidParams struct {
	Err error
}

func (i *ErrInvalidParams) Error() string {
	return fmt.Sprintf("erro de parâmetro. %v", i.Err.Error())
}

type ErrRequestBody struct {
	Err error
}

func (i *ErrRequestBody) Error() string {
	return fmt.Sprintf("erro com a request. %v", i.Err.Error())
}
