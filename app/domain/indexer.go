package domain

import (
	"regexp"
	"sort"
	"strings"
	"time"
)

func NewIndexer() *Indexer {
	i := &Indexer{blogList: NewBlogList(), projectList: NewProjectList()}
	i.Init()
	return i
}

type Indexer struct {
	blogList    *BlogList
	projectList *ProjectList
}

type SearchResult struct {
	BlogResults    []SearchResultItem
	ProjectResults []SearchResultItem
}

type SearchResultItem struct {
	Name     string
	Abstract string
	Link     string
	Type     string
	Date     time.Time
	Score    int
}

type ByScore []SearchResultItem

func (l ByScore) Len() int      { return len(l) }
func (l ByScore) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l ByScore) Less(i, j int) bool {
	if l[i].Score != l[j].Score {
		return l[i].Score > l[j].Score
	}
	return l[i].Name < l[j].Name
}

func (i *Indexer) Init() {
	go i.blogList.Index()
	go i.projectList.Index()
}

func (b *BlogList) Search(terms []string, ch chan SearchResultItem) {
	for _, item := range b.items {
		if item.Active {
			score := 0
			for _, term := range terms {
				for key, _ := range item.Index {
					if strings.Index(key, term) != -1 {
						score += item.Index[key]
					}
				}
			}
			if score > 0 {
				ch <- SearchResultItem{
					Name:     item.Name,
					Abstract: item.Abstract,
					Link:     "blog/" + item.Slug,
					Type:     "blog",
					Date:     item.Date.Time,
					Score:    score,
				}
			}
		}
	}
	close(ch)
}

func (b *ProjectList) Search(terms []string, ch chan SearchResultItem) {
	for _, item := range b.Items {
		if item.Active {
			score := 0
			for _, term := range terms {
				for key, _ := range item.Index {
					if strings.Index(key, term) != -1 {
						score += item.Index[key]
					}
				}
			}
			if score > 0 {
				ch <- SearchResultItem{
					Name:     item.Name,
					Abstract: item.Abstract,
					Link:     "proj/" + item.Slug,
					Type:     "project",
					Score:    score,
				}
			}
		}
	}
	close(ch)
}

func getResults(ch chan SearchResultItem) []SearchResultItem {
	items := make([]SearchResultItem, 0, 10)

	for item := range ch {
		if item.Score > 0 {
			items = append(items, item)
		}
	}

	sort.Sort(ByScore(items))
	return items
}

func (i *Indexer) Search(term string) SearchResult {
	stripChars, _ := regexp.Compile("[^\\w|\\s]")
	terms := strings.Fields(strings.ToLower(stripChars.ReplaceAllString(term, "")))
	projChannel := make(chan SearchResultItem)
	blogChannel := make(chan SearchResultItem)

	go i.projectList.Search(terms, projChannel)
	go i.blogList.Search(terms, blogChannel)

	return SearchResult{
		BlogResults:    getResults(blogChannel),
		ProjectResults: getResults(projChannel),
	}
}
