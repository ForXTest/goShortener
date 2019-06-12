package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"urlShotener/cmd/url"
)

type UrlModelStub struct {}

func (model *UrlModelStub) GetFullUrlByShort(shortUrl string) (string, error) {

	if shortUrl == "shorturl" {
		return "http://test.ru/", nil
	}

	if shortUrl == "errorurl" {
		return "", errors.New("sql: database is closed")
	}

	return "", nil
}

func (model *UrlModelStub) SaveShortUrl(entity *url.UrlEntity) (error) {

	return nil
}

func TestMakeShortUrl(t *testing.T) {
	test := assert.New(t)

	url := MakeShortUrl();

	test.Equal(ShortUrlLength, len(url));
}

func TestStringShuffle(t *testing.T) {
	test := assert.New(t)

	str := "test string"

	test.NotEqual(str, StringShuffle(str))
}

//func TestSaveShortUrl(t *testing.T) {
//
//	test := assert.New(t)
//
//	client, server := initTestServer()
//
//	defer server.Close()
//
//	response, err := client.Post(
//		fmt.Sprintf("%s/short/", server.URL),
//		gin.MIMEPOSTForm,
//		strings.NewReader("url=http://test.ru/lalalalalalalalala"),
//	)
//
//	defer response.Body.Close()
//
//	test.NoError(err)
//
//	test.Equal(http.StatusOK, response.StatusCode)
//	//test.Equal(http.StatusOK, response.Body.Read())
//}

func TestGetFullUrl(t *testing.T) {

	test := assert.New(t)

	client, server := initTestServer()

	defer server.Close()

	response, err := client.Get(fmt.Sprintf("%s/%s", server.URL, "shorturl"))

	defer response.Body.Close()

	test.NoError(err)
	test.Equal(http.StatusMovedPermanently, response.StatusCode)

	responseNotFound, err := client.Get(fmt.Sprintf("%s/%s", server.URL, "othershorturl"))

	defer responseNotFound.Body.Close()

	test.NoError(err)
	test.Equal(http.StatusNotFound, responseNotFound.StatusCode)

	responseError, err := client.Get(fmt.Sprintf("%s/%s", server.URL, "errorurl"))

	defer responseError.Body.Close()

	test.NoError(err)
	test.Equal(http.StatusInternalServerError, responseError.StatusCode)
}

func initTestServer() (*http.Client, *httptest.Server) {

	urlHandlers := NewShortUrlHandlers(&UrlModelStub{})

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/:short", urlHandlers.GetFullUrl)
	router.POST("/short/", urlHandlers.SaveShortUrl)

	server := httptest.NewServer(router)

	//not follow redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		} ,
	}

	return client, server
}