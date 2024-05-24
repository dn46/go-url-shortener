package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://gin-gonic.com/docs/examples/bind-single-binary-with-template/"
	shortURL := "Jsz4k57oAX"
	salt := "randomSalt"

	SaveUrlMapping(shortURL, initialLink, salt)

	retrievedUrl := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
