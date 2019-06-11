package handlers

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"log"
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"encoding/hex"
	"urlShotener/cmd/url"
)

type ShortUrlHandlers struct {
	urlModel url.UrlModelInterface;
}

func NewShortUrlHandlers(urlModel url.UrlModelInterface) (*ShortUrlHandlers) {
	return &ShortUrlHandlers { urlModel: urlModel }
}

func (handler *ShortUrlHandlers) SaveShortUrl (context *gin.Context) {

	fullUrl := context.PostForm("url")
	hash := md5.Sum([]byte(fullUrl))
	hashString := hex.EncodeToString(hash[:])

	shortUrl := MakeShortUrl()

	fmt.Printf("short url %+v", shortUrl)

	urlEntity := url.NewUrlEntity(shortUrl, hashString, fullUrl)

	err := handler.urlModel.SaveShortUrl(urlEntity)

	if err != nil {
		log.Fatalf("Failed execute query: %v\n", err)
		context.Status(500)
	}

	context.JSON( 200, gin.H{
		"fullUrl": fullUrl,
		"short": shortUrl,
		"hash": hashString,
	})
}

func (handler *ShortUrlHandlers) GetFullUrl (context *gin.Context) {

	shortUrl := context.Param("short")

	fullUrl, err := handler.urlModel.GetFullUrlByShort(shortUrl)
	if err != nil {
		log.Fatalf("Failed execute query: %v\n", err)
		context.Status(500)
	}

	if (fullUrl != "") {
		context.Redirect(301, fullUrl)
		return
	}

	context.Status(404)
}

func StringShuffle(str string) string {
	runes := []rune(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, len(runes))
	for i, v := range r.Perm(len(runes)) {
		s[i] = runes[v]
	}

	return string(s)
}

func MakeShortUrl() (string) {

	length := 10
	allowSymbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	repeatCount := int(math.Ceil(float64(length)/float64(len(allowSymbols))))
	longString := StringShuffle(strings.Repeat(allowSymbols, repeatCount))

	return longString[0:length]
}

