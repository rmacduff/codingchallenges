package ccwc

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type WordCount struct {
	input      io.Reader
	output     io.Writer
	countBytes bool
}

type option func(*WordCount) error

func WithInput(input io.Reader) option {
	return func(wc *WordCount) error {
		if input == nil {
			return errors.New("input cannot be nil")
		}
		wc.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(wc *WordCount) error {
		if output == nil {
			return errors.New("output cannot be nil")
		}
		wc.output = output
		return nil
	}
}

func FromArgs(args []string) option {
	return func(wc *WordCount) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		CountBytes := fset.Bool("c", false, "count bytes")
		fset.SetOutput(wc.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		wc.countBytes = *CountBytes

		args = fset.Args()
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		wc.input = f
		return nil
	}
}

func New(options ...option) (WordCount, error) {
	wc := WordCount{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range options {
		err := opt(&wc)
		if err != nil {
			return WordCount{}, err
		}
	}
	return wc, nil
}

func (wc WordCount) CountBytes() (int, error) {
	chars := 0
	scanner := bufio.NewScanner(wc.input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if err := scanner.Err(); err != nil {
			return chars, err
		}
		chars += len(bytes)
	}
	return chars, nil
}

func RunCLI() {
	wc, err := New(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if wc.countBytes {
		bytes, err := wc.CountBytes()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, bytes)
	} else {
		fmt.Println("only char count is supported")
	}
}
