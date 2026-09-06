[![Docs](https://img.shields.io/badge/hello-Godoc-blue.svg?label&logo=go)](https://pkg.go.dev/github.com/walles/twin#section-documentation)
[![Linux CI](https://github.com/walles/twin/actions/workflows/linux-ci.yml/badge.svg)](https://github.com/walles/twin/actions/workflows/linux-ci.yml)
[![Windows CI](https://github.com/walles/twin/actions/workflows/windows-ci.yml/badge.svg)](https://github.com/walles/twin/actions/workflows/windows-ci.yml)

Twin is a low-level library for drawing to the terminal screen: you own the
widgets and layout, twin owns the terminal. Originally built for the
[moor](https://github.com/walles/moor) pager, it's been battle-tested across
Linux, macOS and Windows ever since.

# Features

* **Clickable hyperlinks** in terminals that support them
* **Automatic color downsampling**, so full 24-bit color degrades gracefully
  on terminals that don't support truecolor
* **Self-healing raw mode**: detects when another program resets your
  terminal settings behind your back and restores them
* **Transparent suspend/resume**: Ctrl-Z drops you to the shell and back
  cleanly
* **Efficient rendering** that only redraws what changed on screen, not the
  whole frame every time
* **Wide-character support**, so CJK and other double-width characters render
  without corrupting the layout
* **Native progress indicators**, reporting task progress to the terminal /
  taskbar, including error and indeterminate states
* **Terminal background color detection**, so your program can adapt to it if
  needed

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
