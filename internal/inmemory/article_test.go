package inmemory

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.citilink.cloud/grpc-skeleton/internal"
	"testing"
)

func TestNewArticleStorage(t *testing.T) {
	storage := NewArticleStorage()
	assert.Len(t, storage.articles, 0)
	assert.Zero(t, storage.lastId)
}

func TestArticleStorage_Create(t *testing.T) {
	tests := []struct {
		name             string
		storageArticles  map[internal.ArticleId]*internal.Article
		storageLastId    internal.ArticleId
		title            string
		content          string
		category         internal.Category
		tags             []string
		isVisible        bool
		expectedArticle  *internal.Article
		expectedArticles map[internal.ArticleId]*internal.Article
		expectedLastId   internal.ArticleId
	}{{
		name:            "empty storage",
		storageArticles: map[internal.ArticleId]*internal.Article{},
		storageLastId:   0,
		title:           "title",
		content:         "Content",
		category:        1,
		tags:            []string{"tag1", "tag2"},
		isVisible:       false,
		expectedArticle: &internal.Article{
			Id:        1,
			Title:     "title",
			Content:   "Content",
			Category:  internal.Category(1),
			Tags:      []string{"tag1", "tag2"},
			IsVisible: false,
		},
		expectedArticles: map[internal.ArticleId]*internal.Article{
			1: {Id: 1, Title: "title", Content: "Content", Category: 1, Tags: []string{"tag1", "tag2"}, IsVisible: false},
		},
		expectedLastId: 1,
	}, {
		name: "not empty storage",
		storageArticles: map[internal.ArticleId]*internal.Article{
			1: {Id: 1, Title: "Title11", Content: "Content11", Category: 1, Tags: []string{"tag1_1"}, IsVisible: false},
			2: {Id: 2, Title: "title2", Content: "content3", Category: 22, Tags: []string{"tag2_1", "2_2"}, IsVisible: true},
		},
		storageLastId: 6,
		title:         "Title7",
		content:       "content 77",
		category:      777,
		tags:          []string{"tag77"},
		isVisible:     false,
		expectedArticle: &internal.Article{
			Id:        7,
			Title:     "Title7",
			Content:   "content 77",
			Category:  internal.Category(777),
			Tags:      []string{"tag77"},
			IsVisible: false,
		},
		expectedArticles: map[internal.ArticleId]*internal.Article{
			1: {Id: 1, Title: "Title11", Content: "Content11", Category: 1, Tags: []string{"tag1_1"}, IsVisible: false},
			2: {Id: 2, Title: "title2", Content: "content3", Category: 22, Tags: []string{"tag2_1", "2_2"}, IsVisible: true},
			7: {Id: 7, Title: "Title7", Content: "content 77", Category: 777, Tags: []string{"tag77"}, IsVisible: false},
		},
		expectedLastId: 7,
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			storage := &ArticleStorage{
				articles: tt.storageArticles,
				lastId:   tt.storageLastId,
			}
			article, err := storage.Create(tt.title, tt.content, tt.category, tt.tags, tt.isVisible)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedArticle, article)
			assert.Equal(t, tt.expectedArticles, storage.articles)
			assert.Equal(t, tt.expectedLastId, storage.lastId)
		})
	}
}

