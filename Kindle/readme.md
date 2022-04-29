# mdMOBI
A simple script to send your Markdown file to your Kindle for proofreading.

## Requirements
- [Pandoc](https://github.com/jgm/pandoc/releases)
- [Kindle Previewer](https://kdp.amazon.com/help/topic/G202131170)
- [Send to Kindle](https://www.amazon.com/gp/sendtokindle) *Make sure this is set-up to your account!*

***

## `bash_profile`
Here we have the following entries for your `.bash_profile`.

### Tool Path
That’s where our scripts live. I think it’s a good idea to have them in your home folder. I have mine in `.tools` for example, so they don’t bother me when I don’t need them.

### Kindle Setup
Making `kindlegen` available to the system. Kindlegen creates the `.mobi` files for us. It used to be standalone, now it’s part of the *Kindle Previewer* application bundle.

***

## `vim` Folder
A `markdown` setting for *Vim* to make the whole deal a *single key-press* affair. Merge the `ftplugin` folder with yours or just drag the file over.