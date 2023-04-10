package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"path"
	"github.com/yanyiwu/gojieba"
)

func init() {
}

func main() {
	_file, _ := exec.LookPath(os.Args[0])
	_pwd, _ := path.Split(_file)
	os.Chdir(_pwd)

	bindaddr := flag.String("b", "0.0.0.0:80", "listen port")
	flag.Parse()

	x := gojieba.NewJieba(
		"dict/jieba.dict.utf8",
		"dict/hmm_model.utf8",
		"dict/user.dict.utf8",
		"dict/idf.utf8",
		"dict/stop_words.utf8",
	)
	defer x.Free()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/words", func(c *gin.Context) {
		key := c.DefaultQuery("key", "")
		words := x.CutForSearch(key,true)
		c.JSON(200, gin.H{
			"message": "pong",
			"words":   words,
		})
	})
	r.GET("/reload", func(c *gin.Context) {
		x.Free();
		x = gojieba.NewJieba(
			"dict/jieba.dict.utf8",
			"dict/hmm_model.utf8",
			"dict/user.dict.utf8",
			"dict/idf.utf8",
			"dict/stop_words.utf8",
		)
		c.JSON(200, gin.H{
			"message": "pong",
			"reload":  true,
		})
	})
	log.Println("notice: bind addr:%v", *bindaddr)
	err := r.Run(*bindaddr)
	log.Println("error: %v", err)
}
