package myRotate

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"Go_Day02/pkg/workerpool"
)

type result struct {
	file string
	path string
	err  error
}

type archiver struct {
	opts *options
}

func newArchiver(args []string) (*archiver, error) {
	opts, err := newOptions(args)
	if err != nil {
		return nil, fmt.Errorf("newOptions: %w", err)
	}
	return &archiver{opts: opts}, nil
}

func (ar *archiver) archiveFile(file string) (string, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return "", fmt.Errorf(errFileInfo, file, err)
	}

	baseName := strings.TrimSuffix(filepath.Base(file), ".log")
	archivePath := filepath.Join(ar.opts.archiveDir,
		fmt.Sprintf("%s_%d.tar.gz", baseName, fileInfo.ModTime().Unix()))

	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return "", fmt.Errorf(errCreateTar, file, err)
	}
	defer archiveFile.Close()

	gzWriter := gzip.NewWriter(archiveFile)
	defer gzWriter.Close()
	twWriter := tar.NewWriter(gzWriter)
	defer twWriter.Close()

	srcFile, err := os.Open(file)
	if err != nil {
		return "", fmt.Errorf(errFileAccess, file, err)
	}
	defer srcFile.Close()

	header := &tar.Header{
		Name:    filepath.Base(file),
		Size:    fileInfo.Size(),
		Mode:    int64(fileInfo.Mode()),
		ModTime: fileInfo.ModTime(),
	}

	if err = twWriter.WriteHeader(header); err != nil {
		return "", fmt.Errorf(errTarHeader, file, err)
	}

	if _, err = io.Copy(twWriter, srcFile); err != nil {
		return "", fmt.Errorf(errArchiveFile, file, err)
	}

	return archivePath, nil
}

func (ar *archiver) run() error {
	filesLen := len(ar.opts.files)
	pool := workerpool.New[result](ar.opts.workerCount, filesLen)
	pool.Start()

	for _, file := range ar.opts.files {
		pool.Submit(newArchiveJob(file, ar))
	}

	hasErrors := false
	opCounter := 0
	for i := 0; i < filesLen; i++ {
		res := <-pool.Results()
		if res.err != nil {
			fmt.Printf("Ошибка при архивации файла \"%s\": %v\n", res.file, res.err)
			hasErrors = true
		} else {
			opCounter++
			fmt.Printf("Файл \"%s\" успешно архивирован в \"%s\"\n", res.file, res.path)
		}
	}

	pool.Stop()

	if hasErrors {
		fmt.Println("При архивации отдельных файлов произошла ошибка.")
	} else {
		fmt.Println("Все файлы успешно архивированы.")
	}
	fmt.Printf("\nИтого:\n- Всего файлов: %d\n- Успешно: %d\n- C ошибками: %d\n",
		filesLen, opCounter, filesLen-opCounter)

	return nil
}
