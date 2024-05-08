package grpc

import (
	"errors"
	"fmt"
	"go.citilink.cloud/grpc-skeleton/internal"
	articlev1 "go.citilink.cloud/grpc-skeleton/internal/api/grpc/gen/citilink/blog/article/v1"
)

var pbCategories = map[internal.Category]articlev1.Category{
	internal.CategoryPeople:  articlev1.Category_CATEGORY_PEOPLE,
	internal.CategoryAnimals: articlev1.Category_CATEGORY_ANIMALS,
	internal.CategoryTravels: articlev1.Category_CATEGORY_TRAVELS,
}

// articleToProtobufMapper маппер статей из доменных структур в protobuf
type articleToProtobufMapper struct{}

func (a *articleToProtobufMapper) mapCategory(category internal.Category) (articlev1.Category, error) {
	pbCategory, ok := pbCategories[category]
	if !ok {
		return articlev1.Category_CATEGORY_INVALID, errors.New("category id is invalid")
	}

	return pbCategory, nil
}

func (a *articleToProtobufMapper) mapArticle(article *internal.Article) (*articlev1.Article, error) {
	category, err := a.mapCategory(article.Category)
	if err != nil {
		return nil, fmt.Errorf("can't map category id: %w", err)
	}

	return &articlev1.Article{
		Id:        int32(article.Id),
		Title:     article.Title,
		Content:   article.Content,
		Category:  category,
		Tags:      article.Tags,
		IsVisible: article.IsVisible,
	}, nil
}
