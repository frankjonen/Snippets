# Keyboard Speed settings: macOS High Sierra

## Nice and fast

```bash
defaults write NSGlobalDomain KeyRepeat -int 2
defaults write NSGlobalDomain InitialKeyRepeat -int 15
```

## Plaid

```bash
defaults write NSGlobalDomain KeyRepeat -int 1
defaults write NSGlobalDomain InitialKeyRepeat -int 10
```

# Caveats

GUI apps don't always honour that setting but rely on the key repeat setting first. To address that set one of the following:

- Global for all apps
```bash
defaults write -g ApplePressAndHoldEnabled -bool false
```
- Specify an app
```bash
defaults write -app MacVim ApplePressAndHoldEnabled -bool false
```

You can also set it globally to false and set it app specific to `true`.
