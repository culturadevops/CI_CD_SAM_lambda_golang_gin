package routers

import (
	"hello-world/handles"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	r.GET("/allseason", handles.AllSeasion)
	r.POST("/addseason", handles.AddSeasion)
	r.POST("/delseason/:id", handles.DelSeasion)
	//r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("season/upload", handles.UploadImage)

	r.POST("/addchapter", handles.AddChapter)
	r.POST("/delchapter/:id", handles.DelChapter)

	r.GET("/allchapterbyidseason/:id", handles.AllChapterByIdSeason)
	r.GET("/chapterbyid/:id", handles.ChapterById)
	r.GET("/chapterrandom", handles.ChapterRandom)
	r.GET("/chapterrandombyidseason/:id", handles.ChapterRandomByIdSeason)

}
