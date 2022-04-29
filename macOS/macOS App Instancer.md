# macOS App Instancer

A little macOS shell script for when you need to have multiple instances of an app open.

***CAVEATS:***
- Instancing on macOS will break 
- autosave behaviour if the app is built around .plists
- cloud sync (if the app does that)
- window status doesn't get saved (only for the last instance you quit)

**Usage**
1. Run the script and follow the on-screen instructions.
2. Basically you type in the name of the app, hit enter and type in the number of instances you want.

---

```bash
#! /bin/bash

echo
echo "$(tput setaf 7;tput bold)          Instance Application $(tput setaf 1)"
echo
read -p "     Enter App Name: $(tput sgr0)" "app_name"
echo -n "$(tput setaf 1;tput bold)"
read -p "     How many times? $(tput sgr0)" inst_count
echo
echo -n "          "
echo "$(tput setaf 7;tput bold)$app_name$(tput sgr0) is$(tput setaf 1;tput bold) GO! $(tput sgr0)"
echo

for (( c=1; c<=$inst_count; c++ ))
do
  open -n -a "$app_name"
done

exit

```
