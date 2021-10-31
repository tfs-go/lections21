package lecture07

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt2Str_Testify(t *testing.T) {
	assert.Equal(t, "7", Int2Str(7))

	// assert.Equal(t, "10", Int2Str(0), "zero value")

	assert.ElementsMatch(t, []int{1, 2, 3}, []int{2, 3, 1})

	assert.InDelta(t, 7, 5+rand.Intn(4), 3)

	require.Equal(t, "7", Int2Str(7))
}
