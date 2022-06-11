package common

import (
	"fmt"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Ffm() {
	wd, _ := os.Getwd()
	path := wd + "\a.wmv"
	outpath := wd + "\a.mp4"
	err := ffmpeg.Input(path).
		Output(outpath, ffmpeg.KwArgs{"c:v": "libx264"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		fmt.Println(err)
	}
}
