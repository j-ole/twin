[![Linux CI](https://github.com/walles/twin/actions/workflows/linux-ci.yml/badge.svg)](https://github.com/walles/twin/actions/workflows/linux-ci.yml)
[![Windows CI](https://github.com/walles/twin/actions/workflows/windows-ci.yml/badge.svg)](https://github.com/walles/twin/actions/workflows/windows-ci.yml)

Twin is a library for drawing to the terminal screen.

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

# TODO

- Review the API based on how it is being used
- Once [ftop](https://github.com/walles/ftop) has been migrated to use this
  twin, add an `ftop` screenshot

## Done

- Make Color implement stdlib's color.Color
- Enable revive's "exported" rule and validate comments for our public surface
- Get CI running for PRs and pushes to main, and add a CI status badge at the
  top of the README.
