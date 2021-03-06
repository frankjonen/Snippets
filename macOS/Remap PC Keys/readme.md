
# Installation

Copy the file `com.local.KeyRemapping.plist` to `~/Library/LaunchAgents`.
To apply the changes log out and log in again.

## Check up
Check the new key setting in the `Terminal` via

```bash
hidutil property --get "UserKeyMapping"
```

# Modifications

This `plist` file makes the following changes

| Previous Key |      New Key     |
|:------------:|:----------------:|
| App Key      | Right Option Key |
| Print Screen | F13              |
| Scroll Lock  | F14              |
| Pause        | F15              |

# Known Issues

On *MAX Keyboards* **Blackbird** keyboard you need to hold <kbd>SHIFT</kbd> when pressing <kbd>ScrLk</kbd> or <kbd>Pause</kbd>. Only <kbd>PrtSc</kbd> works without shift.
