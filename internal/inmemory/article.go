// TODO: это вымышленный пример, удалите в реальном приложении
package inmemory

import (
	"errors"
	"go.citilink.cloud/grpc-skeleton/internal"
	"sort"
	"sync"
)

type ArticleStorage struct {
	mu       sync.RWMutex
	articles map[internal.ArticleId]*internal.Article
	lastId   internal.ArticleId
}

func NewArticleStorage() *ArticleStorage {
	return &ArticleStorage{
		articles: make(map[internal.ArticleId]*internal.Article),
	}
}

func (a *ArticleStorage) Create(title string, content string, category internal.Category, tags []string, isVisible bool) (*internal.Article, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	article := &internal.Article{
		Id:        a.lastId + 1,
		Title:     title,
		Content:   content,
		Category:  category,
		Tags:      tags,
		IsVisible: isVisible,
	}

	a.articles[article.Id] = article
	a.lastId = article.Id

	return article, nil
}

func (a *ArticleStorage) Update(id internal.ArticleId, title string, content string, category internal.Category, tags []string, isVisible bool) (*internal.Article, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	article, ok := a.articles[id]
	if !ok {
		return nil, errors.New("article doesn't exist")
	}

	article.Title = title
	article.Content = content
	article.Category = category
	article.Tags = tags
	article.IsVisible = isVisible

	return article, nil
}

func (a *ArticleStorage) Get(id internal.ArticleId) (*internal.Article, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.articles[id], nil
}

func (a *ArticleStorage) Filter(categories []internal.Category, tags []string, onlyVisible bool) ([]*internal.Article, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	articles := make([]*internal.Article, 0)
	for _, article := range a.articles {
		categoryFound := false
		for _, category := range categories {
			if category == article.Category {
				categoryFound = true
			}
		}

		tagFound := false
		for _, tag := range tags {
			for _, articleTag := range article.Tags {
				if articleTag == tag {
					tagFound = true
				}
			}
		}

		if (len(categories) == 0 || categoryFound) && (len(tags) == 0 || tagFound) && (!onlyVisible || article.IsVisible) {
			articles = append(articles, article)
		}
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Id < articles[j].Id
	})

	return articles, nil
}

func (a *ArticleStorage) Delete(id internal.ArticleId) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	delete(a.articles, id)

	return nil
}
