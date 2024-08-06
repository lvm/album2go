# album2go

Downloads a video from Youtube and splits tracks based on a tracklist. It'll also try to save ID3 Tags.

The next iteration of [youtube-album](https://github.com/lvm/youtube-album) but rewritten (in Golang) and simplified.

## Usage 

```
album2go <YouTube URL> --artist <artist> --album <album> --tracklist <tracklist> --output <directory> [--verbose]
```

## Tracklist format

Examples of valid tracklist formats are either of these:

``` 
1. Song Name 01:23:45
2. Song Name - 01:23:45
3) Song Name 01:23:45
4) Song Name - 01:23:45
``` 

That is:

``` 
1. Anthropogenic End Transmission 2:16
2. The Geocide 3:43
3. Be Still Our Bleeding Hearts 3:54
4. Vulturous 4:59
[...]
``` 

## Third party dependencies

This software uses `ffmpeg` (required) and `id3v2` (optional, nice to have) to work, one to slice and one to tag audio files respectivelly. 

### `ffmpeg`  

* GNU/Linux, `apt install ffmpeg libmp3lame0` (or your distro equivalent) 
* macOS, `brew install ffmpeg`
* Windows, follow [this link](https://www.ffmpeg.org/download.html#build-windows)

### `id3v2`

* GNU/Linux, `apt install id3v2` (or your distro equivalent)
* macOS, `brew install id3v2`
* Windows, I'm afraid you're out of luck, you'll have to tag them manually

## License 

See [LICENSE](LICENSE)
