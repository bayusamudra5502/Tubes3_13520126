package controller

import (
	"mime/multipart"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func FileValidation(c *gin.Context, file *multipart.FileHeader) (string, bool) {
		// File Validation
		f, err := file.Open()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": "KO",
				"message": "Error when opening file",
				"data": err,
			})
			return "", false
		}
	
		data := make([]byte, file.Size)
		_, err = f.Read(data)
	
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": "KO",
				"message": "Error when reading file",
				"data": err,
			})
			return "", false
		}
	
		regex, err := regexp.Compile("^[ACTG]+$")
	
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": "KO",
				"message": "Error when creating regular expression",
				"data": err,
			})
			return "", false
		}
	
		if !regex.Match(data) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "KO",
				"message": "DNA Sequence is not valid",
				"data": nil,
			})
			return "", false
		}

		return string(data), true
}