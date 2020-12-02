package operator

import (
	"os"

	"github.com/pkg/errors"
)

// CreateLocalOptions is used to configure a local Response instance through the operator and
// is provided to CreateLocal as the second argument.
type CreateLocalOptions struct {
	Path            string
	EncryptionKey   string
	EnableDeveloper bool
}

// CreateLocal creates a local Response instance on this machine using the provided options.
func (o *Operator) CreateLocal(name string, opts *CreateLocalOptions) error {
	if err := os.MkdirAll(opts.Path, 0700); err != nil {
		return errors.Wrap(err, "unable to create instance directory")
	}

	return nil
}
