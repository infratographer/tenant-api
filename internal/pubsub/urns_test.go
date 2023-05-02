package pubsub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.infratographer.com/x/gidx"
)

func TestNewURN(t *testing.T) {
	tests := []struct {
		name string
		kind string
		id   gidx.PrefixedID
		want string
	}{
		{
			name: "example urn",
			kind: "testThing",
			id:   "abcdefg-abcd-abcdefg",
			want: "urn:infratographer:testThing:abcdefg-abcd-abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newURN(tt.kind, tt.id)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newURN(t *testing.T) {
	tests := []struct {
		name string
		kind string
		id   gidx.PrefixedID
		want string
	}{
		{
			name: "example",
			kind: "foo",
			id:   "abcdefg-abcd-abcdefg",
			want: "urn:infratographer:foo:abcdefg-abcd-abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newURN(tt.kind, tt.id)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewTenantURN(t *testing.T) {
	tests := []struct {
		name string
		id   gidx.PrefixedID
		want string
	}{
		{
			name: "example",
			id:   "abcdefg-abcd-abcdefg",
			want: "urn:infratographer:tenants:abcdefg-abcd-abcdefg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTenantURN(tt.id)
			assert.Equal(t, tt.want, got)
		})
	}
}
