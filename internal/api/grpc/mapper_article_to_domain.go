package grpc

import (
	"errors"
	"go.citilink.cloud/grpc-skeleton/internal"
	articlev1 "go.citilink.cloud/grpc-skeleton/internal/api/grpc/gen/citilink/blog/article/v1"
)

var domainCategories = map[articlev1.Category]internal.Category{
	articlev1.Category_CATEGORY_PEOPLE:  internal.CategoryPeople,
	articlev1.Category_CATEGORY_ANIMALS: internal.CategoryAnimals,
	articlev1.Category_CATEGORY_TRAVELS: internal.CategoryTravels,
}

// articleToDomainMapper маппер статей из protobuf в доменные структуры
type articleToDomainMapper struct{}

func (a *articleToDomainMapper) mapCategory(pbCategory articlev1.Category) (internal.Category, error) {
	category, ok := domainCategories[pbCategory]
	if !ok {
		return 0, errors.New("category id is invalid")
	}

	return category, nil
}
