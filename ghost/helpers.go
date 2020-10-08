package ghost

import (
	"errors"
	"strings"
)

func beautifyError(res *errorsResponse) error {
	var contexts []string
	for _, e := range res.Errors {
		contexts = append(contexts, e.Context)
	}
	return errors.New(strings.Join(contexts, ", "))
}
