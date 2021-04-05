package log

import "testing"

func TestL(t *testing.T) {
	L(Red, "Helloworld")
	L(Green, "Helloworld")
	L(Yellow, "Helloworld")
	L(Blue, "Helloworld")
	L(Magenta, "Helloworld")
	L(Cyan, "Helloworld")
	L(White, "Helloworld")
}

func TestLW(t *testing.T) {
	LW(Red, "Helloworld")
	LW(Green, "Helloworld")
	LW(Yellow, "Helloworld")
	LW(Blue, "Helloworld")
	LW(Magenta, "Helloworld")
	LW(Cyan, "Helloworld")
	LW(White, "Helloworld")
}
