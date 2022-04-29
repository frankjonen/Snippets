# Symbolic Link for Beorg

- Create a Symbolic Link for Beorg (https://beorgapp.com)
- *FROM:* Your local iCloud Beorg Documents folder
-   *TO:* Your Documents folder

Now you can just visit `~/Documents/org` in Emacs and have it sync semi-live to Beorg on your iPad, iPhone or other Macs.

```bash
ln -s ~/Library/Mobile\ Documents/iCloud\~com\~appsonthemove\~beorg/Documents/org ~/Documents/org
```
