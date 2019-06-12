package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"urlShotener/cmd/url"
)

const ShortUrlLength = 10;

const AllowSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

type ShortUrlHandlers struct {
	urlModel url.UrlModelInterface;
}

func NewShortUrlHandlers(urlModel url.UrlModelInterface) (*ShortUrlHandlers) {
	return &ShortUrlHandlers { urlModel: urlModel }
}

func (handler *ShortUrlHandlers) SaveShortUrl (context *gin.Context) {

	fullUrl := context.PostForm("url")

	if fullUrl == "" {
		responseBadRequest( context,"Not passed or empty required parameter 'url'")
		return
	}

	hash := GetSringHash(fullUrl)
	shortUrl := MakeShortUrl()

	urlEntity := url.NewUrlEntity(shortUrl, hash, fullUrl)

	err := handler.urlModel.SaveShortUrl(urlEntity)

	if err != nil {
		log.Printf("Failed execute query: %v\n", err)
		responseServerError(context)
		return
	}

	context.JSON( http.StatusOK, gin.H{
		"fullUrl": fullUrl,
		"short":   shortUrl,
		"hash":    hash,
	})
}

func (handler *ShortUrlHandlers) GetFullUrl (context *gin.Context) {

	shortUrl := context.Param("short")

	fullUrl, err := handler.urlModel.GetFullUrlByShort(shortUrl)
	if err != nil {
		log.Printf("Failed execute query: %v\n", err)
		responseServerError(context)
		return
	}

	if (fullUrl != "") {
		context.Redirect(http.StatusMovedPermanently, fullUrl)
		return
	}

	responseNotFound(context)
}

func GetSringHash(str string) string {

	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
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

	repeatCount := int(math.Ceil(float64(ShortUrlLength)/float64(len(AllowSymbols))))
	longString := StringShuffle(strings.Repeat(AllowSymbols, repeatCount))

	return longString[0:ShortUrlLength]
}

