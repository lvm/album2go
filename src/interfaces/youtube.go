package interfaces

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
)

type (
	IYouTube interface {
		DownloadAudio(url, filename string) error
	}

	YouTube struct {
		Client youtube.Client
	}
)

func NewYouTubeClient() *YouTube {
	return &YouTube{Client: youtube.Client{}}
}

func (yt *YouTube) GetVideo(url string) (*youtube.Video, error) {
	video, err := yt.Client.GetVideo(url)
	if err != nil {
		return nil, fmt.Errorf("could't retrieve youtube content: %w", err)
	}
	return video, nil
}

func (yt *YouTube) GetStream(video *youtube.Video, format *youtube.Format) (io.ReadCloser, error) {
	stream, _, err := yt.Client.GetStream(video, format)
	if err != nil {
		return nil, fmt.Errorf("no valid audio stream: %w", err)
	}
	return stream, nil
}

func (yt *YouTube) DownloadAudio(url, filename string) error {
	video, err := yt.GetVideo(url)
	if err != nil {
		return err
	}

	audioFormats := video.Formats.Type("audio")
	if len(audioFormats) == 0 {
		return fmt.Errorf("no audio formats available")
	}

	var bestFormat *youtube.Format
	for _, format := range audioFormats {
		if bestFormat == nil || format.Quality > bestFormat.Quality {
			bestFormat = &format
		}
	}

	if bestFormat == nil {
		return fmt.Errorf("no valid audio format found")
	}

	stream, err := yt.GetStream(video, bestFormat)
	if err != nil {
		return fmt.Errorf("no valid audio stream: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("couldn't create youtube audio file: %w", err)
	}
	defer file.Close()

	_, err = file.ReadFrom(stream)
	return err
}
