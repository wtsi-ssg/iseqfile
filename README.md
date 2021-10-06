# iseqfile
Drop-in replacement for a Sanger-specific imeta query for getting iRODS seq paths, to get the result fast.

# Install

```
go install github.com/wtsi-ssg/iseqfile@latest

```

To use as a drop-in replacement for imeta (but note it only supports one specific query!), you could add an alias to your shell login script, eg.

```
alias imeta=$(go env GOPATH)/bin/iseqfile
```

# Usage

Assuming you've aliased imeta to iseqfile, then:

```
imeta qu -z seq -d type = cram and target = 1 and lane = 2 and tag_index = 23 and id_run = 41208
```

NB: this is currently the ONLY kind if query supported.
-z is ignored; it will always be "seq".
target is ignored; it will always be "1".
id_run, lane and tag_index are required, but can be in any order.
type is used, but defaults to "cram".

This will return results in the same format as imeta, but should do so consistently in about 1 second, instead of up to many minutes when using the real imeta.
