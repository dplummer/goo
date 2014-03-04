Goo
---

A Go port of the liquid template language.

Testing
-------

Install the `inotify-utils` package.

Run:

```
inotifywait -m -r -e close_write --exclude '.*\.test' . | while read line; do go test ; done
```
