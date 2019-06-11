package url

type UrlModelInterface interface {
	GetFullUrlByShort(shortUrl string) (string, error)
	SaveShortUrl(entity *UrlEntity) (error)
}
