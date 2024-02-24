package methods

import (
	"fmt"
	"log"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ImageToVideoWithSecond(imagePath string, second int) {
	mp4FilePath := strings.Replace(imagePath, ".jpg", ".mp4", -1)
	err := ffmpeg.Input(imagePath, ffmpeg.KwArgs{"loop": "1"}).Output(fmt.Sprint(mp4FilePath), ffmpeg.KwArgs{"t": fmt.Sprintf("%d", second)}).OverWriteOutput().ErrorToStdOut().Run()
	if nil != err {
		log.Fatal(err)
	}
}

func OverlayVideoOnVideo(frontVideoPath string, backgroundVideoPath string, startSecond int) {
	overlay := ffmpeg.Input(frontVideoPath).Filter("colorkey", ffmpeg.Args{"0x4fff00:0.1:0.2"})
	err := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input(backgroundVideoPath),
			overlay,
		}, "overlay", ffmpeg.Args{"1:1"}, ffmpeg.KwArgs{"enable": "gte(t,1)"}).
		Output("./output/out1.mp4").OverWriteOutput().ErrorToStdOut().Run()
	if nil != err {
		log.Fatal(err)
	}
}
