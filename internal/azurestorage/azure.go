package azurestorage

import (
	"context"
	"fmt"
	"intern_traning/config"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/pkg/errors"
)

const (
	MAXIMUM_FILE_SIZE_MB        int64 = 10
	SAS_MAXIMUM_HOUR_EXPIRATION int64 = 2
)

// AzureStorage is an interface for Azure Storage
type AzureStorage interface {
	CreateUploadSASURL(ctx context.Context, fileName string) (string, error)
	CreateDownloadSASURL(ctx context.Context, fileName string) (string, error)
}

// impl is an implementation of AzureStorage
type impl struct {
	client                   *container.Client
	maximumFileSizeMB        int64
	sasMaximumHourExpiration int64
}

// NewAzureStorage creates a new instance of AzureStorage
func NewAzureStorage(config config.AzureStorageConfig) (AzureStorage, error) {
	cred, err := azblob.NewSharedKeyCredential(config.AccountName, config.AccountKey)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("failed to create shared key credential: %w", err))
	}

	service, err := azblob.NewClientWithSharedKeyCredential(config.ServiceURL, cred, nil)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("failed to create service client: %w", err))
	}

	container := service.ServiceClient().NewContainerClient(config.Container)

	maximumFileSizeMb := MAXIMUM_FILE_SIZE_MB
	if config.MaximumFileSizeMB > 0 {
		maximumFileSizeMb = config.MaximumFileSizeMB
	}

	sasMaximumHourExpiration := SAS_MAXIMUM_HOUR_EXPIRATION
	if config.SASMaximumHourExpiration > 0 {
		sasMaximumHourExpiration = config.SASMaximumHourExpiration
	}

	return impl{
		client:                   container,
		maximumFileSizeMB:        maximumFileSizeMb,
		sasMaximumHourExpiration: sasMaximumHourExpiration,
	}, nil
}

// CreateUploadSASURL creates a SAS URL for uploading a file
func (i impl) CreateUploadSASURL(ctx context.Context, fileName string) (string, error) {
	expireTime := time.Now().UTC().Add(time.Duration(i.sasMaximumHourExpiration) * time.Hour)

	url, err := i.client.NewBlobClient(fileName).GetSASURL(sas.BlobPermissions{
		Write: true,
	}, expireTime, nil)
	if err != nil {
		return "", errors.WithStack(fmt.Errorf("failed to create SAS URL: %w", err))
	}

	return url, nil
}

// CreateDownloadSASURL creates a SAS URL for downloading a file
func (i impl) CreateDownloadSASURL(ctx context.Context, fileName string) (string, error) {
	expireTime := time.Now().UTC().Add(time.Duration(i.sasMaximumHourExpiration) * time.Hour)

	url, err := i.client.NewBlobClient(fileName).GetSASURL(sas.BlobPermissions{
		Read: true,
	}, expireTime, nil)
	if err != nil {
		return "", errors.WithStack(fmt.Errorf("failed to create SAS URL: %w", err))
	}

	return url, nil
}
