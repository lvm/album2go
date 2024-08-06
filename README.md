# album2go

Downloads a video from Youtube and splits tracks based on a tracklist. It'll also try to save ID3 Tags.

The next iteration of [youtube-album](https://github.com/lvm/youtube-album) but rewritten (in Golang) and simplified.

## Usage 

```
album2go <YouTube URL> --artist <artist> --album <album> --tracklist <tracklist> --output <directory> [--verbose]
```

## Tracklist format

An example of a valid tracklist format is:

``` 
<trackno>. <trackname> <min>:<sec>
``` 

Which translates to:

``` 
1. Anthropogenic End Transmission 2:16
2. The Geocide 3:43
3. Be Still Our Bleeding Hearts 3:54
4. Vulturous 4:59
[...]
``` 

## License 

See [LICENSE](LICENSE)
