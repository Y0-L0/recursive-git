package git

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
	"path"
)

func getObject(base string, sha GitSha) (string, error) {
	objectPath := path.Join(base, ".git/objects", string(sha[:2]), string(sha[2:]))
	objectBytes, err := os.ReadFile(objectPath)
	if err != nil {
		return "", err
	}

	reader, err := zlib.NewReader(bytes.NewReader(objectBytes))
	if err != nil {
		return "", err
	}
	defer reader.Close()

	var object bytes.Buffer
	_, err = io.Copy(&object, reader)
	if err != nil {
		return "", err
	}

	return object.String(), nil
}
