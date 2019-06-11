package url

import "urlShotener/cmd/database"

type UrlModel struct {
	DB *database.MySQL;
}

type UrlEntity struct {
	shortUrl, hash, fullUrl string
}

func NewUrlEntity(shortUrl string, hash string, fullUrl string) (*UrlEntity) {
	return &UrlEntity { shortUrl: shortUrl, hash: hash, fullUrl: fullUrl }
}

func (model *UrlModel) GetFullUrlByShort(shortUrl string) (string, error) {

	result, err := model.DB.Master.Query("SELECT link FROM short_links WHERE short = ? LIMIT 1", shortUrl);

	defer result.Close()

	if err != nil {
		return "", err
	}

	var fullUrl string

	for result.Next() {
		err = result.Scan(
			&fullUrl,
		)

		if err != nil {
			return "", err
		}
	}

	return fullUrl, nil
}

func (model *UrlModel) SaveShortUrl(entity *UrlEntity) (error) {

	_, err := model.DB.Master.Exec(
		"INSERT INTO short_links (short, hash, link) VALUES (?, ?, ?);",
		entity.shortUrl,
		entity.hash,
		entity.fullUrl,
	);

	if err != nil {
		return err
	}

	return nil
}
