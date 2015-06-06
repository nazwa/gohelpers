package tokens

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenLengths(t *testing.T) {
	assert.Len(t, <-Token4, 4)
	assert.Len(t, <-Token10, 10)
	assert.Len(t, <-Token24, 24)
	assert.Len(t, <-Token50, 50)
}
