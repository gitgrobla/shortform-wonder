package utils

import (
	"os/exec"
)

func ExecuteTranscription() string {
	cmd := exec.Command("python", "python\\scribe.py")

	// Start the Python script
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Wait for the Python script to finish
	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	return "temp/audio.srt"
}
