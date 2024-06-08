package service

import (
	"context"
	"intern_traning/ent"
	"intern_traning/internal/azurestorage"
	"intern_traning/internal/util"
	"net/http"

	"go.uber.org/zap"
)

type StorageService interface {
	CreateAttachmentSASURL(ctx context.Context, input ent.AttachmentInput) (*ent.AttachmentResponse, error)
}

type storageSvcImpl struct {
	objectStorage azurestorage.AzureStorage
	logger        *zap.Logger
}

func NewStorageService(objectStorage azurestorage.AzureStorage, logger *zap.Logger) StorageService {
	return storageSvcImpl{
		objectStorage: objectStorage,
		logger:        logger,
	}
}

func (svc storageSvcImpl) CreateAttachmentSASURL(ctx context.Context, input ent.AttachmentInput) (*ent.AttachmentResponse, error) {
	fileName := input.Folder.String() + "/" + input.ID + input.FileName
	if input.Action.String() == "UPLOAD" {
		url, err := svc.objectStorage.CreateUploadSASURL(ctx, fileName)

		if err != nil {
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagUncategorized)
		}
		return &ent.AttachmentResponse{
			FileName: input.FileName,
			URL:      url,
			Action:   input.Action,
			ID:       input.ID,
		}, nil
	} else {
		url, err := svc.objectStorage.CreateDownloadSASURL(ctx, fileName)
		if err != nil {
			return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagUncategorized)
		}
		return &ent.AttachmentResponse{
			FileName: input.FileName,
			URL:      url,
			Action:   input.Action,
			ID:       input.ID,
		}, nil
	}
}
