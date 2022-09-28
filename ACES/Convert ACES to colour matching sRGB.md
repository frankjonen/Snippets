# Convert an ACES .exr to a matching sRGB file

This script is geared towards [Affinity](https://affinity.serif.com) suite users. While we can design and hand-off in proper ACES formats, we can’t export colour matching files. This script solves that issue.

---

### Dependencies
- [Homebrew](https://brew.sh)
- [OCIO](https://formulae.brew.sh/formula/opencolorio)
- [OIIO](https://formulae.brew.sh/formula/openimageio)

---

## About

While you’ve probably already have a working OCIO installation, you may not be familiar with the OpenImageIO toolkit. This is what gives us the `oiiotool` an image processing app that lets us utilise our OCIO configuration to convert images between color description types. This way we can go from OCIO manages spaces to ICC profile managed spaces with ease.

Pop the following code into a shell script, and run it followed by the name of the file you’re looking to convert.

#### Example Setup
1. I’ve named the script `acesout` and gave it execution permissions via `chmod u+x acesout`.  
2. `acesout samplefile.exr`

## Caveats

1. The filename cannot have spaces. There is currently no way of `oiiotool` to accept spaces in a filename, even quoted and double quoted.
2. The script is made for single files. If your filenames are `sequence.001.exr` it will recognise `001` as the suffix and remove from there to get the base image name. While still correct behaviour, it may cause issues with your file tracking processes. 


## The Script


```bash

#!/bin/bash

fileIn="$1"
fileOut=${fileIn%.*}

oiiotool "$fileIn" --colorconvert "ACES - ACEScg" "Output - sRGB (D60 sim.)" --ch R,G,B -o:type=uint16 "$fileOut".tif

```

#### Disclaimer
1. If you’ve designed towards the D60 simulation, use `"Output - sRGB (D60 sim.)"` instead.
2. If you want to keep the alpha channel, use `--ch R,G,B,A` instead.
3. For 8-bit files use `-o:type=UINT8` instead, the file format is inferred by the suffix after "`$fileOut`".

