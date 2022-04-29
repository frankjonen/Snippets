# Homebrew Basics: macOS High Sierra

## Before new installs

```bash
brew cleanup

```

> This removes all source code and any other crud. Do this if you haven't done it yet.

## Update the Homebrew index

```bash
brew update
```

> This just updates the index, not to fret

## Upgrade all applications

```bash
brew upgrade
```

> This actually downloads and upgrades the outdated items

## Clean up after yourselves 

```bash
brew cleanup                           
```

> This is the right place to do it, so you don't have to remember it next time.

> When you uninstall a brew formula, `cleanup` is also a good idea before you re-install something.


### Bonus Tip

```bash
brew upgrade && brew cleanup
```

> Download and install in sequence. The `&&` means that the second command will only fire if the first one ran successfully.