func TestArticleStorage_Update(t *testing.T) {
	tests := []struct {
		name             string
		storageArticles  map[internal.ArticleId]*internal.Article
		storageLastId    internal.ArticleId
		id               int
		title            string
		content          string
		category         internal.Category
		tags             []string
		isVisible        bool
		expectedErr      error
		expectedArticle  *internal.Article
		expectedArticles map[internal.ArticleId]*internal.Article
		expectedLastId   internal.ArticleId
	}{{
		name: "doesn't exist error",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 1, Tags: []string{}, IsVisible: false},
		},
		storageLastId:   3,
		id:              1,
		title:           "title",
		content:         "Content",
		category:        1,
		tags:            []string{"tag1", "tag2"},
		isVisible:       true,
		expectedErr:     errors.New("article doesn't exist"),
		expectedArticle: nil,
		expectedArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 1, Tags: []string{}, IsVisible: false},
		},
		expectedLastId: 3,
	}, {
		name: "ok",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 22, Tags: []string{}, IsVisible: false},
			3: {Id: 3, Title: "title", Content: "Content", Category: 333, Tags: []string{}, IsVisible: false},
			4: {Id: 4, Title: "title", Content: "Content", Category: 4, Tags: []string{}, IsVisible: true},
		},
		storageLastId: 8,
		id:            3,
		title:         "new title3",
		content:       "new content 33",
		category:      32323,
		tags:          []string{"tag1", "tag2"},
		isVisible:     true,
		expectedErr:   nil,
		expectedArticle: &internal.Article{
			Id:        3,
			Title:     "new title3",
			Content:   "new content 33",
			Category:  internal.Category(32323),
			Tags:      []string{"tag1", "tag2"},
			IsVisible: true,
		},
		expectedArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 22, Tags: []string{}, IsVisible: false},
			3: {Id: 3, Title: "new title3", Content: "new content 33", Category: 32323, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title", Content: "Content", Category: 4, Tags: []string{}, IsVisible: true},
		},
		expectedLastId: 8,
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			storage := &ArticleStorage{
				articles: tt.storageArticles,
				lastId:   tt.storageLastId,
			}
			article, err := storage.Update(internal.ArticleId(tt.id), tt.title, tt.content, tt.category, tt.tags, tt.isVisible)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedArticle, article)
			assert.Equal(t, tt.expectedArticles, storage.articles)
			assert.Equal(t, tt.expectedLastId, storage.lastId)
		})
	}
}

func TestArticleStorage_Get(t *testing.T) {
	tests := []struct {
		name            string
		storageArticles map[internal.ArticleId]*internal.Article
		storageLastId   internal.ArticleId
		id              int
		expectedArticle *internal.Article
	}{{
		name: "nil article",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 22},
			3: {Id: 3, Title: "title", Content: "Content", Category: 333},
		},
		storageLastId:   4,
		id:              1,
		expectedArticle: nil,
	}, {
		name: "ok",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 22, Tags: []string{"tag2", "tag22"}, IsVisible: false},
			3: {Id: 3, Title: "title2", Content: "Content2", Category: 333, Tags: []string{}, IsVisible: true},
		},
		storageLastId: 4,
		id:            2,
		expectedArticle: &internal.Article{
			Id:        2,
			Title:     "title",
			Content:   "Content",
			Category:  internal.Category(22),
			Tags:      []string{"tag2", "tag22"},
			IsVisible: false,
		},
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			storage := &ArticleStorage{
				articles: tt.storageArticles,
				lastId:   tt.storageLastId,
			}
			article, err := storage.Get(internal.ArticleId(tt.id))
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedArticle, article)
			assert.Equal(t, tt.storageArticles, storage.articles)
			assert.Equal(t, tt.storageLastId, storage.lastId)
		})
	}
}

func TestArticleStorage_Delete(t *testing.T) {
	tests := []struct {
		name             string
		storageArticles  map[internal.ArticleId]*internal.Article
		storageLastId    internal.ArticleId
		id               int
		expectedArticles map[internal.ArticleId]*internal.Article
		expectedLastId   internal.ArticleId
	}{{
		name: "nonexistent id",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 22},
			3: {Id: 3, Title: "title", Content: "Content", Category: 333},
		},
		storageLastId: 4,
		id:            1,
		expectedArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title", Content: "Content", Category: 22},
			3: {Id: 3, Title: "title", Content: "Content", Category: 333},
		},
		expectedLastId: 4,
	}, {
		name: "existing id",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 22},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 333},
		},
		storageLastId: 4,
		id:            3,
		expectedArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 22},
		},
		expectedLastId: 4,
	}}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			storage := &ArticleStorage{
				articles: tt.storageArticles,
				lastId:   tt.storageLastId,
			}
			err := storage.Delete(internal.ArticleId(tt.id))
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedArticles, storage.articles)
			assert.Equal(t, tt.expectedLastId, storage.lastId)
		})
	}
}

