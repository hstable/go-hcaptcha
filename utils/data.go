package utils

import (
	"math/rand"
	"sync"
)

var (
	// FrameSize is the size of the hCaptcha frame.
	FrameSize = [2]int{400, 600}
	// TileImageSize is the size of the tile image.
	TileImageSize = [2]int{123, 123}
	// TileImageStartPosition is the start position of the tile image.
	TileImageStartPosition = [2]int{11, 130}
	// TileImagePadding is the padding between the tile images.
	TileImagePadding = [2]int{5, 6}
	// VerifyButtonPosition is the position of the verify button.
	VerifyButtonPosition = [2]int{314, 559}

	// TilesPerPage is the number of tiles per page.
	TilesPerPage = 9
	// TilesPerRow is the number of tiles per row.
	TilesPerRow = 3
)

// widgetCharacters are the characters used in randomly generated widget IDs.
var widgetCharacters = []rune("abcdefghijkmnopqrstuvwxyz0123456789")

// WidgetID generates a new random widget ID.
func WidgetID() string {
	b := make([]rune, Between(10, 12))
	for i := range b {
		b[i] = widgetCharacters[rand.Intn(len(widgetCharacters))]
	}
	return string(b)
}

var onceInitVersion sync.Once

// InitVersion initializes the version, asset version and agent data.
func InitVersion() {
	onceInitVersion.Do(func() {
		updateVersion()
		updateAssetVersion()
	})
}
