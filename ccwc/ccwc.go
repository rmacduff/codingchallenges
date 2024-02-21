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
	countLines bool
	countWords bool
	countChars bool
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
		countBytes := fset.Bool("c", false, "count bytes")
		countLines := fset.Bool("l", false, "count lines")
		countWords := fset.Bool("w", false, "count words")
		countChars := fset.Bool("m", false, "count chars")
		fset.SetOutput(wc.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		wc.countBytes = *countBytes
		wc.countLines = *countLines
		wc.countWords = *countWords
		wc.countChars = *countChars

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
	bytes := 0
	scanner := bufio.NewScanner(wc.input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return bytes, err
		}
		bytes++
	}
	return bytes, nil
}

func (wc WordCount) CountLines() (int, error) {
	lines := 0
	scanner := bufio.NewScanner(wc.input)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return lines, err
		}
		lines++
	}
	return lines, nil
}

func (wc WordCount) CountWords() (int, error) {
	words := 0
	scanner := bufio.NewScanner(wc.input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return words, err
		}
		words++
	}
	return words, nil
}

func (wc WordCount) CountChars() (int, error) {
	chars := 0
	scanner := bufio.NewScanner(wc.input)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return chars, err
		}
		chars++
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

	switch {
	case wc.countBytes:
		bytes, err := wc.CountBytes()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%8d %s\n", bytes, wc.input.(*os.File).Name())
	case wc.countLines:
		lines, err := wc.CountLines()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%8d %s\n", lines, wc.input.(*os.File).Name())
	case wc.countWords:
		words, err := wc.CountWords()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%8d %s\n", words, wc.input.(*os.File).Name())
	case wc.countChars:
		chars, err := wc.CountChars()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%8d %s\n", chars, wc.input.(*os.File).Name())
	default:
		fmt.Println("unsupported option")
	}
}
