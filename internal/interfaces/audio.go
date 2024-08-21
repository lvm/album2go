package interfaces

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/lvm/album2go/pkg/logger"
)

type (
	IAudioProcessor interface {
		Slice(inputFile string, startTime, endTime time.Time, outputFile string) error
		Tag(num int, artist, album, title, file string) error
	}

	AudioProcessor struct{}
)

func NewAudioProcessor() *AudioProcessor {
	return &AudioProcessor{}
}

func checkCommand(cmdName string) error {
	if runtime.GOOS == "windows" {
		cmdName += ".exe"
	}

	_, err := exec.LookPath(cmdName)
	return err
}

func (p *AudioProcessor) Slice(inputFile string, startTime, endTime time.Time, outputFile string) error {
	if err := checkCommand("ffmpeg"); err != nil {
		return fmt.Errorf("ffmpeg is not installed (or not in $PATH): %w", err)
	}

	cmd := exec.Command(
		"ffmpeg", "-v", "quiet", "-y", "-i", inputFile,
		"-ss", startTime.Format("15:04:05"), "-to", endTime.Format("15:04:05"),
		"-c:a", "libmp3lame", outputFile,
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error slicing audio with ffmpeg: %w", err)
	}
	return nil
}

func (p *AudioProcessor) Tag(num int, artist, album, title, file string) error {
	if err := checkCommand("id3v2"); err != nil {
		logger.Info(fmt.Sprintf("id3v2 is not installed (or not in $PATH): %s", err.Error()))
		return nil
	}

	cmd := exec.Command("id3v2", "-a", artist, "-A", album, "-t", title, "-T", strconv.Itoa(num), file)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error tagging file with id3v2: %w", err)
	}
	return nil
}
