package test

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/aws/eks-anywhere/pkg/filewriter"
)

var UpdateGoldenFiles = flag.Bool("update", false, "update golden files")

func AssertFilesEquals(t *testing.T, gotPath, wantPath string) {
	t.Helper()
	gotFile := ReadFile(t, gotPath)
	processUpdate(t, wantPath, gotFile)
	wantFile := ReadFile(t, wantPath)

	if gotFile != wantFile {
		cmd := exec.Command("diff", wantPath, gotPath)
		result, err := cmd.Output()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() == 1 {
					t.Fatalf("Results diff expected actual:\n%s", string(result))
				}
			}
		}
		t.Fatalf("Files are different got =\n  %s \n want =\n %s\n%s", gotFile, wantFile, err)
	}
}

func AssertContentToFile(t *testing.T, gotContent, wantFile string) {
	t.Helper()
	if wantFile == "" {
		return
	}
	processUpdate(t, wantFile, gotContent)

	fileContent := ReadFile(t, wantFile)
	if gotContent != fileContent {
		diff, err := computeDiffBetweenContentAndFile([]byte(gotContent), wantFile)
		if err != nil {
			t.Fatalf("Content doesn't match file got =\n%s\n\n\nwant =\n%s\n", gotContent, fileContent)
		}
		if diff != "" {
			t.Fatalf("Results diff expected actual for %s:\n%s", wantFile, string(diff))
		}
	}
}

func contentEqualToFile(gotContent []byte, wantFile string) (bool, error) {
	if wantFile == "" && len(gotContent) == 0 {
		return false, nil
	}

	fileContent, err := ioutil.ReadFile(wantFile)
	if err != nil {
		return false, err
	}

	return bytes.Equal(gotContent, fileContent), nil
}

func computeDiffBetweenContentAndFile(content []byte, file string) (string, error) {
	cmd := exec.Command("diff", "-u", file, "-")
	cmd.Stdin = bytes.NewReader([]byte(content))
	result, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return string(result), nil
		}

		return "", fmt.Errorf("computing the difference between content and file %s: %v", file, err)
	}
	return "", nil
}

func processUpdate(t *testing.T, filePath, content string) {
	if *UpdateGoldenFiles {
		if err := ioutil.WriteFile(filePath, []byte(content), 0o644); err != nil {
			t.Fatalf("failed to update golden file %s: %v", filePath, err)
		}
		log.Printf("Golden file updated: %s", filePath)
	}
}

func ReadFileAsBytes(t *testing.T, file string) []byte {
	bytesRead, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("File [%s] reading error in test: %v", file, err)
	}

	return bytesRead
}

func ReadFile(t *testing.T, file string) string {
	return string(ReadFileAsBytes(t, file))
}

func NewWriter(t *testing.T) (dir string, writer filewriter.FileWriter) {
	dir, err := ioutil.TempDir(".", SanitizePath(t.Name())+"-")
	if err != nil {
		t.Fatalf("error setting up folder for test: %v", err)
	}

	t.Cleanup(cleanupDir(t, dir))
	writer, err = filewriter.NewWriter(dir)
	if err != nil {
		t.Fatalf("error creating writer with folder for test: %v", err)
	}
	return dir, writer
}

func cleanupDir(t *testing.T, dir string) func() {
	return func() {
		if !t.Failed() {
			os.RemoveAll(dir)
		}
	}
}

var sanitizePathChars = regexp.MustCompile(`[^\w-]`)

const sanitizePathReplacementChar = "_"

// SanitizePath sanitizes s so its usable as a path name. For safety, it assumes all characters that are not
// A-Z, a-z, 0-9, _ or - are illegal and replaces them with _.
func SanitizePath(s string) string {
	return sanitizePathChars.ReplaceAllString(s, sanitizePathReplacementChar)
}
