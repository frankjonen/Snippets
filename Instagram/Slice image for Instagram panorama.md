# Slice image for Instagram panorama

Crop an image into 3 slices for an Instagram panorama # For best results the image should be a 3x1 aspect ratio. If you have a 2x1 aspect ratio image, change the `3x1@` to a `2x1@` The `%d` part is a placeholder for the automatic numbering.

```bash
convert "File Name.tiff" -crop 3x1@ slice-%d.png
```
