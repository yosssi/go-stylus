package stylus

import "testing"

func TestCompile(t *testing.T) {
	pathc, errc := Compile("test/1.styl")

	select {
	case <-pathc:
	case err := <-errc:
		t.Error(err)
	}
}

func TestCompile_cmdErr(t *testing.T) {
	pathc, errc := Compile("not_exist_file")

	select {
	case <-pathc:
		t.Error("err should be returned")
	case <-errc:
	}
}

func TestCompile_invalidMsg(t *testing.T) {
	pathc, errc := Compile("-V")

	select {
	case <-pathc:
		t.Error("err should be returned")
	case <-errc:
	}
}

func Test_execCmd_err(t *testing.T) {
	outc, errc := execCmd("not_exist_file")

	select {
	case <-outc:
		t.Error("err should be returned")
	case <-errc:
	}
}

func Test_execCmd(t *testing.T) {
	outc, errc := execCmd("test/2.styl")

	select {
	case <-outc:
	case err := <-errc:
		t.Error(err)
	}
}
