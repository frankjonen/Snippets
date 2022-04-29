# Run script in new Terminal Window: macOS High Sierra

## This goes into `.bash_profile

```bash
HISTCONTROL=ignoreboth
```

> This excludes anything with a space before it from your `.bash_history`.


```bash
function triggerBunny() {
osascript <<EOD
	tell application "Terminal"
		set newTab to do script "$@"
	end tell	
EOD
}
```

Call AppleScript to open a new Terminal window/tab and run the alias piped into it via the variable `triggerBunny example` would run the `example` alias in a new Terminal window/tab.

## This goes into our `.inputrc`

We set up F9 to run the `triggerBunny` function with the alias `example` appended as a flag.

```bash
"\e[20~": " triggerBunny example \C-m"
```
