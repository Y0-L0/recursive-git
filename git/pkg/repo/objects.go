package repo

import (
	"bytes"
	"compress/zlib"
	"io"
	"log/slog"
	"os"
	"path"
)

func getObject(base string, sha GitSha) (string, error) {
	objectPath := path.Join(base, ".git/objects", string(sha[:2]), string(sha[2:]))
	objectBytes, err := os.ReadFile(objectPath)
	if err != nil {
		slog.Error("Failed to read git object file", "sha", sha)
		return "", err
	}

	reader, err := zlib.NewReader(bytes.NewReader(objectBytes))
	if err != nil {
		slog.Error("Failed to initialize zlib reader", "sha", sha)
		return "", err
	}
	defer reader.Close()

	var byteObject bytes.Buffer
	_, err = io.Copy(&byteObject, reader)
	if err != nil {
		slog.Error("Failed to decompress git object", "sha", sha)
		return "", err
	}

	object := byteObject.String()

	slog.Debug("Read git object from file", "sha", sha, "object", object)
	return object, nil
}
