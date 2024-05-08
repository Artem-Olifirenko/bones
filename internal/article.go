// TODO: это вымышленный пример, удалите в реальном приложении
package internal

//go:generate mockgen -typed -source=article.go -destination=./article_mock_test.go -package=internal
type ArticleId int
type Category int

const (
	CategoryPeople Category = iota + 1
	CategoryAnimals
	CategoryTravels
)

// Article статья
type Article struct {
	// Идентификатор
	Id ArticleId
	// Название
	Title string
	// Контент
	Content string
	// Идентификатор категории
	Category Category
	// теги
	Tags []string
	// видимость
	IsVisible bool
}

// ArticleStorage хранилище статей
type ArticleStorage interface {
	// Create создает статью
	Create(title string, content string, category Category, tags []string, isVisible bool) (*Article, error)
	// Update обновляет статью
	Update(id ArticleId, title string, content string, category Category, tags []string, isVisible bool) (*Article, error)
	// Get получает статью
	Get(id ArticleId) (*Article, error)
	// Filter фильтрует статьи
	Filter(categories []Category, tags []string, onlyVisible bool) ([]*Article, error)
	// Delete удаляет статьи
	Delete(id ArticleId) error
}
