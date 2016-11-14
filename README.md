# go-xmmsclient [![Travis](https://api.travis-ci.org/dsvensson/go-xmmsclient.svg)](https://travis-ci.org/dsvensson/go-xmmsclient) [![Coverage Status](https://coveralls.io/repos/github/dsvensson/go-xmmsclient/badge.svg)](https://coveralls.io/github/dsvensson/go-xmmsclient)

A first exploration of the Go programming language with a target of implementing a full xmmsclient library to communicate with the [XMMS2](https://github.com/xmms2/xmms2-devel) music player.

## What's in here?

* `xmmsclient`
    * The actual library that's budding.
* `genipc`
    * A hacky code generator that takes [ipc.xml](https://github.com/xmms2/xmms2-devel/blob/master/src/ipc.xml) and generates the API.
* `examples/test`
    * A weird test client that uses whatever is being worked on right now.

## License?

[ISC](https://opensource.org/licenses/ISC), as far from [tl;dr](https://www.gnu.org/licenses/gpl-3.0.txt) as you can possibly get.
