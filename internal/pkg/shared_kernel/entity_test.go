package shared_kernel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	t.Parallel()

	id := NewID()
	assert.NotNil(t, id)
}

func TestStringToID(t *testing.T) {
	t.Parallel()

	_, err := StringToID("fd14c028-5f56-488a-8c29-3186fd62395c")
	assert.Nil(t, err)
}
