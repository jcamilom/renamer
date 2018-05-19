# renamer

A cli tool to change filenames using regular expresions.

## How to use it

```
$ go get -u github.com/jcamilom/renamer

# walk the current dir (subdirs included!) and promp a
# message when a file matches this regular expression
# "^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$"
$ renamer
```
