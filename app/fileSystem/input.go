package fileSystem

import (
	"bufio"
	"io"
)

type Input struct {
	scanner    *bufio.Scanner
	line       string
	lineNumber int
	charIndex  int
	isEOF      bool
	isEOLN     bool
}

func CreateInput(reader io.Reader) *Input {
	result := &Input{
		lineNumber: 0,
		charIndex:  0,
	}

	result.scanner = bufio.NewScanner(reader)
	result.isEOF = false
	result.isEOLN = true
	result.line = ""
	return result
}

func (i *Input) Scan() bool {
	if i.isEOLN {
		return false
	}

	i.charIndex++
	if i.charIndex >= len(i.line) {
		i.isEOLN = true
	}

	return !i.isEOLN
}

func (i *Input) ScanLn() bool {
	if !i.Scan() {
		return i.scanLn()
	}

	return true
}

func (i *Input) scanLn() bool {
	if i.isEOF {
		return false
	}

	i.isEOF = !i.scanner.Scan()
	i.line = i.scanner.Text() + "\n"
	i.lineNumber++
	i.charIndex = 0
	i.isEOLN = i.charIndex >= len(i.line)
	if i.isEOLN {
		return i.scanLn()
	}

	return true
}

func (i *Input) ScanBack(numberOfChar int) int {
	i.charIndex = max(0, i.charIndex-numberOfChar)
	i.isEOLN = false
	return i.charIndex
}

func (i *Input) Byte() byte {
	if i.isEOLN {
		return 0
	}
	return i.line[i.charIndex]
}

func (i *Input) IsEOF() bool {
	return i.isEOLN && i.isEOF
}

func (i *Input) IsEOLN() bool {
	return i.isEOLN
}

func (i *Input) GetLineNumber() int {
	return i.lineNumber
}

func (i *Input) GetCharIndex() int {
	return i.charIndex
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
