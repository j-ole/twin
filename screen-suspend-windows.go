//go:build windows

package twin

import "fmt"

func (screen *terminalScreen) suspend() error {
	return fmt.Errorf("suspend is not supported on windows")
}
