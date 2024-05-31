package utils

import (
	"fmt"
	"os"
	"strconv"
)

func HandleVideo(source *string, start float64, end float64, height int, width int, avs string) {

	//check if source file exists
	if _, err := os.Stat(*source); os.IsNotExist(err) {
		panic("source file does not exist")
	}

	v, p := LoadVideo(*source)
	av, ap := LoadVideo(avs)

	d, e := strconv.ParseFloat(p.Streams[0].Duration, 32)
	ad, ae := strconv.ParseFloat(ap.Streams[0].Duration, 32)

	if end == 0 {
		end = d
	}

	if start > end {
		panic("start time is greater than end time")
	}

	if e != nil {
		panic(e)
	}

	if ae != nil {
		panic(ae)
	}

	if start > d || start > ad {
		panic("start time is greater than the duration of the video")
	}

	if end > d || end > ad {
		panic("end time is greater than the duration of the video")
	}

	t := SliceVideo(v, start, end)
	dd := end - start

	rs, re := RandomTimestamp(ad, dd)
	at := SliceVideo(av, rs, re)
	fmt.Println("additional video start: ", rs)
	fmt.Println("additional video end: ", re)

	mv := MergeVideos(t, at, height, width)

	o := TimestampFileName()
	fmt.Println("output file: ", o)
	ExportVideo(mv, o, start, dd)
}
