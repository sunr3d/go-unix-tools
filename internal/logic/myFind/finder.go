package myFind

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type finder struct {
	opts *options
}

func newFinder(args []string) (*finder, error) {
	opts, err := newOptions(args)
	if err != nil {
		return nil, fmt.Errorf("newOptions: %w", err)
	}
	return &finder{opts: opts}, nil
}

func (f *finder) printSymlink(path string) {
	target, err := os.Readlink(path)
	if err != nil {
		fmt.Printf("%s -> [broken]\n", path)
		return
	}

	if !filepath.IsAbs(target) {
		target = filepath.Join(filepath.Dir(path), target)
	}

	if _, err = os.Stat(target); err != nil {
		fmt.Printf("%s -> [broken]\n", path)
	} else {
		fmt.Printf("%s -> %s\n", path, target)
	}
}

func (f *finder) matchesExtension(path string) bool {
	if f.opts.extFilter == "" {
		return true
	}
	ext := filepath.Ext(path)
	return len(ext) > 1 && ext[1:] == f.opts.extFilter
}

func (f *finder) scanDirectory() error {
	return filepath.WalkDir(f.opts.filepath, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			if errors.Is(err, os.ErrPermission) {
				return nil
			}
			return fmt.Errorf(errWalkDir, f.opts.filepath, err)
		}

		if entry.Type()&os.ModeSymlink != 0 {
			if f.opts.showLinks {
				f.printSymlink(path)
			}
			return nil
		}

		isDir := entry.IsDir()
		if (isDir && f.opts.showDirs) ||
			(!isDir && f.opts.showFiles && f.matchesExtension(path)) {
			fmt.Println(path)
		}

		return nil
	})
}

func (f *finder) run() error {
	info, err := os.Stat(f.opts.filepath)
	if err != nil {
		return fmt.Errorf(errFileInfo, f.opts.filepath, err)
	}

	if !info.IsDir() {
		return fmt.Errorf(errInvalidPath, f.opts.filepath)
	}

	return f.scanDirectory()
}
