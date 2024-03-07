package utils

import (
	"fmt"
	"io"
	"os"
)

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func CopyFile(name, srcpath, destpath string) error {

	content, err := os.ReadFile(srcpath + "\\" + name)
	if err != nil {
		return err

	}

	err = os.WriteFile(destpath+name, content, 0755)
	if err != nil {
		return err

	}

	return nil
}

func CopyDir(src string, dest string) error {

	if dest[:len(src)] == src {
		return fmt.Errorf("cannot copy a folder into the folder itself")
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return fmt.Errorf("Source " + file.Name() + " is not a directory!")
	}

	err = os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {

		if f.IsDir() {

			err = CopyDir(src+"/"+f.Name(), dest+"/"+f.Name())
			if err != nil {
				return err
			}

		}

		if !f.IsDir() {

			content, err := os.ReadFile(src + "/" + f.Name())
			if err != nil {
				return err

			}

			err = os.WriteFile(dest+"/"+f.Name(), content, 0755)
			if err != nil {
				return err

			}

		}

	}

	return nil
}

func MoveDir(src string, dest string) error {

	if dest[:len(src)] == src {
		return fmt.Errorf(`cannot copy a folder into the folder itself`)
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return fmt.Errorf("Source " + file.Name() + " is not a directory!")
	}

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {

		if f.IsDir() {

			err = MoveDir(src+"/"+f.Name(), dest+"/"+f.Name())
			if err != nil {
				return err
			}

		}

		if !f.IsDir() {

			content, err := os.ReadFile(src + "/" + f.Name())
			if err != nil {
				return err

			}

			err = os.WriteFile(dest+"/"+f.Name(), content, 0755)
			if err != nil {
				return err

			}
			/*
				// The copy was successful, so now delete the original file
				err = os.Remove(src + "/" + f.Name())
				if err != nil {
					return fmt.Errorf("Failed removing original file: %s", err)
				}
			*/

		}

	}

	err = os.RemoveAll(src)

	if err != nil {
		return fmt.Errorf("failed removing original path: %s", err)
	}
	return nil
}
