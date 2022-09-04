# Batch rename a file in all subfolders

If you have a file with the same name in a lot of folders within your project folder and want to rename them, here’s how.


```bash

find . -type f -name "Filename.ext" -execdir mv {} "Filename Changed.ext" \;

```

Just change `Filename.ext` to the file you’re searching and `Filename Changed.ext` to its new name.
