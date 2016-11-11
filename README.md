# go-xmmsclient [![Travis](https://api.travis-ci.org/dsvensson/go-xmmsclient.svg)](https://travis-ci.org/dsvensson/go-xmmsclient) [![Coverage Status](https://coveralls.io/repos/github/dsvensson/go-xmmsclient/badge.svg)](https://coveralls.io/github/dsvensson/go-xmmsclient)

A first exploration of the Go programming language with a target of implementing a full xmmsclient library to communicate with the [XMMS2](https://github.com/xmms2/xmms2-devel) music player.

## Stuff to figure out

* go generate the whole API from [ipc.xml](https://github.com/xmms2/xmms2-devel/blob/master/src/ipc.xml)
  * integrate properly with go:generate
  * XmmsList/XmmsDict params today that can be []string, or even ...string, map[..]...
  * Emit API documentation
  * Generate constants/enums.
  * Generate broadcasts/signals
  * Check enum parameters against the int value passed in.
  * Heuristics
    * Better naming of some functions
    * Maybe generate more functions for some commands and hide params internally
