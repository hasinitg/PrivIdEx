package util

import "fmt"

type InvalidArgumentError struct {
	GivenNumberOfArgs,
	ExpectedNumberOfArgs int
}

func (err InvalidArgumentError) Error() string {
	return fmt.Sprintf("Invalid number of arguments: %v. Expected: %v", err.GivenNumberOfArgs, err.ExpectedNumberOfArgs)
}
