package image

import (
	"errors"
	"log"
	"os"
	"runtime"

	"github.com/cp16net/go-image2ascii/logger"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

const (
	// TempOutputFile is the location of the temp file.
	// ideally this should be an actual /tmp/file but had issues with
	// it for weird reasons so this will have to do for now.
	// MULTITHREADED WILL NOT WORK WITH THIS
	TempOutputFile = "local-file-output.png"
)

// Banner struct object
type Banner struct {
	fontLocation string
	Input        string
}

func (b *Banner) setup() error {
	//check os for font location
	if runtime.GOOS == "windows" {
		err := errors.New("sorry this wont work on windows [dont know where the fonts live]")
		return err
	} else if runtime.GOOS == "linux" {
		b.fontLocation = "/usr/share/fonts/truetype/freefont/FreeMonoBold.ttf"
	} else if runtime.GOOS == "darwin" {
		b.fontLocation = "/Library/Fonts/Courier New Bold.ttf"
	}
	return nil
}

// Execute method to create a string of text to a banner image.
// The idea for this is to create a banner of text for a dot matrix
// printer to print out.
func (b Banner) Execute() (*Image, error) {
	if err := b.setup(); err != nil {
		return nil, err
	}
	if err := b.createTextImage(); err != nil {
		return nil, err
	}

	// load file from path
	f, err := os.Open(TempOutputFile)
	if err != nil {
		logger.Logger.Error("**ERROR** file failed to load", err)
		return nil, err
	}
	defer f.Close()
	defer os.Remove(TempOutputFile)

	i := Image{}
	img, err := i.Execute(f, -1, -1)
	if err != nil {
		logger.Logger.Error("**ERROR** converting image", err)
		return nil, err
	}

	return img, nil
}

func (b Banner) createTextImage() error {
	const S = 1024
	dc := gg.NewContext(S, S)
	w, h := dc.MeasureString(b.Input)
	h = 80
	w = w * 7.5

	dc = gg.NewContext(int(w), int(h))
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(b.fontLocation, 80); err != nil {
		return err
	}
	dc.DrawStringAnchored(b.Input, w/2, h/2, 0.5, 0.35)

	dc.SavePNG(TempOutputFile)
	src, err := imaging.Open(TempOutputFile)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}
	dst := imaging.Rotate270(src)
	imaging.Save(dst, TempOutputFile)
	return nil
}
