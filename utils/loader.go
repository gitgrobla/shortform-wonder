package utils

//load the video file
//return a pointer to the video file
//so that it can be modified in the main function

import (
	"encoding/json"
	"fmt"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type _Stream struct {
	Index    int    `json:"index"`
	Duration string `json:"duration"`
}

type _Probe struct {
	Streams []_Stream `json:"streams"`
}

// LoadVideo => stream, duration
func LoadVideo(source string) (*ffmpeg_go.Stream, _Probe) {
	p, e := ffmpeg_go.Probe(source)

	if e != nil {
		panic(e)
	}

	_probe := &_Probe{}
	e = json.Unmarshal([]byte(p), _probe)

	if e != nil {
		panic(e)
	}

	return ffmpeg_go.Input(source), *_probe
}

func GetRandomAdditionalSource() string {
	files, e := os.ReadDir("attention\\")
	if e != nil {
		panic(e)
	}

	fmt.Println(files)

	files = FilterFiles(files)

	if len(files) == 0 {
		panic("no files found in ./attention/")
	}

	return "attention\\" + files[RandomNumber(0, len(files)-1)].Name()
}

func GetInputFiles() []os.DirEntry {
	files, e := os.ReadDir("input\\")
	if e != nil {
		panic(e)
	}

	return FilterFiles(files)
}

func CreateIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
}

func ExportVideo(video *ffmpeg_go.Stream, output string, s float64, d float64) {

	CreateIfNotExist("temp")
	err0 := ffmpeg_go.Output([]*ffmpeg_go.Stream{video}, "temp\\audio.mp3", ffmpeg_go.KwArgs{
		"map": "0:a",
		"ss":  s,
		"t":   d,
	}).OverWriteOutput().Run()

	if err0 != nil {
		panic(err0)
	}

	p := ExecuteTranscription()

	_vs := video.Filter("subtitles", ffmpeg_go.Args{p})

	if !FileExists(p) {
		Cleanup()
		panic("srt file does not exist")
	}

	o := ffmpeg_go.Output([]*ffmpeg_go.Stream{_vs}, output, ffmpeg_go.KwArgs{
		"map":    "0",
		"ss":     s,
		"t":      d,
		"c:v":    "h264",
		"c:a":    "aac",
		"strict": "experimental",
	})

	err := o.Run()

	if err != nil {
		Cleanup()
		panic(err)
	}

	Cleanup()
}

func FilterFiles(files []os.DirEntry) []os.DirEntry {
	var filteredFiles []os.DirEntry
	ext := ".gitkeep"

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if len(file.Name()) < len(ext) {
			continue
		}

		if file.Name()[len(file.Name())-len(ext):] != ext {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles
}

func Cleanup() {
	os.Remove("temp\\audio.mp3")
	os.Remove("temp\\audio.srt")
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}
