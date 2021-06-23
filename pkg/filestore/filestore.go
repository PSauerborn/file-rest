package filestore

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/PSauerborn/file-rest/internal/pkg/filestore"
	db "github.com/PSauerborn/file-rest/internal/pkg/filestore/persistence"
	"github.com/PSauerborn/file-rest/internal/pkg/utils"
)

// function used to generate new API instance
func NewFilestoreAPI(persistence filestore.FileStorePersistence) *gin.Engine {
	// set instance of filde persistence globally in module
	filestore.SetFilePersistence(persistence)

	// generate new gin router with default middleware
	r := gin.Default()
	r.GET("/filestore/health", filestore.HealthCheckHandler)
	// define routes used to retrieve files
	r.GET("/filestore/files", filestore.ListFilesHandler)
	r.GET("/filestore/file/:fileId/content", filestore.GetFileHandler)
	r.GET("/filestore/file/:fileId/meta", filestore.GetFileMetadataHandler)

	// define routes used to create and modify files
	r.POST("/filestore/file", filestore.CreateFileHandler)
	r.PUT("/filestore/file/:fileId", filestore.PutFileHandler)
	r.PUT("/filestore/file/:fileId/archive", filestore.ArchiveFileHandler)
	r.DELETE("/filestore/file/:fileId", filestore.DeleteFileHandler)

	r.POST("/filestore/search", filestore.SearchFilesHandler)
	return r
}

// function used to generate new instance of API accessor
func NewAccessor(host, protocol string, port int) filestore.FileStoreAPIAccessor {
	baseAccessor := utils.NewBaseAccessor(host, protocol, port)
	return filestore.FileStoreAPIAccessor{
		BaseAPIAccessor: baseAccessor,
	}
}

// function used to generate new instance of postgres persistence
func NewPostgresPersistence(connectionString, basePath string) *db.PostgresPersistence {
	directories := []string{
		basePath,
		fmt.Sprintf("%s/archive", basePath),
	}
	// generate required directories st start time
	for _, dir := range directories {
		os.MkdirAll(dir, os.ModePerm)
	}
	basePersistence := utils.NewBasePersistence(connectionString)
	// generate new instance of base persistence
	return &db.PostgresPersistence{
		BasePostgresPersistence: basePersistence,
		BaseFilePath:            basePath,
	}
}
