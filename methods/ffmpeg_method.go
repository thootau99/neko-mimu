package methods

import (
	"fmt"
	TextContent "github.com/thootau/neko-mimu/pkg/textContent"
	VideoContent "github.com/thootau/neko-mimu/pkg/videoContent"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"strings"

	"github.com/u2takey/go-utils/uuid"
)

func ImageToVideoWithSecond(imagePath string, second int) {
	mp4FilePath := strings.Replace(imagePath, ".jpg", ".mp4", -1)
	err := ffmpeg.Input(imagePath, ffmpeg.KwArgs{"loop": "1"}).Output(fmt.Sprint(mp4FilePath), ffmpeg.KwArgs{"t": fmt.Sprintf("%d", second)}).OverWriteOutput().ErrorToStdOut().Run()
	if nil != err {
		log.Fatal(err)
	}
}

func generateVideo(contents []VideoContent.VideoContent, subtitles []TextContent.TextContent) {
	validateResult := VideoContent.ValidateContent(contents)
	if !validateResult {
		log.Fatal("Invalid video content: Only one background and one layer one video content are allowed.")
	}
	fileUUID := uuid.NewUUID()
	filePath := fmt.Sprintf("./output/%s.mp4", fileUUID)
}

func OverlayVideoOnVideo(frontVideoPath string, backgroundVideoPath string, startSecond int) string {
	outputFileNameUUID := uuid.NewUUID()
	outputPath := fmt.Sprintf("./output/%s.mp4", outputFileNameUUID)
	overlay := ffmpeg.Input(frontVideoPath).Filter("colorkey", ffmpeg.Args{"0x4fff00:0.1:0.2"})
	err := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input(backgroundVideoPath),
			overlay,
		}, "overlay", ffmpeg.Args{"1:1"}, ffmpeg.KwArgs{"enable": "gte(t,1)"}).Output(outputPath).OverWriteOutput().ErrorToStdOut().Run()

	if nil != err {
		log.Fatal(err)
	}

	return outputPath
}

func PutTextOnVide(videoPath string, text string, x int, y int) {
	err := ffmpeg.Input(videoPath).Drawtext("割", 10, 10, true, ffmpeg.KwArgs{"fontfile": "/font/NotoSansCJKjp-Regular.otf", "fontsize": "240"}).Output(videoPath).OverWriteOutput().ErrorToStdOut().Run()

	if nil != err {
		log.Fatal(err)
	}
}
