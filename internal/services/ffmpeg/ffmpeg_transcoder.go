package ffmpeg

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func VideoTranscoder(videoFileKey string) error {

	cmd := exec.Command("ffmpeg", "-i", videoFileKey, "-b:v", "13000k", "13000.mp4")

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("getting stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("starting ffmpeg: %w", err)
	}

	go func() {
		buf := new(bytes.Buffer)
		io.Copy(buf, stderr)
		if buf.Len() > 0 {
			fmt.Println("FFmpeg stderr:", buf.String())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("running ffmpeg: %w", err)
	}

	return nil
}
