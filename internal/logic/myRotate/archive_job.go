package myRotate

import (
	"fmt"
	"path/filepath"
)

type archiveJob struct {
	file     string
	archiver *archiver
}

func newArchiveJob(file string, archiver *archiver) *archiveJob {
	return &archiveJob{
		file:     file,
		archiver: archiver,
	}
}

func (aj *archiveJob) Process() result {
	filename := filepath.Base(aj.file)
	fmt.Printf("Начата архивация файла: %s...\n", filename)

	path, err := aj.archiver.archiveFile(aj.file)
	return result{
		file: filename,
		path: path,
		err:  err,
	}
}
