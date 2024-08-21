package usecases

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lvm/album2go/internal/domain"
	"github.com/lvm/album2go/internal/interfaces"
	"github.com/lvm/album2go/internal/utils"
)

type (
	ITrackUsecase interface {
		ProcessTracklist(tracklistFile, artist, album string) ([]domain.Track, error)
		ProcessAudioFile(outputDir, filename string, track domain.Track) error
	}

	TrackUsecase struct {
		AudioProcessor interfaces.IAudioProcessor
	}
)

func NewTrackUsecase(audioProcessor interfaces.IAudioProcessor) *TrackUsecase {
	return &TrackUsecase{AudioProcessor: audioProcessor}
}

func (u *TrackUsecase) ProcessTracklist(tracklistFile, artist, album string) ([]domain.Track, error) {
	file, err := os.Open(tracklistFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tracks []domain.Track
	scanner := bufio.NewScanner(file)
	startTime := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	for scanner.Scan() {
		line := scanner.Text()

		track, err := utils.ParseTrack(artist, album, line, startTime)
		if err == nil {
			startTime = track.EndTime
			tracks = append(tracks, track)
		} else {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (u *TrackUsecase) ProcessAudioFile(outputDir, filename string, track domain.Track) error {
	outputFilename := fmt.Sprintf("%d_-_%s.mp3", track.Num, strings.ReplaceAll(track.Title, " ", "_"))
	outputPath := filepath.Join(outputDir, outputFilename)

	if err := u.AudioProcessor.Slice(filename, track.StartTime, track.EndTime, outputPath); err != nil {
		return fmt.Errorf("error slicing audio: %w", err)
	}

	if err := u.AudioProcessor.Tag(track.Num, track.Artist, track.Album, track.Title, outputPath); err != nil {
		return fmt.Errorf("error adding ID3 tags: %w", err)
	}

	return nil
}
