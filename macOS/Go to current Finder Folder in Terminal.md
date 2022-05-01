
# This goes into your `.bash_profile`

```bash
cdf () {
	# cd into forefront Finder window.
	local path=$(osascript -e 'tell app "Finder" to POSIX path of (insertion location as alias)')
	echo "$path"
	cd "$path"
}
```

