package examples

import (
	"testing"
)

func TestForSelectTimeout(t *testing.T) {
	ForSelectVisitURL([]string{
		"https://github.com",
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com/",
	})
}
