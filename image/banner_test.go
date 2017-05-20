package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBannerExecute(t *testing.T) {
	i := Banner{Input: "test"}
	im, err := i.Execute()
	assert.NoError(t, err)
	assert.NotNil(t, im)
	assert.NotNil(t, im.Data)
}

// need to work out the test here for faliures at least happy path covers ~70%
// func TestBannerExecuteError(t *testing.T) {
// 	i := Banner{Input: ""}
// 	im, err := i.Execute()
// 	assert.Error(t, err)
// 	assert.Nil(t, im)
// 	// assert.NotNil(t, im)
// 	// assert.NotNil(t, im.Data)
// }
