package main

import (
	"fmt"
	"time"

	methods "github.com/thootau/neko-mimu/methods"
)

func main() {

	const _TEMPLATE_PATH = "./templates"
	const _OUTPUT_PATH = "./output"
	methods.ImageToVideoWithSecond(fmt.Sprintf("%s/single_room.jpg", _TEMPLATE_PATH), 10)
	outputPath := methods.OverlayVideoOnVideo(
		fmt.Sprintf("%s/banana.mp4", _TEMPLATE_PATH),
		fmt.Sprintf("%s/single_room.mp4", _TEMPLATE_PATH),
		1,
	)
	time.Sleep(1000)
	methods.PutTextOnVide(outputPath, "割", 10, 10)

}
