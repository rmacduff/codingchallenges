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

func TestCountLines(t *testing.T) {
	t.Parallel()

	args := []string{"-l", "testdata/lines.txt"}
	wc, err := ccwc.New(ccwc.FromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 4

	got, err := wc.CountLines()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
<<<<<<< Updated upstream
=======

func TestCountWords(t *testing.T) {
	t.Parallel()

	args := []string{"-w", "testdata/words.txt"}
	wc, err := ccwc.New(ccwc.FromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 7

	got, err := wc.CountWords()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestCountChars(t *testing.T) {
	t.Parallel()

	args := []string{"-w", "testdata/chars.txt"}
	wc, err := ccwc.New(ccwc.FromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 62

	got, err := wc.CountChars()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
>>>>>>> Stashed changes
