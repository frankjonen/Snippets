
# Replace Relative Stepbacks in HTML documents

If you have a bunch of `../../` in your local website folder from testing static sites right from the Finder, this can help you. It replaces all occurrences of `../` with a single `/` representing the site's root. Instead of stepping back from nested taxonomies to link assets, you can just go from the site's root down. Much more reliable.

This script creates backups with the `.bup` suffix of each changed file, so you have a fallback and can start over in case it got messed up.

You can remove the backups later on with:


```bash
find . -name "*.bup" -delete
```



## The Commands


### For single steps (`../`)

```bash
find . -name '*.html' | xargs sed -i .bup 's/\.\.\//\//g
```

### For double steps (`../../`)

```bash
find . -name '*.html' | xargs sed -i .bak 's/(\.\.\/){2}/\//g
```

### Deeper?

For deeper nesting change the `{2}` to the number of repetitions you need.
