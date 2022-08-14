package gwkerr

import "github.com/pkg/errors"

func ErrInvalidTaskParam(err error) error {
	return errors.Wrap(err, "invalid param")
}
