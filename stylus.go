package stylus

import (
	"bytes"
	"fmt"
	"os/exec"
)

const stylusCmd = "stylus"

var sucPrefix = []byte{32, 32, 27, 91, 57, 48, 109, 99, 111, 109, 112, 105, 108, 101, 100, 27, 91, 48, 109, 32}

// Compile compiles the Stylus file of the specified path and
// returns two channles: One returns the path of the compiled
// CSS file. Another returns the error when it occurs.
func Compile(path string) (<-chan string, <-chan error) {
	pathc := make(chan string)
	errc := make(chan error)

	go func() {
		outc, cmdErrc := execCmd(path)

		var out []byte

		select {
		case out = <-outc:
		case err := <-cmdErrc:
			errc <- err
			return
		}

		if !bytes.HasPrefix(out, sucPrefix) {
			errc <- fmt.Errorf("command's output message should have prefix %q [actual: %q]", string(sucPrefix), string(out))
			return
		}

		pathc <- string(bytes.TrimSpace(bytes.TrimPrefix(out, sucPrefix)))
	}()

	return pathc, errc
}

// execCmd executes the Stylus command.
func execCmd(path string) (<-chan []byte, <-chan error) {
	outc := make(chan []byte)
	errc := make(chan error)

	go func() {
		out, err := exec.Command(stylusCmd, path).Output()
		if err != nil {
			errc <- err
			return
		}
		outc <- out
	}()

	return outc, errc
}
