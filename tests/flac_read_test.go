package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestFLACRead(t *testing.T) {
	asrt := assert.New(t)

	flac, err := tag.ReadFile("BeeMoved.flac")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}
	defer flac.Close()

	title, err := flac.GetTitle()
	asrt.NoError(err)
	asrt.Equal("Bee Moved", title)

	artist, err := flac.GetArtist()
	asrt.NoError(err)
	asrt.Equal("Blue Monday FM", artist)

	album, err := flac.GetAlbum()
	asrt.NoError(err)
	asrt.Equal("Bee Moved", album)

	albumArtist, err := flac.GetAlbumArtist()
	asrt.NoError(err)
	asrt.Equal("Blue Monday FM", albumArtist)

	picture, err := flac.GetPicture()
	asrt.NoError(err)
	out, err := ioutil.TempFile("", "flacTst.png")
	asrt.NoError(err)
	defer os.Remove(out.Name())
	err = png.Encode(out, picture)
	asrt.NoError(err)
	cmp := compareFiles("flac.png", out.Name())
	asrt.Equal(true, cmp)
}
