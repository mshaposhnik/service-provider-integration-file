package main

import (
	"context"
	"io"
)

type Provider interface {
	accept(repository string) bool
	fetch(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error)
}
