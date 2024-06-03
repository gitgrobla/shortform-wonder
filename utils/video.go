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

	if start > d {
		panic("start time is greater than the duration of the video")
	}

	if end > d {
		panic("end time is greater than the duration of the video")
	}

	dd := end - start
	rs, re := RandomTimestamp(ad, dd)
	aad := re - rs

	ExportTempVideo(v, true, start, dd)
	ExportTempVideo(av, false, rs, aad)

	mv := MergeVideos(height, width)

	o := TimestampFileName()
	fmt.Println("output file: ", o)
	ExportVideo(mv, o, start, dd)
}
