package filestore

import (
	"github.com/gin-gonic/gin"

	"github.com/PSauerborn/project-alpha/internal/pkg/filestore"
	db "github.com/PSauerborn/project-alpha/internal/pkg/filestore/persistence"
	"github.com/PSauerborn/project-alpha/internal/pkg/utils"
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
	r.GET("/filestore/file/contents/:fileId", filestore.GetFileHandler)
	r.GET("/filestore/file/meta/:fileId", filestore.GetFileMetadataHandler)

	// define routes used to create and modify files
	r.POST("/filestore/file", filestore.CreateFileHandler)
	r.PUT("/filestore/file/:fileId", filestore.PutFileHandler)
	r.DELETE("/filestore/file/:fileId", filestore.DeleteFileHandler)
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
	basePersistence := utils.NewBasePersistence(connectionString)
	// generate new instance of base persistence
	return &db.PostgresPersistence{
		BasePostgresPersistence: basePersistence,
		BaseFilePath:            basePath,
	}
}