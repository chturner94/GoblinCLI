package goblincli

import (
	"bufio"
	"os"

	"github.com/muesli/termenv"
)

type Terminal struct {
	output termenv.Output
	input  *bufio.Reader
}

func NewTerminal() *Terminal {
	return &Terminal{
		output: *termenv.DefaultOutput(),
		input:  bufio.NewReader(os.Stdout),
	}
}

func (t *Terminal) WriteLine(s string) error {
	_, err := t.output.WriteString(s + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (t *Terminal) ReadLine() (string, error) {
	line, err := t.input.ReadString(10)
	if err != nil {
		return "", err
	}
	return line[:len(line)-1], nil
}

// investigate the Write in the termenv
