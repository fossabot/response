package operator

import (
	"os"
	"path"

	"github.com/dgraph-io/badger/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Operator is a management plane used for controlling Response instances.
type Operator struct {
	Path             string
	Data             *badger.DB
	Log              *zerolog.Logger
	PostgresURL      string
	PostgresHost     string
	PostgresPort     string
	PostgresUsername string
	PostgresPassword string
}

// OptionFn is used to configure the operator.
type OptionFn = func(operator *Operator) error

// New creates a new Operator to manage Response instances.
func New(opts ...OptionFn) (*Operator, error) {
	op := &Operator{}

	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get home directory")
	}
	op.Path = path.Join(homeDir, ".responseop")

	for _, fn := range opts {
		if err := fn(op); err != nil {
			return nil, errors.Wrapf(err, "unable to configure Operator")
		}
	}

	// setup log file and logger
	logfile, err := os.OpenFile(path.Join(op.Path, "operator.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open log file")
	}

	// check if operator exists in path, if not create it
	if _, err := os.Stat(op.Path); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(op.Path, 0700); err != nil {
				return nil, errors.Wrap(err, "unable to create operator at path")
			}
		} else {
			return nil, errors.Wrap(err, "unable to detect directory status")
		}
	}

	// init logger
	logger := zerolog.New(logfile)
	op.Log = &logger

	// initialize and open our kv store!
	badgerOptions := badger.DefaultOptions(path.Join(op.Path, "data"))
	badgerOptions.Logger = NewBadgerLogConverter(logger)

	data, err := badger.Open(badgerOptions)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open bdb")
	}
	op.Data = data

	return op, nil
}

// Close should be called when the Operator is no longer in-use, allowing cleanup of any open
// files and directories.
func (o *Operator) Close() {
	o.Data.Close()
}

// IsInitialized determines if the operator has been initialized within the given path, typically
// a directory within a User's homedir.
func (o *Operator) IsInitialized() bool {
	if _, err := os.Stat(path.Join(o.Path, "data")); err != nil {
		return false
	}

	return true
}