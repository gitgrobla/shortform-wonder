package utils

import (
	"strconv"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func SliceVideo(video *ffmpeg_go.Stream, start float64, end float64) *ffmpeg_go.Stream {
	return video.Trim(ffmpeg_go.KwArgs{"start": start, "end": end}).SetPts("PTS-STARTPTS").Get("v")
	// return video.Trim(ffmpeg_go.KwArgs{"start": start, "end": end}).SetPts("PTS-STARTPTS").Get("v")
}

func SliceAudio(audio *ffmpeg_go.Stream, start float64, end float64) *ffmpeg_go.Stream {
	return audio.Filter("atrim", ffmpeg_go.Args{}, ffmpeg_go.KwArgs{"start": start, "end": end}).Filter("asetpts", ffmpeg_go.Args{"PTS-STARTPTS"})
}

func MergeVideos(v0 *ffmpeg_go.Stream, v2 *ffmpeg_go.Stream, height int, width int) *ffmpeg_go.Stream {
	scaledHeight := height / 2
	widthStr := strconv.Itoa(width)
	heightStr := strconv.Itoa(scaledHeight)

	v1Scaled := v0.Filter("scale", ffmpeg_go.Args{widthStr, heightStr})
	v2Scaled := v2.Filter("scale", ffmpeg_go.Args{widthStr, heightStr})

	mv := ffmpeg_go.Filter([]*ffmpeg_go.Stream{v1Scaled, v2Scaled}, "vstack", ffmpeg_go.Args{})

	return mv
}
