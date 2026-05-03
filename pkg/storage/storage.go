// Package storage defines constructors for kbsink persistence. Types live in core (Storage interface)
// and internal/store; kb-sink-md defaults to local disk via NewLocalStorage.
package storage

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kbsink-org/kbsink/internal/store"
	"github.com/kbsink-org/kbsink/pkg/core"
)

// NewLocalStorage writes output under a filesystem root directory.
func NewLocalStorage(root string) core.Storage {
	return store.NewLocalStorage(root)
}

// NewS3Storage uploads markdown and assets to an S3-compatible bucket.
func NewS3Storage(client *s3.Client, bucket, prefix string) (core.Storage, error) {
	return store.NewS3Storage(client, bucket, prefix)
}
