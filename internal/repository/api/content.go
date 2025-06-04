package repository

import (
	"encoding/json"
	"errors"
	"net/http"
)

type QuoteRepository interface {
	GetRandomQuote() (string, string, string, error)
}

type quoteRepository struct{}

func NewQuoteRepository() QuoteRepository {
	return &quoteRepository{}
}

func (r *quoteRepository) GetRandomQuote() (string, string, string, error) {
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	var result []map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", "", err
	}

	if len(result) == 0 {
		return "", "", "", errors.New("empty response from quote API")
	}

	return result[0]["q"], result[0]["a"], result[0]["h"], nil
}
