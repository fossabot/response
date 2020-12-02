package operator

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v2"
	"github.com/pkg/errors"
)

// InstanceState is the managed instance state.
type InstanceState int

// Represents the state of the Response instance.
const (
	InstanceStarted InstanceState = iota
	InstanceStopped
)

// ManagedInstance is an instance of Response managed by the Operator.
type ManagedInstance struct {
	Name  string        `json:"name"`
	Path  string        `json:"path"`
	State InstanceState `json:"state"`
}

func (o *Operator) getInstanceKey(key string) []byte {
	return []byte(fmt.Sprintf("managed-instance.%s", key))
}

func (o *Operator) createInstance(instance *ManagedInstance) error {
	instance.State = InstanceStopped

	return o.Data.Update(func(txn *badger.Txn) error {
		data, err := json.Marshal(txn)
		if err != nil {
			return errors.Wrap(err, "failed to marshal instance to bytes")
		}

		if err := txn.Set(o.getInstanceKey(instance.Name), data); err != nil {
			return errors.Wrap(err, "failed to set instance by key when creating")
		}

		return errors.Wrap(txn.Commit(), "failed to commit creation")
	})
}
