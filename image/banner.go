package image

import (
	"errors"
	"io"

	"github.com/cp16net/go-image2ascii/logger"
)

type Banner struct {
}

func (b Banner) Execute(f io.Reader) (*Image, error) {
	if f == nil {
		logger.Logger.Error("nil input given")
		return nil, errors.New("nill input given")
	}

	img := Image{}
	return &img, nil
}
