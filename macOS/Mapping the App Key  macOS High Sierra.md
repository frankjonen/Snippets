# Mapping the App Key: macOS High Sierra

To use the `App Key` (found on non-Mac keyboards) you have to remap it manually from the Terminal. 
The keyboard modifier mapper in macOS only lets you change: Caps Lock, Control, Option, Command.

For off-beat keys you need to use the *Human Interface Device Utility*, short *hidutil*.


## Set `App Key` to `Right Option Key`

```bash
hidutil property --set '{"UserKeyMapping":[{"HIDKeyboardModifierMappingSrc":0x700000065,"HIDKeyboardModifierMappingDst":0x7000000e6}]}'

# Set `App Key` to `Right GUI Key` (Command Key on Mac, Windows key on Windows, etc)
hidutil property --set '{"UserKeyMapping":[{"HIDKeyboardModifierMappingSrc":0x700000065,"HIDKeyboardModifierMappingDst":0x7000000e7}]}'
```

Reference: https://developer.apple.com/library/content/technotes/tn2450/_index.html
