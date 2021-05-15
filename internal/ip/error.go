package ip

import "fmt"

// NotSupported errors when a given provider is not supported.
type NotSupported struct {
	Name string
}

func (e *NotSupported) Error() string {
	return fmt.Sprintf("Provider %s not supported.", e.Name)
}
