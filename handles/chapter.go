package handles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type chapter struct {
	Name      string `form:"Name" json:"Name" `
	Imagen    string `form:"Image" json:"Image" `
	Id_season uint   `form:"Id_season" json:"Id_season" `
	Resourse  string `form:"Resourse" json:"Resourse" `
	Order     uint   `form:"Order" json:"Order" `
}

func AddChapter(c *gin.Context) {
	var json chapter
	c.BindJSON(&json)
	x, _ := chapters.Add(json.Id_season, json.Name, json.Imagen, json.Resourse, json.Order)

	c.JSON(http.StatusOK, x)

}

func DelChapter(c *gin.Context) {
	x := chapters.Del(GetParamID(c))
	c.JSON(http.StatusOK, gin.H{
		"mylist": x,
	})
}
