package grpc

import (
	"context"
	"fmt"

	"go.citilink.cloud/citizap"
	"go.citilink.cloud/citizap/cores/swaplevel"
	citizap_factory "go.citilink.cloud/citizap/factory/ctx"
	"go.citilink.cloud/grpc-skeleton/internal"
	articlev1 "go.citilink.cloud/grpc-skeleton/internal/api/grpc/gen/citilink/blog/article/v1"
	storev1 "go.citilink.cloud/grpc-skeleton/internal/specs/grpcclient/gen/citilink/store/store/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ArticleServer сервер статей
type ArticleServer struct {
	storage          internal.ArticleStorage
	storeClient      storev1.StoreAPIClient
	toDomainMapper   *articleToDomainMapper
	toProtobufMapper *articleToProtobufMapper
	loggerFactory    citizap_factory.Factory
	articlev1.UnimplementedArticleAPIServer
}

func NewArticleServer(
	storage internal.ArticleStorage,
	storeClient storev1.StoreAPIClient,
	loggerFactory citizap_factory.Factory,
) *ArticleServer {
	return &ArticleServer{
		storage:          storage,
		storeClient:      storeClient,
		loggerFactory:    loggerFactory,
		toDomainMapper:   &articleToDomainMapper{},
		toProtobufMapper: &articleToProtobufMapper{},
	}
}

func (a *ArticleServer) Create(ctx context.Context, in *articlev1.CreateRequest) (*articlev1.CreateResponse, error) {
	// Логгер создаётся через фабрику
	logger := a.loggerFactory.Create(ctx)
	// Обернуть логгер в декоратор, который будет подменять уровень логирования
	// с Error на Info для ошибок отмены контекста
	logger = swaplevel.WrapCore(logger, zap.InfoLevel, swaplevel.IsContextCanceledErr)

	title := in.GetTitle()
	category, err := a.toDomainMapper.mapCategory(in.GetCategory())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("can't map category id: %s", err))
	}

	article, err := a.storage.Create(title, in.GetContent(), category, in.GetTags(), in.GetIsVisible())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't create article: %s", err))
	}

	pbArticle, err := a.toProtobufMapper.mapArticle(article)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't map article: %s", err))
	}

	// Вымышленный пример обращения к другому мс
	_, err = a.storeClient.GetSpaceIds(ctx, &storev1.GetSpaceIdsRequest{})
	if err != nil {
		// Если контекст был отменён, ошибка будет залогирована с уровнем Info
		logger.Error("failed to get spaces ids from store", citizap.Error(err))
		return nil, status.Error(codes.Internal, "failed to get spaces ids from store")
	}

	return &articlev1.CreateResponse{Article: pbArticle}, nil
}

func (a *ArticleServer) Update(_ context.Context, in *articlev1.UpdateRequest) (*articlev1.UpdateResponse, error) {
	id := in.GetId()
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	article, err := a.storage.Get(internal.ArticleId(id))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("can't get article: %s", err))
	}
	if article == nil {
		return nil, status.Error(codes.NotFound, "article doesn't exist")
	}

	title := article.Title
	if in.GetTitle() != nil {
		title = in.GetTitle().GetValue()
	}

	content := article.Content
	if in.GetContent() != nil {
		content = in.GetContent().GetValue()
	}

	category := article.Category
	if in.GetCategory() != nil {
		category, err = a.toDomainMapper.mapCategory(in.GetCategory().GetValue())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("can't map category id: %s", err))
		}
	}

	tags := article.Tags
	if in.GetTags() != nil {
		tags = in.GetTags().GetValues()
	}

	isVisible := article.IsVisible
	if in.GetIsVisible() != nil {
		isVisible = in.GetIsVisible().GetValue()
	}

	article, err = a.storage.Update(article.Id, title, content, category, tags, isVisible)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't update article: %s", err))
	}

	pbArticle, err := a.toProtobufMapper.mapArticle(article)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't map article: %s", err))
	}

	return &articlev1.UpdateResponse{Article: pbArticle}, nil
}

func (a *ArticleServer) Get(_ context.Context, in *articlev1.GetRequest) (*articlev1.GetResponse, error) {
	id := in.GetId()
	article, err := a.storage.Get(internal.ArticleId(int(id)))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't get article: %s", err))
	}

	pbArticle, err := a.toProtobufMapper.mapArticle(article)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't map article: %s", err))
	}

	return &articlev1.GetResponse{Article: pbArticle}, nil
}

func (a *ArticleServer) Filter(_ context.Context, in *articlev1.FilterRequest) (*articlev1.FilterResponse, error) {
	categories := make([]internal.Category, 0, len(in.GetCategories()))
	for _, pbCategory := range in.GetCategories() {
		category, err := a.toDomainMapper.mapCategory(pbCategory)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("can't map category id: %s", err))
		}

		categories = append(categories, category)
	}

	articles, err := a.storage.Filter(categories, in.GetTags(), in.GetOnlyVisible())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't filter articles: %s", err))
	}

	pbArticles := make([]*articlev1.Article, 0, len(articles))
	for _, article := range articles {
		pbArticle, err := a.toProtobufMapper.mapArticle(article)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't map article: %s", err))
		}
		pbArticles = append(pbArticles, pbArticle)
	}

	return &articlev1.FilterResponse{
		Articles: pbArticles,
	}, nil
}

func (a *ArticleServer) Delete(_ context.Context, in *articlev1.DeleteRequest) (*articlev1.DeleteResponse, error) {
	for _, id := range in.GetIds() {
		err := a.storage.Delete(internal.ArticleId(id))
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't delete article: %s", err))
		}
	}

	return &articlev1.DeleteResponse{}, nil
}
