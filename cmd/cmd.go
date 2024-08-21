package cmd

import (
	"fmt"
	"os"

	"github.com/lvm/album2go/internal/interfaces"
	"github.com/lvm/album2go/internal/usecases"
	"github.com/lvm/album2go/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	yt      *interfaces.YouTube
	ap      *interfaces.AudioProcessor
	uc      *usecases.TrackUsecase
	verbose bool
)

func init() {
	yt = interfaces.NewYouTubeClient()
	ap = interfaces.NewAudioProcessor()
	uc = usecases.NewTrackUsecase(ap)

	rootCmd.Flags().String("artist", "Unknown Artist", "Artist name")
	rootCmd.Flags().String("album", "Unknown Album", "Album name")
	rootCmd.Flags().String("tracklist", "", "Path to tracklist file")
	rootCmd.Flags().String("output", "", "Path to store audio files")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")

	logger.SetVerbose(verbose)
}

var rootCmd = &cobra.Command{
	Use:   "album2go",
	Short: "Downloads a video from YouTube and splits tracks based on a tracklist.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logger.Error("Usage: %s", fmt.Errorf("album2go <YouTube URL> --artist <artist> --album <album> --tracklist <tracklist> --output <directory> [--verbose]"))
			return
		}

		youtubeURL := args[0]
		tracklistFile, _ := cmd.Flags().GetString("tracklist")
		artist, _ := cmd.Flags().GetString("artist")
		album, _ := cmd.Flags().GetString("album")
		outputDir, _ := cmd.Flags().GetString("output")

		if err := run(youtubeURL, tracklistFile, artist, album, outputDir); err != nil {
			logger.Error("Error: %v", err)
		}
	},
}

func run(youtubeURL, tracklistFile, artist, album, outputDir string) error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	audioFilename := "audio.mp4"
	if err := yt.DownloadAudio(youtubeURL, audioFilename); err != nil {
		return fmt.Errorf("error downloading YouTube audio: %w", err)
	}
	logger.Info(fmt.Sprintf("YouTube video %s downloaded successfully!\n", youtubeURL))

	tracks, err := uc.ProcessTracklist(tracklistFile, artist, album)
	if err != nil {
		return fmt.Errorf("error processing tracklist: %w", err)
	}

	for _, track := range tracks {
		if err := uc.ProcessAudioFile(outputDir, audioFilename, track); err != nil {
			return fmt.Errorf("error slicing audio for track %d: %w", track.Num, err)
		}
	}

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}
