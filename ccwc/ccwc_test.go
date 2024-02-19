package ccwc_test

import (
	"ccwc"
	"testing"
)

func TestCountBytes(t *testing.T) {
	t.Parallel()

	args := []string{"-c", "testdata/bytes.txt"}
	wc, err := ccwc.New(ccwc.FromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 27

	got, err := wc.CountBytes()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
