# Python Setup macOS Mojave and up

## My setup to make learning Python easier

1. First we get the most recent Python version. As of this writing, that's 3.8
2. If you don't use Homebrew yet, now's a good time to fix that.
	1. Open Terminal and go like this:
	```bash
	/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
	```

Done? Now do:

```bash
brew install python@3.8
```

## Let's get us a decent Vim for the Desktop

```bash
brew install macvim
```

Homebrew installed apps cannot be seen by Spotlight, so you can't use that to launch these apps. Here's a workaround:

- Create an Automator app that runs this shell script
	```bash
	open /usr/local/opt/macvim/MacVim.app $@
	```

- Or just get it readymade with Spotlight launchability from here
	https://macvim-dev.github.io/macvim/

> Difference? The brew version only has Python3 built-in (:py3 commands).

## Cleanup and set it up

```bash
brew cleanup
```

Brewage is done, now we set our system up to have **both**, the legacy Python many macOS apps need to function and the current one that we want to learn. Without screwing up either.

### Put this in your .bash_profile 

```bash
export PYTHONPATH=/usr/local/opt/python@3.8/bin
export LDFLAGS="-L/usr/local/opt/python@3.8/lib"
export PKG_CONFIG_PATH="/usr/local/opt/python@3.8/lib/pkgconfig"

export PATH=$PATH:$PYTHONPATH
```

PIP for us is now PIP3 and the PYTHON command is now PYTHON3. Test it by typing `PYTHON -V` and `PYTHON3 -V` If everything went correctly, you now have both versions available to you.

Brew's MacVim is here: `/usr/local/Cellar/macvim/[the version number]/MacVim.app` Double click that or type `open MacVim.app` and once launched, drag the icon where you want it in your Dock

The easiest way to run your Python script in Vim is via:

```vim
:!python3 %
```

The `%` part is your current buffer. You can set this up as a key-binding for running your code with one key-press.

That's basically it. Now you can wander off and customise the rest to your heart's content.
