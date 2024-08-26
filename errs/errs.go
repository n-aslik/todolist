package errs

import "errors"

var (
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrTasksNotFound               = errors.New("ErrTasksNotFound")
	ErrIncorrectUsernameorPassword = errors.New("ErrIncorrectUsernameorPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
)
