package methods

import (
	"fmt"
	"log/slog"

	TextContent "github.com/thootau/neko-mimu/pkg/textContent"
	VideoContent "github.com/thootau/neko-mimu/pkg/videoContent"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"github.com/u2takey/go-utils/uuid"
)

func ImageToVideoWithSecond(imagePath string, second int) *ffmpeg.Stream {
	return ffmpeg.Input(imagePath, ffmpeg.KwArgs{"loop": "1", "t": fmt.Sprintf("%d", second)})
}

func generateVideo(contents []VideoContent.VideoContent, subtitles []TextContent.TextContent) string {
	validateResult := VideoContent.ValidateContent(contents)
	if !validateResult {
		slog.Error("Invalid video content: Content validation failed.")
	}
	fileUUID := uuid.NewUUID()
	filePath := fmt.Sprintf("./output/%s.mp4", fileUUID)

	return filePath
}

func OverlayVideoOnVideo(frontVideo *ffmpeg.Stream, backgroundVideo *ffmpeg.Stream, startSecond int) *ffmpeg.Stream {
	// outputFileNameUUID := uuid.NewUUID()
	// outputPath := fmt.Sprintf("./output/%s.mp4", outputFileNameUUID)
	overlay := frontVideo.Filter("colorkey", ffmpeg.Args{"0x4fff00:0.1:0.2"})
	return ffmpeg.Filter(
		[]*ffmpeg.Stream{
			backgroundVideo,
			overlay,
		}, "overlay", ffmpeg.Args{"1:1"}, ffmpeg.KwArgs{"enable": "gte(t,1)"})
}

func PutTextOnVide(video *ffmpeg.Stream, text string, x int, y int) {
	err := video.Drawtext("割", 10, 10, true, ffmpeg.KwArgs{"fontfile": "/font/NotoSansCJKjp-Regular.otf", "fontsize": "240"}).Output(videoPath).OverWriteOutput().ErrorToStdOut().Run()

	if nil != err {
		slog.Any("Error", err)
	}
}
