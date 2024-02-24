package videoContent

type VideoContent struct {
	ContentUri   string
	Second       int
	isBackground bool
	Layer        int
	Scale        float32
	Position     Position
}

type Position struct {
	x float32
	y float32
}

func GetBackgroundContentCount(videoContents []VideoContent) int {
	count := 0
	for _, videoContent := range videoContents {
		if videoContent.isBackground {
			count++
		}
	}
	return count
}

func GetLayerOneCount(videoContents []VideoContent) int {
	count := 0
	for _, videoContent := range videoContents {
		if videoContent.Layer == 1 {
			count++
		}
	}
	return count
}

func ValidateContent(videoContents []VideoContent) bool {
	return GetBackgroundContentCount(videoContents) == 1 && GetLayerOneCount(videoContents) == 1
}
