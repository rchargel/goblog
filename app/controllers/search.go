package controllers

import (
	"github.com/rchargel/goblog/app/domain"
	"github.com/robfig/revel"
)

var indexer = domain.NewIndexer()

type Search struct {
	*revel.Controller
}

func (c Search) Search(terms string) revel.Result {
	searchResults := indexer.Search(terms)

	return c.Render(terms, searchResults)
}
