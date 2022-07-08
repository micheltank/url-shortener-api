package domain

import "github.com/lithammer/shortuuid/v4"

type Url struct {
	hash     string
	longUrl  string
	shortUrl string
}

func NewUrl(longUrl string) Url {
	hash := shortuuid.NewWithNamespace(longUrl)
	return Url{
		hash:     hash,
		longUrl:  longUrl,
		shortUrl: "", // todo: build
	}
}

func NewShortenedUrl(hash, longUrl string) Url {
	return Url{
		hash:     hash,
		longUrl:  longUrl,
		shortUrl: "", // todo: build
	}
}

func (u Url) GetShortUrl() string {
	return u.shortUrl
}

func (u Url) GetLongUrl() string {
	return u.longUrl
}
