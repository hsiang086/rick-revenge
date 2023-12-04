package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ItsJimi/gif/pkg/convert"
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"github.com/kkdai/youtube/v2"
)

func main() {
	entries, err := os.ReadDir(".\\video")
	if err != nil {
		panic(err)
	}
	if len(entries) == 0 {
		downloadVideo()
	}
	inpath := ".\\video"
	outpath := ".\\gif"
	convertMp4ToPng(inpath, outpath)
	inpath = ".\\gif\\*.gif"
	generateAscii(inpath)
}

func convertMp4ToPng(inpath string, outpath string) {
	options := convert.Options{
		FPS:   10,
		Scale: -1,
		Crop:  "",
	}

	err := convert.FromFolder(inpath, outpath, options)
	if err != nil {
		panic(err)
	}
}

func generateAscii(path string) {
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{100, 50}
	flags.Colored = true
	flags.SaveTxtPath = "."
	flags.SaveImagePath = "."
	flags.Braille = true
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	asciiArt, err := aic_package.Convert(path, flags)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", asciiArt)
}

func downloadVideo() {
	videoId := "dQw4w9WgXcQ"
	client := youtube.Client{}
	video, err := client.GetVideo(videoId)
	if err != nil {
		panic(err)
	}
	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	file, err := os.Create(".\\video\\rick.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
