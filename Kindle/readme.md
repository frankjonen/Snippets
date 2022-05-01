# mdMOBI
A simple script to send a Markdown file to your Kindle for proofreading.

***

## How it’s used

### The quick-fire way:
1. Open the Markdown file in Vim.
2. Hit the `F12` key.
3. Select the Kindle device.

### From the Terminal, without Vim
1. Navigate to the folder with your Markdown file.
2. Type `mdmobi "My Markdown File"` and hit `Enter`.
3. Select the Kindle device.

> **Note 1:** The `" "` around the filename, that just makes it easy to deal with spaces.

> **Note 2:** Leave the `.md` suffix off. The script adds that automatically and uses the document name to generate all other filenames.

> **Tip:** You can add this [script](https://github.com/frankjonen/Snippets/blob/main/macOS/Go%20to%20current%20Finder%20Folder%20in%20Terminal.md) to your `.bash_profile` and only have to type `cdf` in Terminal to get to the currently active Finder window.


***


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
