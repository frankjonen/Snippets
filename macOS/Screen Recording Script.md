# Screen Recording Script.md

This is macOS focused, on other systems you need to change `-f avfoundation` to one of the following:

- Linux: `x11grab`
- Windows: `gdigrab` all displays as one
	- `dshow` discrete recording

Check for screen to record from via:

- `ffmpeg -f avfoundation -list_devices true -i ""`
	- And set it after -i (my screen is 0)

- To add audio, set the audio device after it
	- Example: `-i (0.1)`

## The Basics

1. Screen recording is ended with `CTRL+C`
2. Script is used as: scriptname filename
3. Output is without audio
4. To add audio set the device as in the example above.
5. Matroska suffix is added for convenience, change it if you like.

> Get FFMPEG via Homebrew. http://brew.sh

```bash
#!/bin/sh

outfile=$1

ffmpeg  -video_size 2560x1440 \
	-framerate 30 \
	-pix_fmt 0rgb \
	-f avfoundation \
	-i 0 \
	-c:v libx264rgb \
	-crf 0 \
	-preset ultrafast \
	-color_range 2 \
	$outfile.mkv
```
