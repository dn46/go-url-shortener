package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://gin-gonic.com/docs/examples/html-rendering/"
	shortLink_1, salt_1 := GenerateShortLink(initialLink_1)

	initialLink_2 := "https://gin-gonic.com/docs/examples/bind-form-data-request-with-custom-struct/"
	shortLink_2, salt_2 := GenerateShortLink(initialLink_2)

	initialLink_3 := "https://gin-gonic.com/docs/examples/bind-single-binary-with-template/"
	shortLink_3, salt_3 := GenerateShortLink(initialLink_3)

	assert.Equal(t, len(shortLink_1), 8)
	assert.Equal(t, len(shortLink_2), 8)
	assert.Equal(t, len(shortLink_3), 8)

	// test that the salts are not empty
	assert.NotEmpty(t, salt_1)
	assert.NotEmpty(t, salt_2)
	assert.NotEmpty(t, salt_3)
}
