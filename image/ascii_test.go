package image

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteWithNil(t *testing.T) {
	i, err := Execute(nil)
	assert.Nil(t, i)
	assert.Error(t, err)
}

func TestExecuteWithBadImage(t *testing.T) {
	f, _ := os.Open("../README.org")
	i, err := Execute(f)
	assert.Nil(t, i)
	assert.Error(t, err)
}

func TestConvertWithNilImage(t *testing.T) {
	i, err := convert(nil)
	assert.Nil(t, i)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "No image found")
}

func TestExecuteWithImage(t *testing.T) {
	f, _ := os.Open("../gopher.png")
	i, err := Execute(f)
	assert.NotNil(t, i)
	assert.Contains(t, i.Data, "@@@@@@@")
	assert.NoError(t, err)
}
