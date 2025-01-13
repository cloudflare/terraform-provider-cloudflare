package apishieldoperation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpointValueEquality(t *testing.T) {
	e1 := NewEndpointValue("/foo/{fooId}/bar/{barId}/baz")
	e2 := NewEndpointValue("/foo/{var1}/bar/{var2}/baz")

	ok, diag := e1.StringSemanticEquals(context.Background(), e2)
	assert.Nil(t, diag)
	assert.True(t, ok)
}
