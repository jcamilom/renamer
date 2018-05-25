# renamer

A cli tool to change filenames using regular expresions.

## How to use it

```
$ go get -u github.com/jcamilom/renamer

# Usage: renamer [<path>] \"<match_regexp>\" \"<replace_string>\"
# If path no provided, current dir is used.
# Be careful, subdirs are always included!

# Example
$ renamer "^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)\$" "\$2 - \$1 - \$3 of \$4.\$5"
```
