package handles

import (
	"hello-world/libs"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Name   string `form:"Name" json:"Name" binding:"required"`
	Imagen string `form:"Image" json:"Image" binding:"required"`
}

func AddSeasion(c *gin.Context) {
	var json data
	c.BindJSON(&json)
	x, err := season.Add(json.Name, json.Imagen)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, x)
	}

}

func DelSeasion(c *gin.Context) {
	x := season.Del(GetParamID(c))
	c.JSON(http.StatusOK, gin.H{
		"mylist": x,
	})
}
func ls(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ls:")
	for _, f := range files {
		log.Println(f.Name())
	}
}

var filepath string

func UploadImage(c *gin.Context) {
	s3 := new(libs.S3Client)
	s3.NewSession("us-east-1")
	_, Header, err := c.Request.FormFile("photo")
	filename := Header.Filename
	log.Println("intento de copiar")
	//log.Println(c.SaveUploadedFile(Header, "/tmp/"+Header.Filename))
	log.Println(c.SaveUploadedFile(Header, "./tmp/"+Header.Filename))
	//upload to the s3 bucket
	MyBucket := "thesimpson"
	ls("./tmp/")
	err = s3.AddFilesToS3("./tmp/"+Header.Filename, MyBucket, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to upload file",
			"uploader": err,
		})
		return
	}

	filepath = filename
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})
}

/*func UpdateSeasion(c *gin.Context) {
	//https://medium.com/wesionary-team/aws-sdk-for-go-and-uploading-a-file-using-s3-bucket-df7425317a40
	ls()
	log.Println("update seasionlisto")
	s3 := new(libs.S3Client)
	s3.NewSession("us-east-1")
	file, Header, _ := c.Request.FormFile("file")
	print(Header)
	name := c.PostForm("name")
	finalname := strings.Join(strings.Fields(name), "")
	log.Println(file)
	//log.Println(c.SaveUploadedFile(file, "/tmp/"+Header.Filename))
	content, err := ioutil.ReadFile("/tmp/" + Header.Filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	//s3.Upload("/tmp/"+file.Filename, "thesimpson", "imagentemporadas/"+finalname+filepath.Ext(file.Filename))

	//s3.AddFilesToS3(f, size, "thesimpson", "imagentemporadas/"+finalname+filepath.Ext(file.Filename))
	//s3.Upload("/tmp/"+file.Filename, "thesimpson", "imagentemporadas/"+finalname+filepath.Ext(file.Filename))
	//x, _ := season.Add(name, "/imagentemporadas/"+finalname+filepath.Ext(file.Filename))
	c.JSON(http.StatusOK, x)

}*/
