package pemeriksaan

import "github.com/gin-gonic/gin"

func PemeriksaanController(r *gin.Engine) {
	r.POST("/check", CheckUp)
	r.GET("/check", GetCheckupHistory)
	r.GET("/check/:id", GetCheckUpById)
	r.DELETE("/check/:id", DeleteCheckupHistory)
}
