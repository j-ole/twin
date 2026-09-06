[![Docs](https://img.shields.io/badge/hello-Godoc-blue.svg?label&logo=go)](https://pkg.go.dev/github.com/walles/twin#section-documentation)
[![Linux CI](https://github.com/walles/twin/actions/workflows/linux-ci.yml/badge.svg)](https://github.com/walles/twin/actions/workflows/linux-ci.yml)
[![Windows CI](https://github.com/walles/twin/actions/workflows/windows-ci.yml/badge.svg)](https://github.com/walles/twin/actions/workflows/windows-ci.yml)

Twin is a library for drawing to the terminal screen, originally built for the
[moor](https://github.com/walles/moor) pager and battle-tested across Linux,
macOS and Windows ever since.

# Demo

The [moor](https://github.com/walles/moor) pager was built using twin.

So is [ftop](https://github.com/walles/ftop):

![ftop screenshot](screenshot.png)

# Installation

```
go get github.com/walles/twin
```

# Usage

* Open a screen using `twin.NewScreen()` or one of its friends
* Use `screen.SetCell()` to draw characters in an off-screen buffer
* Use `screen.Show()` to send the buffer contents to the terminal
* Listen to events on the `screen.Events()` channel, for keyboard, mouse and
  window resize
* Close the screen when done using `screen.Close()`

Twin opens an alternate screen buffer that it draws into.

See the full API docs at
[pkg.go.dev/github.com/walles/twin](https://pkg.go.dev/github.com/walles/twin#section-documentation).

# Making a new release

1. `git tag --annotate vX.Y.Z`, note the leading `v` in the version number.
   Write something descriptive in the annotation message.
1. `git push --tags`

# TODO

## Done

- Add an `ftop` screenshot demonstrating what twin can do
- Make Color implement stdlib's color.Color
- Enable revive's "exported" rule and validate comments for our public surface
- Get CI running for PRs and pushes to main, and add a CI status badge at the
  top of the README.
- Review the API based on how it is being used
