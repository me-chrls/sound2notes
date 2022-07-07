package upload

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ConfigureUpload(router *gin.Engine) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 16 << 20 // 16MiB
}

func Upload(context *gin.Context, fileKey string) bool {
	// Single file
	file, err := context.FormFile(fileKey)

	// The file cannot be received.
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return false
	}

	// Retrieve file information
	//extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	//newFileName := uuid.New().String() + extension

	filepath := "uploads/" + file.Filename // TODO: generate Unique Filename
	// The file is received, so let's save it
	if err := context.SaveUploadedFile(file, filepath); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		log.Print(err)
		return false
	}

	filepath = "http://localhost:8080/" + filepath // TODO: setup host IP address with port dynamically
	// File saved successfully. Return proper result
	context.JSON(http.StatusOK, gin.H{
		"message":  "Your file has been successfully uploaded.",
		"filepath": filepath,
	})
	return true
}
