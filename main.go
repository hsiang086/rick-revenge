package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"github.com/kkdai/youtube/v2"
	"github.com/schollz/progressbar/v3"
)

func main() {
	entries, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}
	isVideoFile := false
	isTextFile := false
	isImageFile := false
	for _, entry := range entries {
		if entry.Name() == "video" {
			isVideoFile = true
		} else if entry.Name() == "images" {
			isImageFile = true
		} else if entry.Name() == "text" {
			isTextFile = true
		}
	}
	if !isVideoFile {
		os.Mkdir("./video", os.ModePerm)
	} else if !isImageFile {
		os.Mkdir("./images", os.ModePerm)
	} else if !isTextFile {
		os.Mkdir("./text", os.ModePerm)
	}
	isVideo, err := os.ReadDir("./video")
	if err != nil {
		panic(err)
	}
	if len(isVideo) == 0 {
		downloadVideo()
	}
	imagePath := "./images"
	videoPath := "./video/rick.mp4"
	isImage, err := os.ReadDir(imagePath)
	if err != nil {
		panic(err)
	}
	if len(isImage) == 0 {
		ffmpegConvert(videoPath, imagePath)
	}
	textPath := "./text"
	imageList, err := os.ReadDir(imagePath)
	if err != nil {
		panic(err)
	}
	bar := progressbar.Default(int64(len(imageList)))
	for _, image := range imageList {
		generateAscii(imagePath+"/"+image.Name(), textPath)
		bar.Add(1)
	}
}

func ffmpegConvert(inpath string, outpath string) {
	print("Converting video to images...")
	cmd := exec.Command("ffmpeg", "-threads 4", "-i", inpath, "-vf", "fps=20", outpath+"/%d.png")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	print("\tDone!\n")
}

func generateAscii(inpath string, outpath string) {
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{100, 50}
	flags.SaveTxtPath = outpath
	flags.Braille = true
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}
	_, err := aic_package.Convert(inpath, flags)
	if err != nil {
		panic(err)
	}
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
	file, err := os.Create("./video/rick.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
