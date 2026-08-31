package twin

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestRuneWidth(t *testing.T) {
	assert.Equal(t, StyledRune{Rune: 'x', Style: Style{}}.Width(), 1)
	assert.Equal(t, StyledRune{Rune: '午', Style: Style{}}.Width(), 2)
}
