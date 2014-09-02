package stylus

import (
	"fmt"
	"os/exec"
	"strings"
)

const stylusCmd = "stylus"

const compSucMsg = "compiled"

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

		tokens := strings.Split(strings.TrimSpace(string(out)), " ")

		if len(tokens) < 2 {
			errc <- fmt.Errorf("command's output is invalid [out: %v]", tokens)
			return
		}

		msg := strings.TrimRight(
			strings.TrimLeft(
				tokens[0],
				string([]byte{27, 91, 57, 48, 109}),
			),
			string([]byte{27, 91, 48, 109}),
		)

		if msg != compSucMsg {
			errc <- fmt.Errorf("command's output message should be %s [actual: %s]", compSucMsg, msg)
			return
		}

		pathc <- tokens[1]
	}()

	return pathc, errc
}

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
