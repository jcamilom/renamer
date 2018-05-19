# renamer

A cli tool to change filenames using regular expresions.

## How to use it

```
$ go get -u github.com/jcamilom/renamer

# walk the current dir (subdirs included!) and renames files
# that match this regular expression:
# "^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$"
# The new new is given by this Replace String:
# "$2 - $1 - $3 of $4.$5"
$ renamer
```
