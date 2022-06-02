package common

import (
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Ffm() {
	err := ffmpeg.Input("a.wmv").
		Output("a.mp4", ffmpeg.KwArgs{"c:v": "libx264"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		fmt.Println(err)
	}
}
