Twin is a library for drawing to the terminal screen.

# Installation

```
go get github.com/walles/twin
```

# Usage
* Open a screen using `twin.NewScreen()` or one of its friends
* Use `screen.SetCell()` to draw characters in an off-screen buffer
* Use `screen.Show()` to send the buffer contents to the terminal
* Listen to events on the `screen.Events()` channel, for keyboard mouse and
  window resize
* Close the screen when done using `screen.Close()`

Twin opens an alternate screen buffer that it draws into.

# TODO

- Make Color implement stdlib's image.Color
- Should we use https://pkg.go.dev/github.com/lucasb-eyer/go-colorful to compute
  distances?
- Enable revive's "exported" rule and validate comments for our public surface
