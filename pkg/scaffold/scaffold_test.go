package scaffold

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	err := New().Generate("./scaffold", "demo")
	require.Nil(t, err)
}
