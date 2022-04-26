package pemeriksaan

import "github.com/gin-gonic/gin"

func PemeriksaanController(r *gin.Engine){
	r.POST("/check", CheckDisease)
}