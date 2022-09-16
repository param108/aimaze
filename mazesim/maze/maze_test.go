package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMaze(t *testing.T) {
	m := NewMaze()

	if err := m.Create(); err != nil {
		assert.Nil(t, err)
	}

}