func TestArticleStorage_Filter(t *testing.T) {
	tests := []struct {
		name             string
		storageArticles  map[internal.ArticleId]*internal.Article
		storageLastId    internal.ArticleId
		categories       []internal.Category
		tags             []string
		onlyVisible      bool
		expectedArticles []*internal.Article
	}{{
		name: "all",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}, IsVisible: false},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}, IsVisible: false},
			5: {Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
			6: {Id: 6, Title: "title6", Content: "Content6", Category: 6, Tags: []string{}, IsVisible: false},
		},
		storageLastId: 6,
		categories:    []internal.Category{},
		tags:          []string{},
		onlyVisible:   false,
		expectedArticles: []*internal.Article{
			{Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}, IsVisible: false},
			{Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			{Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
			{Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
			{Id: 6, Title: "title6", Content: "Content6", Category: 6, Tags: []string{}},
		},
	}, {
		name: "category",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
			5: {Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
			6: {Id: 6, Title: "title6", Content: "Content6", Category: 6, Tags: []string{}},
		},
		storageLastId: 6,
		categories:    []internal.Category{4, 5},
		tags:          []string{},
		onlyVisible:   false,
		expectedArticles: []*internal.Article{
			{Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
			{Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
		},
	}, {
		name: "tag",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
			5: {Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
			6: {Id: 6, Title: "title6", Content: "Content6", Category: 6, Tags: []string{}},
		},
		storageLastId: 6,
		categories:    []internal.Category{},
		tags:          []string{"tag1"},
		onlyVisible:   false,
		expectedArticles: []*internal.Article{
			{Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}},
			{Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			{Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
		},
	}, {
		name: "tag&visible",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
			5: {Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
			6: {Id: 6, Title: "title6", Content: "Content6", Category: 6, Tags: []string{}},
		},
		storageLastId: 6,
		categories:    []internal.Category{},
		tags:          []string{"tag2", "tag3"},
		onlyVisible:   true,
		expectedArticles: []*internal.Article{
			{Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			{Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
		},
	}, {
		name: "category&tag",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
			5: {Id: 5, Title: "title5", Content: "Content5", Category: 5, Tags: []string{"tag3"}, IsVisible: true},
			6: {Id: 6, Title: "title6", Content: "Content6", Category: 6, Tags: []string{}},
		},
		storageLastId: 6,
		categories:    []internal.Category{3},
		tags:          []string{"tag1"},
		onlyVisible:   false,
		expectedArticles: []*internal.Article{
			{Id: 3, Title: "title3", Content: "Content3", Category: 3, Tags: []string{"tag1", "tag2"}, IsVisible: true},
		},
	}, {
		name: "category&visible",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 2, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title4", Content: "Content4", Category: 4, Tags: []string{"tag1", "tag2"}},
			5: {Id: 5, Title: "title5", Content: "Content5", Category: 2, Tags: []string{"tag3"}, IsVisible: true},
			6: {Id: 6, Title: "title6", Content: "Content6", Category: 6, Tags: []string{}},
		},
		storageLastId: 6,
		categories:    []internal.Category{2},
		tags:          []string{},
		onlyVisible:   true,
		expectedArticles: []*internal.Article{
			{Id: 3, Title: "title3", Content: "Content3", Category: 2, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			{Id: 5, Title: "title5", Content: "Content5", Category: 2, Tags: []string{"tag3"}, IsVisible: true},
		},
	}, {
		name: "category&tag&visible",
		storageArticles: map[internal.ArticleId]*internal.Article{
			2: {Id: 2, Title: "title2", Content: "Content2", Category: 2, Tags: []string{"tag1", "tag2"}},
			3: {Id: 3, Title: "title3", Content: "Content3", Category: 2, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			4: {Id: 4, Title: "title4", Content: "Content4", Category: 2, Tags: []string{"tag1", "tag2"}},
			5: {Id: 5, Title: "title5", Content: "Content5", Category: 2, Tags: []string{"tag2"}, IsVisible: true},
			6: {Id: 6, Title: "title6", Content: "Content6", Category: 2, Tags: []string{}},
		},
		storageLastId: 6,
		categories:    []internal.Category{2},
		tags:          []string{"tag2"},
		onlyVisible:   true,
		expectedArticles: []*internal.Article{
			{Id: 3, Title: "title3", Content: "Content3", Category: 2, Tags: []string{"tag1", "tag2"}, IsVisible: true},
			{Id: 5, Title: "title5", Content: "Content5", Category: 2, Tags: []string{"tag2"}, IsVisible: true},
		},
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			storage := &ArticleStorage{
				articles: tt.storageArticles,
				lastId:   tt.storageLastId,
			}
			articles, err := storage.Filter(tt.categories, tt.tags, tt.onlyVisible)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedArticles, articles)
			assert.Equal(t, tt.storageArticles, storage.articles)
			assert.Equal(t, tt.storageLastId, storage.lastId)
		})
	}
}
