package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatalln("Need a project name")
	}

	moduleName := args[0]

	replaceImport := "github.com/aleksess/nifty/bootstrap"

	fmt.Println("INITIALISING GO MODULE")
	if err := exec.Command("go", "mod", "init", moduleName).Run(); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("COPYING BOOTSTRAP")
	if err := exec.Command("git", "clone", "https://github.com/aleksess/nifty").Run(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := copyDir("./nifty/bootstrap", "."); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("REMOVING nifty dir since no longer needed")
	if err := os.RemoveAll("./nifty"); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("RENAMING bootstrap TO %s\n", moduleName)

	if err := replaceTextInDir(".", replaceImport, moduleName); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("DOWNLOADING DEPS")
	if err := exec.Command("go", "get", "-u", "github.com/aleksess/nifty").Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

// copyDir copies a whole directory recursively, omitting .gitkeep files and the .git directory
func copyDir(src string, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", src, err)
	}

	for _, entry := range entries {
		if entry.Name() == ".gitkeep" || entry.Name() == ".git" {
			continue
		}

		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		fi, err := os.Stat(srcPath)
		if err != nil {
			return fmt.Errorf("failed to stat file %s: %v", srcPath, err)
		}

		switch mode := fi.Mode(); {
		case mode.IsDir():
			if err := os.MkdirAll(dstPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", dstPath, err)
			}
			if err := copyDir(srcPath, dstPath); err != nil {
				return fmt.Errorf("failed to copy directory %s to %s: %v", srcPath, dstPath, err)
			}
		case mode.IsRegular():
			if err := copyFile(srcPath, dstPath); err != nil {
				return fmt.Errorf("failed to copy file %s to %s: %v", srcPath, dstPath, err)
			}
		}
	}

	return nil
}

// copyFile copies a single file from src to dst
func copyFile(src string, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat file %s: %v", src, err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %v", src, err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %v", dst, err)
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return fmt.Errorf("failed to copy content from %s to %s: %v", src, dst, err)
	}

	// Handle empty files
	if sourceFileStat.Size() == 0 {
		if err := destination.Truncate(0); err != nil {
			return fmt.Errorf("failed to truncate file %s: %v", dst, err)
		}
	}

	if err := os.Chmod(dst, sourceFileStat.Mode()); err != nil {
		return fmt.Errorf("failed to set permissions on file %s: %v", dst, err)
	}

	return nil
}

// replaceTextInDir recursively goes through all files in a directory and replaces occurrences of oldText with newText
func replaceTextInDir(dir string, oldText string, newText string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read the file
		input, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", path, err)
		}

		// Replace oldText with newText
		output := strings.ReplaceAll(string(input), oldText, newText)

		// Write the modified content back to the file
		if err := ioutil.WriteFile(path, []byte(output), info.Mode()); err != nil {
			return fmt.Errorf("failed to write file %s: %v", path, err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk the directory %s: %v", dir, err)
	}

	return nil
}
