package main

import (
	"crypto/rand"
	"os"
)

type Config struct {
	Iter int
	Zero bool
}

func (conf Config) DeleteFile(path string) error {

	//filling the file with junk data
	for i := 0; i < conf.Iter; i++ {
		if err := fillJunk(path, true); err != nil {
			return err
		}
	}

	//filling the file with zeros
	if conf.Zero {
		if err := fillJunk(path, false); err != nil {
			return err
		}
	}

	//removing the file from filesystem
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func fillJunk(path string, random bool) error {
	f, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return err
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	buff := make([]byte, info.Size())
	if random {
		if _, err := rand.Read(buff); err != nil {
			return err
		}
	}

	_, err = f.WriteAt(buff, 0)
	return err
}
