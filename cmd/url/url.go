package url

import "urlShotener/cmd/database"

type Url struct {
	ShortUrl, Hash, FullUrl string
}

func GetFullUrlByShort(DB *database.MySQL, shortUrl string) (string, error) {

	result, err := DB.Master.Query("SELECT link FROM short_links WHERE short = ? LIMIT 1", shortUrl);

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

func SaveShortUrl(DB *database.MySQL, shortUrl string, hashString string, fullUrl string) (error) {

	_, err := DB.Master.Exec(
		"INSERT INTO short_links (short, hash, link) VALUES (?, ?, ?);",
		shortUrl,
		hashString,
		fullUrl,
	);

	if err != nil {
		return err
	}

	return nil
}
