package videoContent

import (
	"log/slog"
	"time"
)

type VideoContent struct {
	ContentUri   string
	StartTime    time.Time
	EndTime      time.Time
	CutStartTime time.Time
	CurEndTime   time.Time
	Duration     int
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

// ValidateContent TODO: Could be improved by adding more validation rules
func ValidateContent(videoContents []VideoContent) bool {
	backgrounds := GetVideoBackground(videoContents)
	if len(backgrounds) < 1 {
		slog.Warn("Missing background content, will use default background")
	}
	if GetLayerOneCount(videoContents) > 1 {
		slog.Warn("Multiple layer one content detected, some content might not be visible")
	}
	if IsBackgroundOverlapping(videoContents) {
		slog.Warn("Backgrounds overlapping detected, some content might not be visible")
	}
	return GetBackgroundContentCount(videoContents) == 1 && GetLayerOneCount(videoContents) == 1
}

func (videoContent *VideoContent) SetBackground() {
	videoContent.isBackground = true
}

func GetVideoBackground(videoContents []VideoContent) []VideoContent {
	var backgrounds []VideoContent
	for _, videoContent := range videoContents {
		if videoContent.isBackground {
			backgrounds = append(backgrounds, videoContent)
		}
	}
	return backgrounds
}

func IsBackgroundOverlapping(videoContents []VideoContent) bool {
	backgrounds := GetVideoBackground(videoContents)
	// Calculate StartTime and duration to check if they are overlapping on editor (StartTime and EndTime)
	for i := 0; i < len(backgrounds); i++ {
		for j := i + 1; j < len(backgrounds); j++ {
			if backgrounds[i].StartTime.Before(backgrounds[j].EndTime) && backgrounds[i].EndTime.After(backgrounds[j].StartTime) {
				return true
			}
		}
	}
	return false
}
