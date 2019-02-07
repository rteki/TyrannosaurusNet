package FileSystemSerializer

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Serialize(rootPath string, compressor io.WriteCloser) error {

	if _, err := os.Stat(rootPath); err != nil {
		fmt.Printf("Can't serialize.")
		return err
	}

	tarWriter := tar.NewWriter(compressor)
	defer tarWriter.Close()

	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())

		if err != nil {
			return err
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		file, err := os.Open(path)

		if err != nil {
			return err
		}

		if _, err := io.Copy(tarWriter, file); err != nil {
			return err
		}

		file.Close()

		return nil
	})

}

func Deserialize() {

}
