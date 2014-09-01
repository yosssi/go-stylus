package stylus

// Compile compiles the Stylus file of the specified path and
// returns two channles: One returns the path of the compiled
// CSS file. Another returns the error when it occurs.
func Compile(path string) (<-chan string, <-chan error) {
	return nil, nil
}
