# endsentence

End your comments with a period. This will check your go files for comments on
methods/functions that are commented but does not have a period in the end of
the sentence.

Make sure to use golint in addition to this.

This has been done before as a pull request to golint but was deemed to be a
small thing so it was skipped. See
[golint/#128](https://github.com/golang/lint/pull/128) for more info.

The code that is in here is based mostly on golint and I wanted to do some
quick hacking around the parser so I decided to implement this small thing.
