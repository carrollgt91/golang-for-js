package main

import (
	"io"
	"os"
)

// Copy copies the file contents from the file
// at path src to the file at path dst. If the src
// file does not exist, it returns an error. If the
// dest file is not writable, it returns an error.
// Does not correctly handle all edge cases.
func Copy(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	// this statement will be executed after this function returns,
	// ensuring that, no matter what happens, this file handler is
	// cleaned up.
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// we can use underscore to indicate we don't want to use
	// a return value
	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	// commit current changes of the file to stable storage
	err = out.Sync()
	return err
}

func main() {
	err := Copy("./data/src.txt", "./data/dest.txt")
	if err != nil {
		panic(err)
	}
}
