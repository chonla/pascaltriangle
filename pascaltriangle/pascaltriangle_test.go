package pascaltriangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var PascalLine = PascalLineMemo

func TestPascalLine1(t *testing.T) {
	line := PascalLine(1)

	assert.Equal(t, []int{1}, line)
}

func TestPascalLine2(t *testing.T) {
	line := PascalLine(2)

	assert.Equal(t, []int{1, 1}, line)
}

func TestPascalLine3(t *testing.T) {
	line := PascalLine(3)

	assert.Equal(t, []int{1, 2, 1}, line)
}
