# go-xmmsclient [![Travis](https://api.travis-ci.org/dsvensson/go-xmmsclient.svg?branch=master)](https://travis-ci.org/dsvensson/go-xmmsclient)

A first exploration of the Go programming language with a target of implementing a full xmmsclient library to communicate with the [XMMS2](https://github.com/xmms2/xmms2-devel) music player.

## Stuff to figure out

* How to drive the networking
 * sync/async calls
 * repeating server side event driven broadcasts (skip signals?)
* go generate the whole API from [ipc.xml](https://github.com/xmms2/xmms2-devel/blob/master/src/ipc.xml)
* How to deal with the highly dynamic API of `xmmsc_medialib_query`.
* What is idiomatic Go-esque style?
 * Naming things
 * Package separation
 * File separation
* Testing, coverage
