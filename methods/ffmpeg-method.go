package methods

import (
	"fmt"

	interfaces "github.com/thootau/neko-mimu/interfaces"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"github.com/u2takey/go-utils/uuid"
)

func backgroundImageToVideoWithSecond(imagePath string, second int) *ffmpeg.Stream {
	return ffmpeg.Input(imagePath, ffmpeg.KwArgs{"loop": "1", "t": fmt.Sprintf("%d", second)})
}

func GenerateVideo(request *interfaces.GenerateVideoRequest) string {
	// validateResult := VideoContent.ValidateContent(contents)
	// if !validateResult {
	// slog.Error("Invalid video content: Content validation failed.")
	// }
	backgroundStream := backgroundImageToVideoWithSecond(request.Background.ContentUri, request.Background.EndSecond-request.Background.StartSecond)
	for _, videoContent := range request.Contents {
		backgroundStream = OverlayVideoOnVideo(&videoContent, backgroundStream)
	}

	fileUUID := uuid.NewUUID()
	filePath := fmt.Sprintf("./output/%s.mp4", fileUUID)

	backgroundStream.OverWriteOutput().Output(filePath).Run()
	return filePath
}

func OverlayVideoOnVideo(videoContent *interfaces.VideoContent, backgroundVideoStream *ffmpeg.Stream) *ffmpeg.Stream {
	// outputFileNameUUID := uuid.NewUUID()
	// outputPath := fmt.Sprintf("./output/%s.mp4", outputFileNameUUID)
	overlay := ffmpeg.Input(videoContent.ContentUri).Filter("colorkey", ffmpeg.Args{"0x4fff00:0.1:0.2"})
	return ffmpeg.Filter(
		[]*ffmpeg.Stream{
			backgroundVideoStream,
			overlay,
		}, "overlay", ffmpeg.Args{
			fmt.Sprintf("%f:%f", videoContent.Position.X, videoContent.Position.Y), // 素材が置くべき座標
		}, ffmpeg.KwArgs{"enable": fmt.Sprintf("gte(t,%d)", videoContent.StartSecond)})
}

func PutTextOnVide(video *ffmpeg.Stream, text string, x int, y int) *ffmpeg.Stream {
	return video.Drawtext("割", 10, 10, true, ffmpeg.KwArgs{"fontfile": "/font/NotoSansCJKjp-Regular.otf", "fontsize": "240"})
}
