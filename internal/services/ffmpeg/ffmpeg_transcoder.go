package ffmpeg

import "os/exec"

func VideoTranscoder(videoFileKey string) error {
	cmd := exec.Command("ffmpeg", "-i", videoFileKey, "-b:v", "13000k", "13000.mp4")
	err := cmd.Run()
	return err
}
