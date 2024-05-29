package home

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		_ = Index().Render(context.Background(), w)
		w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World!", doc.Find("h1").Text())
}
