package main

import (
	interfaces "github.com/thootau/neko-mimu/interfaces"
	"github.com/thootau/neko-mimu/methods"
)

func main() {

	const _TEMPLATE_PATH = "./templates"
	const _OUTPUT_PATH = "./output"

	testVideoContents := interfaces.GenerateVideoRequest{
		Contents: []interfaces.VideoContent{
			interfaces.VideoContent{
				ContentUri:   "./templates/rinkoy.mp4",
				StartSecond:  0,
				EndSecond:    8,
				IsBackground: false,
				Layer:        0,
				Scale:        1,
				Position:     interfaces.Position{X: 100.0, Y: 100.0},
			},
		},
		Background: interfaces.VideoContent{
			ContentUri:   "./templates/single_room.jpg",
			StartSecond:  0,
			EndSecond:    8,
			IsBackground: true,
			Layer:        0,
			Scale:        1,
			Position:     interfaces.Position{X: 100.0, Y: 100.0},
		},
	}

	methods.GenerateVideo(&testVideoContents)

}
