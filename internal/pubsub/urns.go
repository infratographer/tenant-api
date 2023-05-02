package pubsub

import (
	"fmt"

	"go.infratographer.com/x/gidx"
)

func newURN(kind string, id gidx.PrefixedID) string {
	return fmt.Sprintf("urn:infratographer:%s:%s", kind, id)
}

// NewTenantURN creates a new tenant URN
func NewTenantURN(id gidx.PrefixedID) string {
	return newURN("tenants", id)
}
