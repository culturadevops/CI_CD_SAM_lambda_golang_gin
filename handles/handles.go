package handles

import (
	"hello-world/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

var chapters *models.ChaptersModel
var season *models.SeasonModel

type User struct {
}
type seasonResponde struct {
	ID         uint   `form:"ID" json:"ID" `
	Name       string `form:"Name" json:"Name" `
	Image      string `form:"Image" json:"Image" `
	ImageFinal string `form:"ImageFinal" json:"ImageFinal" `
	Chapters   uint   `form:"Chapters" json:"Chapters" `
}

type chapterResponde struct {
	ID            uint   `form:"ID" json:"ID" `
	Name          string `form:"Name" json:"Name" `
	Image         string `form:"Image" json:"Image" `
	ImageFinal    string `form:"ImageFinal" json:"ImageFinal" `
	Order         uint   `form:"Chapters" json:"Chapters" `
	Resourse      string `form:"Resourse" json:"Resourse" `
	ResourseFinal string `form:"ResourseFinal" json:"ResourseFinal" `
	Id_season     uint   `form:"Id_season" json:"Id_season" `
}

func GetLink(myBucket string, myKey string) string {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myKey),
	})
	urlStr, _ := req.Presign(15 * time.Minute)
	return urlStr
}

func GetParamID(c *gin.Context) uint {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	return uint(id)
}
func responseChapter(c *gin.Context, x models.ChaptersModel) {

	/*c.JSON(http.StatusOK, gin.H{
		"Resourse":  "https://thesimpson.s3.amazonaws.com/" + x.Resourse,
		"Name":      x.Name,
		"Imagen":    "https://thesimpson.s3.amazonaws.com/" + x.Image,
		"Order":     x.Order,
		"Id_season": x.Id_season,
	})*/

	c.JSON(http.StatusOK, gin.H{

		"Name":          x.Name,
		"ImagenFinal":   GetLink("thesimpson", x.Image),
		"Imagen":        x.Image,
		"Resourse":      x.Resourse,
		"ResourseFinal": GetLink("thesimpson", x.Resourse),
		"Order":         x.Order,
		"Id_season":     x.Id_season,
	})

}
func random(numero uint) uint {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return uint(r1.Intn(int(numero)))
}
func ChapterRandomById(c *gin.Context, Id_season uint) {
	sesion := season.SeasonById(Id_season)
	x := chapters.ChapterById(random(sesion.Chapters))
	responseChapter(c, x)
}
func ChapterRandomByIdSeason(c *gin.Context) {
	ChapterRandomById(c, GetParamID(c))
}
func ChapterRandom(c *gin.Context) {
	ChapterRandomById(c, random(2))
}
func ChapterById(c *gin.Context) {
	x := chapters.ChapterById(GetParamID(c))
	responseChapter(c, x)
}

func AllChapterByIdSeason(c *gin.Context) {
	x := chapters.ListAllByIdSeason(GetParamID(c))
	var dataRequest []chapterResponde
	for _, x := range x {
		var data1 chapterResponde
		data1.ID = x.ID
		data1.Name = x.Name
		data1.ImageFinal = GetLink("thesimpson", x.Image)
		data1.Image = x.Image
		data1.ResourseFinal = GetLink("thesimpson", x.Resourse)
		data1.Resourse = x.Resourse
		data1.Order = x.Order
		data1.Id_season = x.Id_season
		dataRequest = append(dataRequest, data1)
	}
	c.JSON(http.StatusOK, gin.H{

		"mylist": dataRequest,
	})
}
func AllSeasion(c *gin.Context) {
	var season *models.SeasonModel
	x := season.ListAll()
	var dataRequest []seasonResponde

	for _, x := range x {
		var data1 seasonResponde
		data1.ID = x.ID
		data1.Name = x.Name
		data1.Image = x.Image
		data1.ImageFinal = GetLink("thesimpson", x.Image)
		data1.Chapters = x.Chapters
		dataRequest = append(dataRequest, data1)
	}
	c.JSON(http.StatusOK, gin.H{

		"mylist": dataRequest,
	})
}
