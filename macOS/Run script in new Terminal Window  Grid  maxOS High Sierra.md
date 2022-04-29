# Run a script in new Terminal Window Grid: macOS High Sierra

## In `.bash_profile

```bash
HISTCONTROL=ignoreboth
```

> This excludes anything with a space before it from your `.bash_history`


```bash
function scriptSpawn(){
# The Apple Script that creates the grid items

osascript <<EOD
    tell application "Finder"
        set desktopBounds to bounds of window of desktop
        set desktopWidth to item 3 of desktopBounds
        set desktopHeight to item 4 of desktopBounds

        set windowWidth to desktopWidth / 3
        set windowHeight to desktopHeight / 2

        -- - - - - - - -  Column 1

        set window_c1_vpos1 to 22
        set window_c1_vpos2 to desktopHeight - 1
        set window_c1_hpos1 to 0
        set window_c1_hpos2 to windowWidth

        -- - - - - - - -  Column 2
        -- - - - - - - - - - - - - - - - - - - - - ROW 1

        -- grid item 1
        set window_1_vpos1 to 22
        set window_1_vpos2 to windowHeight
        set window_1_hpos1 to 0
        set window_1_hpos2 to windowWidth

        -- grid item 2
        set window_2_vpos1 to 22
        set window_2_vpos2 to windowHeight
        set window_2_hpos1 to windowWidth
        set window_2_hpos2 to 2 * windowWidth

        -- grid item 3
        set window_3_vpos1 to 22
        set window_3_vpos2 to windowHeight
        set window_3_hpos1 to 2 * windowWidth
        set window_3_hpos2 to 3 * windowWidth

        -- - - - - - - - - - - - - - - - - - - - - ROW 2

        -- grid item 4
        set window_4_vpos1 to windowHeight
        set window_4_vpos2 to 2 * windowHeight
        set window_4_hpos1 to 0
        set window_4_hpos2 to windowWidth

        -- grid item 5
        set window_5_vpos1 to windowHeight
        set window_5_vpos2 to 2 * windowHeight
        set window_5_hpos1 to windowWidth
        set window_5_hpos2 to 2 * windowWidth

        -- grid item 6
        set window_6_vpos1 to windowHeight
        set window_6_vpos2 to 2 * windowHeight
        set window_6_hpos1 to 2 * windowWidth
        set window_6_hpos2 to 3 * windowWidth
    end tell

	tell application "Terminal"
		set newTab to do script "$1"
		set bounds of front window to {window_$2_hpos1, window_$2_vpos1, window_$2_hpos2, window_$2_vpos2}
	end tell
EOD
}

```

> Make sure EOD does not have a trailing space!

### Usage

Call the AppleScript to open a new Terminal window/tab and run the alias piped into it via the variable input.

#### Window Sizes and Positions

These are derived from your screensize. The 'c1' parameter gets you a full-height panel pinned to the left.
- `scriptSpawn example 1` would run the `example` alias in a new Terminal window anchored top left and a size of 1/3rd your screen width and half its height.

## Creating the Shortcut

### Creating or editing the `.inputrc` file

We're setting up `F9` to run the `scriptSpwan` function with the alias `example` and the position flag `1` appended as flags. Or use your own scripts and tools instead of an alias. For example I have `htop` mapped to a key for `c1`.

```bash
"\e[20~": " scriptSpawn example 1 && clear \C-m"
```

#### Overview

Here's a basic overview how the script is set up visually.

```txt
# - - - - - - - - - - - - - - - -
# Script Spawn Grid Overview
# - - - - - - - - - - - - - - - -
#     1 | 2 | 3
#     4 | 5 | 6

#    c1 | 2 | 3
#       | 5 | 6
```

