package cliutil

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// FindRepoRoot walks up from start and returns the first directory containing go.mod.
func FindRepoRoot(start string) (string, error) {
	d := start
	for {
		if _, err := os.Stat(filepath.Join(d, "go.mod")); err == nil {
			return d, nil
		}
		nd := filepath.Dir(d)
		if nd == d { // reached filesystem root
			return "", errors.New("go.mod not found up the tree")
		}
		d = nd
	}
}

// ModulePath reads the go.mod in the repository root and returns the declared module path.
func ModulePath(repoRoot string) (string, error) {
	gomod := filepath.Join(repoRoot, "go.mod")
	f, err := os.Open(gomod)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	return "", errors.New("module path not found in go.mod")
}
