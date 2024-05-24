package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink_1, salt_1 := GenerateShortLink(initialLink_1)

	initialLink_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_2, salt_2 := GenerateShortLink(initialLink_2)

	initialLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3, salt_3 := GenerateShortLink(initialLink_3)

	assert.Equal(t, len(shortLink_1), 8)
	assert.Equal(t, len(shortLink_2), 8)
	assert.Equal(t, len(shortLink_3), 8)

	// test that the salts are not empty
	assert.NotEmpty(t, salt_1)
	assert.NotEmpty(t, salt_2)
	assert.NotEmpty(t, salt_3)
}
