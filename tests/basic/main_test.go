package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test: from inside the basic folder

func TestAddOne(t *testing.T) {
	result := AddOne(3)
	assert.Equal(t, 4, result, "AddOne should return the number plus one")
}
