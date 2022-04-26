package pemeriksaan

import "github.com/gin-gonic/gin"

func PemeriksaanController(r *gin.Engine) {
	r.POST("/check", CheckUp)
	r.GET("/check", GetCheckupHistory)
	r.DELETE("/check/:id", DeleteCheckupHistory)
}
