package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// ZipReader handles reading from ZIP archives
type ZipReader struct {
	zipFile *os.File
	reader  *zip.Reader
}

// NewZipReader creates a new ZipReader instance
func NewZipReader(zipPath string) (*ZipReader, error) {
	zipFile, err := os.Open(zipPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open zip file: %w", err)
	}

	fileInfo, err := zipFile.Stat()
	if err != nil {
		zipFile.Close()
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	reader, err := zip.NewReader(zipFile, fileInfo.Size())
	if err != nil {
		zipFile.Close()
		return nil, fmt.Errorf("failed to create zip reader: %w", err)
	}

	return &ZipReader{
		zipFile: zipFile,
		reader:  reader,
	}, nil
}

// GetFirstFile returns a reader for the first file in the archive
func (z *ZipReader) GetFirstFile() (io.ReadCloser, error) {
	if len(z.reader.File) == 0 {
		return nil, fmt.Errorf("empty zip archive")
	}

	return z.reader.File[0].Open()
}

// Close closes the zip file
func (z *ZipReader) Close() error {
	return z.zipFile.Close()
}
