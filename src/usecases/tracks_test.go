package usecases

import (
	"testing"

	"github.com/lvm/album2go/src/interfaces"

	"os"

	"github.com/stretchr/testify/assert"
)

var (
	artist string = "Cattle Decapitation"
	album  string = "Death Atlas"
)

func TestProcessTracklist(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testtracklist*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	mockData := `1. Anthropogenic End Transmission 2:16
2. The Geocide 3:43
3. Be Still Our Bleeding Hearts 3:54
4. Vulturous 4:59
5. The Great Dying 1:13
6. One Day Closer To The End Of The World 3:47
7. Bring Back The Plague 4:28
8. Absolute Destitute 4:36
9. The Great Dying II 1:06
10. Finish Them 2:57
11. With All Disrespect 4:31
12. Time's Cruel Curtain 5:32
13. The Unerasable Past 2:51
14. Death Atlas 9:15
`
	expectedTitles := []string{
		"Anthropogenic End Transmission",
		"The Geocide",
		"Be Still Our Bleeding Hearts",
		"Vulturous",
		"The Great Dying",
		"One Day Closer To The End Of The World",
		"Bring Back The Plague",
		"Absolute Destitute",
		"The Great Dying II",
		"Finish Them",
		"With All Disrespect",
		"Time's Cruel Curtain",
		"The Unerasable Past",
		"Death Atlas",
	}

	_, err = tempFile.Write([]byte(mockData))
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	ap := &interfaces.AudioProcessor{}
	uc := NewTrackUsecase(ap)

	tracks, err := uc.ProcessTracklist(tempFile.Name(), artist, album)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Validate the result
	assert.Equal(t, 14, len(tracks), "expected 14 tracks")
	for i, track := range tracks {
		assert.Equal(t, expectedTitles[i], track.Title, "unexpected track title")
	}
}

func TestProcessTracklist_Error(t *testing.T) {
	invalidFilePath := "tracklist_idasfhaj.txt"

	ap := &interfaces.AudioProcessor{}
	uc := NewTrackUsecase(ap)

	tracks, err := uc.ProcessTracklist(invalidFilePath, artist, album)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	assert.Nil(t, tracks, "expected nil tracks")
}
