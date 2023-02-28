# Vim-Plug Cleaner

If you’re using Vim-Plug instead of installing Vim plugins by hand, you might have noticed your plug-in folder swelling up to gigantic sizes. 

This is due to a bug in Vim-Plug that downloads the **entire** GitHub repository into your plugin folder. The developers of Vim-Plug deem this impossible to fix. So here’s quick script.

I just named it `vimclean` and put it in my scripts folder.

> Example: Cleanup of UltiSnips

```bash
vimclean

Which Plugin to clean up: ultisnips
Done
```

`ultisnips` is the plug-in’s folder name here.


```bash
#! /bin/bash

echo -n "Which Plugin to clean up: "
read name

cleaner () {
	rm -f .gitignore *.md *.txt Docker*.* Makefile Pipfile Pipfile*.* *.ini *.py pylintrc &&
	rm -rf .github test ctags docker
	echo "Done."
}

echo "Cleaning: $name"

( cd .vim/plugged/"$name" && cleaner )
```