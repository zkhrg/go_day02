package rotater

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func addFileToTar(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filename)

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	if _, err := io.Copy(tw, file); err != nil {
		return err
	}

	return nil
}

func createTarGz(tarGzFilename string, file string, wg *sync.WaitGroup) error {
	defer wg.Done()
	tarGzFile, err := os.Create(tarGzFilename)
	if err != nil {
		return err
	}
	defer tarGzFile.Close()

	gw := gzip.NewWriter(tarGzFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	if err := addFileToTar(tw, file); err != nil {
		return err
	}

	return nil
}

func LogRotate(dest string, timestamp string, args []string) {
	var wg sync.WaitGroup

	for _, file := range args {
		fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))
		wg.Add(1)
		go createTarGz(dest+"/"+fileName+"_"+timestamp+".tar.gz", file, &wg)
	}
	wg.Wait()
}
