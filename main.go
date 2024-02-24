package main

import (
	"fmt"

	methods "github.com/thootau/neko-mimu/methods"
)

func main() {

	const _TEMPLATE_PATH = "./templates"
	const _OUTPUT_PATH = "./output"

	methods.OverlayVideoOnVideo(
		fmt.Sprintf("%s/banana.mp4", _TEMPLATE_PATH),
		fmt.Sprintf("%s/single_room.jpg", _TEMPLATE_PATH),
		1,
	)

}
