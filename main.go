package main

import (
	"flag"
	"fmt"
	"os"
	"sfw/utils"
)

func main() {
	// run subcommand
	run_cmd := flag.NewFlagSet("run", flag.ExitOnError)

	// run subcommand flags
	run_source := run_cmd.String("source", "", "source of the video file")

	// timestamps
	run_start := run_cmd.String("start", "00:00", "start time")
	run_end := run_cmd.String("end", "00:00", "end time")
	run_height := run_cmd.Int("height", 1920, "height of the video")
	run_width := run_cmd.Int("width", 1080, "width of the video")
	run_additional := run_cmd.String("additional", "", "additional video")
	run_random := run_cmd.Bool("random", true, "random video")

	if len(os.Args) < 2 {
		fmt.Println("subcommand 'run' is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		handleRun(run_cmd, run_source, run_start, run_end, run_height, run_width, run_additional, run_random)
	case "bulk":
		handleBulkRun(run_cmd, run_source, run_start, run_end, run_height, run_width, run_additional, run_random)
	default:
		fmt.Println("subcommand not found")
		os.Exit(1)
	}

}

func handleBulkRun(run_cmd *flag.FlagSet, source *string, start *string, end *string, height *int, width *int, additional *string, random *bool) {
	run_cmd.Parse(os.Args[2:])

	files := utils.GetInputFiles()

	if *height <= 0 {
		fmt.Println("valid height is required")
		os.Exit(1)
	}

	if *width <= 0 {
		fmt.Println("valid width is required")
		os.Exit(1)
	}

	if *additional == "" && !*random {
		fmt.Println("additional video is required or use -random flag")
		os.Exit(1)
	}

	for _, file := range files {
		source := "input\\" + file.Name()

		var av string

		if *random {
			av = utils.GetRandomAdditionalSource()
		} else {
			av = *additional
		}

		handleRun(run_cmd, &source, start, end, height, width, &av, random)
	}

}

func handleRun(run_cmd *flag.FlagSet, source *string, start *string, end *string, height *int, width *int, additional *string, random *bool) {
	run_cmd.Parse(os.Args[2:])

	if *source == "" {
		fmt.Println("source is required")
		os.Exit(1)
	}

	if *height <= 0 {
		fmt.Println("valid height is required")
		os.Exit(1)
	}

	if *width <= 0 {
		fmt.Println("valid width is required")
		os.Exit(1)
	}

	var av string

	if *additional == "" && !*random {
		fmt.Println("additional video is required or use -random flag")
		os.Exit(1)
	}

	if *random {
		av = utils.GetRandomAdditionalSource()
	} else {
		av = *additional
	}

	sValid := utils.ValidateTimestamp(*start)
	eValid := utils.ValidateTimestamp(*end)

	if !sValid {
		fmt.Println("start time is invalid")
		os.Exit(1)
	}

	if !eValid {
		fmt.Println("end time is invalid")
		os.Exit(1)
	}

	utils.HandleVideo(source, utils.ConvertToSeconds(*start), utils.ConvertToSeconds(*end), *height, *width, av)
}
