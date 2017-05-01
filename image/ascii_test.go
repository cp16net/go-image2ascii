package image

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteWithNil(t *testing.T) {
	i := Image{}
	im, err := i.Execute(nil, -1, -1)
	assert.Nil(t, im)
	assert.Error(t, err)
}

func TestExecuteWithBadImage(t *testing.T) {
	f, _ := os.Open("../README.org")
	i := Image{}
	im, err := i.Execute(f, -1, -1)
	assert.Nil(t, im)
	assert.Error(t, err)
}

func TestConvertWithNilImage(t *testing.T) {
	i, err := convert(nil, -1, -1)
	assert.Nil(t, i)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "No image found")
}

func TestExecuteWithImage(t *testing.T) {
	f, _ := os.Open("../gopher.png")
	i := Image{}
	im, err := i.Execute(f, -1, -1)
	assert.NotNil(t, im)
	assert.Contains(t, im.Data, "@@@@@@@")
	assert.NoError(t, err)
}
