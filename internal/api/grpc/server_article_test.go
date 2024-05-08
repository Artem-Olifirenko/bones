package grpc

//go:generate mockgen -typed -destination=mock/article_storage.go -package=mock go.citilink.cloud/grpc-skeleton/internal ArticleStorage
//go:generate mockgen -typed -destination=mock/store_client.go -package=mock go.citilink.cloud/grpc-skeleton/internal/specs/grpcclient/gen/citilink/store/store/v1 StoreAPIClient

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	citizap_factory_ctx "go.citilink.cloud/citizap/factory/ctx"
	"go.citilink.cloud/grpc-skeleton/internal"
	articlev12 "go.citilink.cloud/grpc-skeleton/internal/api/grpc/gen/citilink/blog/article/v1"
	"go.citilink.cloud/grpc-skeleton/internal/api/grpc/mock"
	storev1 "go.citilink.cloud/grpc-skeleton/internal/specs/grpcclient/gen/citilink/store/store/v1"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestArticleServer_Create_InvalidCategoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Create(context.Background(), &articlev12.CreateRequest{
		Title:    "title",
		Category: articlev12.Category_CATEGORY_INVALID,
		Tags:     []string{"tag1", "tag2"},
	})

	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.InvalidArgument, "can't map category id: category id is invalid"), err)
}

func TestArticleServer_Create_StorageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Create(
			gomock.Eq("title"),
			gomock.Eq(""),
			gomock.Eq(internal.CategoryPeople),
			gomock.Eq([]string{"tag2"}),
			gomock.Eq(false)).
		Times(1).
		Return(nil, errors.New("artificial storage error"))

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Create(context.Background(), &articlev12.CreateRequest{
		Title:    "title",
		Category: 1,
		Tags:     []string{"tag2"},
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't create article: artificial storage error"), err)
}

func TestArticleServer_Create_MapArticleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Create(
			gomock.Eq("title"),
			gomock.Eq("content"),
			gomock.Eq(internal.CategoryAnimals),
			gomock.Eq([]string{"tag1", "tag11", "tag2_1"}),
			gomock.Eq(true)).
		Times(1).
		Return(&internal.Article{
			Id:        1,
			Title:     "title",
			Content:   "content",
			Category:  0,
			Tags:      []string{"tag1", "tag11", "tag2_1"},
			IsVisible: true,
		}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Create(context.Background(), &articlev12.CreateRequest{
		Title:     "title",
		Content:   "content",
		Category:  2,
		Tags:      []string{"tag1", "tag11", "tag2_1"},
		IsVisible: true,
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't map article: can't map category id: category id is invalid"), err)
}

func TestArticleServer_Create_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Create(
			gomock.Eq("title"),
			gomock.Eq("content"),
			gomock.Eq(internal.CategoryAnimals),
			gomock.Eq([]string{"tag1", "tag11", "tag2_1"}),
			gomock.Eq(true)).
		Times(1).
		Return(&internal.Article{
			Id:        5,
			Title:     "title",
			Content:   "content",
			Category:  internal.CategoryAnimals,
			Tags:      []string{"tag1", "tag11", "tag2_1"},
			IsVisible: true,
		}, nil)

	storeAPIClientMock := mock.NewMockStoreAPIClient(ctrl)
	storeAPIClientMock.EXPECT().
		GetSpaceIds(gomock.Eq(ctx), gomock.Eq(&storev1.GetSpaceIdsRequest{})).
		Return(&storev1.GetSpaceIdsResponse{}, nil)

	server := NewArticleServer(storage, storeAPIClientMock, citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Create(ctx, &articlev12.CreateRequest{
		Title:     "title",
		Content:   "content",
		Category:  2,
		Tags:      []string{"tag1", "tag11", "tag2_1"},
		IsVisible: true,
	})
	assert.Equal(t, &articlev12.CreateResponse{Article: &articlev12.Article{
		Id:        5,
		Title:     "title",
		Content:   "content",
		Category:  2,
		Tags:      []string{"tag1", "tag11", "tag2_1"},
		IsVisible: true,
	}}, resp)
	assert.NoError(t, err)
}

func TestArticleServer_Update_IdIsRequiredError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)
	storage.EXPECT().
		Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Update(context.Background(), &articlev12.UpdateRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.InvalidArgument, "id is required"), err)
}

func TestArticleServer_Update_StorageGetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(1))).
		Times(1).
		Return(nil, errors.New("artificial storage get error"))
	storage.EXPECT().
		Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Update(context.Background(), &articlev12.UpdateRequest{Id: 1})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.InvalidArgument, "can't get article: artificial storage get error"), err)
}

func TestArticleServer_Update_StorageNilArticleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(3))).
		Times(1).
		Return(nil, nil)
	storage.EXPECT().
		Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Update(context.Background(), &articlev12.UpdateRequest{Id: 3})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.NotFound, "article doesn't exist"), err)
}

func TestArticleServer_Update_InvalidCategoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(1))).
		Times(1).
		Return(&internal.Article{Id: 1}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Update(context.Background(), &articlev12.UpdateRequest{
		Id: 1,
		Category: &articlev12.UpdateRequest_CategoryValue{
			Value: articlev12.Category_CATEGORY_INVALID,
		},
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.InvalidArgument, "can't map category id: category id is invalid"), err)
}

func TestArticleServer_Update_StorageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(1))).
		Times(1).
		Return(&internal.Article{Id: 1}, nil)
	storage.EXPECT().
		Update(
			gomock.Eq(internal.ArticleId(1)),
			gomock.Eq(""),
			gomock.Eq(""),
			gomock.Eq(internal.CategoryAnimals),
			gomock.Eq([]string{"tag1", "tag2"}),
			gomock.Eq(false)).
		Times(1).
		Return(nil, errors.New("artificial storage update error"))

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Update(context.Background(), &articlev12.UpdateRequest{
		Id:       1,
		Category: &articlev12.UpdateRequest_CategoryValue{Value: 2},
		Tags:     &articlev12.UpdateRequest_TagsValues{Values: []string{"tag1", "tag2"}},
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't update article: artificial storage update error"), err)
}

func TestArticleServer_Update_MapArticleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(44))).
		Times(1).
		Return(&internal.Article{
			Id:        44,
			Title:     "title was",
			Content:   "was content",
			Category:  3,
			Tags:      []string{"tag1", "tag11", "tag2_1"},
			IsVisible: true,
		}, nil)
	storage.EXPECT().
		Update(
			gomock.Eq(internal.ArticleId(44)),
			gomock.Eq("new title"),
			gomock.Eq("content upd"),
			gomock.Eq(internal.CategoryTravels),
			gomock.Eq([]string{"tag1", "tag11", "tag2_1"}),
			gomock.Eq(true)).
		Times(1).
		Return(&internal.Article{
			Id:        44,
			Title:     "content upd title",
			Content:   "content",
			Category:  0,
			Tags:      []string{"tag1", "tag11", "tag2_1"},
			IsVisible: true,
		}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Update(context.Background(), &articlev12.UpdateRequest{
		Id:      44,
		Title:   &articlev12.UpdateRequest_TitleValue{Value: "new title"},
		Content: &articlev12.UpdateRequest_ContentValue{Value: "content upd"},
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't map article: can't map category id: category id is invalid"), err)
}

func TestArticleServer_Update_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(2))).
		Times(1).
		Return(&internal.Article{
			Id:        2,
			Title:     "title",
			Content:   "content",
			Category:  1,
			Tags:      []string{"tag1", "tag2"},
			IsVisible: true,
		}, nil)
	storage.EXPECT().
		Update(
			gomock.Eq(internal.ArticleId(2)),
			gomock.Eq("title"),
			gomock.Eq("content"),
			gomock.Eq(internal.CategoryPeople),
			gomock.Eq([]string{"tag22"}),
			gomock.Eq(false)).
		Times(1).
		Return(&internal.Article{
			Id:        2,
			Title:     "title",
			Content:   "content",
			Category:  1,
			Tags:      []string{"tag22"},
			IsVisible: false,
		}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Update(context.Background(), &articlev12.UpdateRequest{
		Id: 2,
		Category: &articlev12.UpdateRequest_CategoryValue{
			Value: 1,
		},
		Tags: &articlev12.UpdateRequest_TagsValues{
			Values: []string{"tag22"},
		},
		IsVisible: &wrapperspb.BoolValue{
			Value: false,
		},
	})
	assert.Equal(t, &articlev12.UpdateResponse{Article: &articlev12.Article{
		Id:        2,
		Title:     "title",
		Content:   "content",
		Category:  1,
		Tags:      []string{"tag22"},
		IsVisible: false,
	}}, resp)
	assert.NoError(t, err)
}

func TestArticleServer_Get_StorageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(4))).
		Times(1).
		Return(nil, errors.New("artificial storage error"))

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Get(context.Background(), &articlev12.GetRequest{Id: 4})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't get article: artificial storage error"), err)
}

func TestArticleServer_Get_MapArticleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(7))).
		Times(1).
		Return(&internal.Article{
			Id:        7,
			Title:     "title",
			Content:   "content",
			Category:  9,
			Tags:      []string{"tag1"},
			IsVisible: false,
		}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Get(context.Background(), &articlev12.GetRequest{Id: 7})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't map article: can't map category id: category id is invalid"), err)
}

func TestArticleServer_Get_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Get(gomock.Eq(internal.ArticleId(20))).
		Times(1).
		Return(&internal.Article{
			Id:        20,
			Title:     "title",
			Content:   "content",
			Category:  internal.CategoryTravels,
			Tags:      []string{"tag1", "tag11", "tag2_1", "tag4"},
			IsVisible: false,
		}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Get(context.Background(), &articlev12.GetRequest{Id: 20})
	assert.Equal(t, &articlev12.GetResponse{Article: &articlev12.Article{
		Id:        20,
		Title:     "title",
		Content:   "content",
		Category:  3,
		Tags:      []string{"tag1", "tag11", "tag2_1", "tag4"},
		IsVisible: false,
	}}, resp)
	assert.NoError(t, err)
}

func TestArticleServer_Filter_InvalidCategoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().Filter(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Filter(context.Background(), &articlev12.FilterRequest{
		Categories: []articlev12.Category{0, 1, 2, 3},
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.InvalidArgument, "can't map category id: category id is invalid"), err)
}

func TestArticleServer_Filter_StorageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Filter(gomock.Eq([]internal.Category{2, 3}), gomock.Eq([]string{"tag"}), gomock.Eq(false)).
		Times(1).
		Return(nil, errors.New("artificial storage error"))

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Filter(context.Background(), &articlev12.FilterRequest{
		Categories:  []articlev12.Category{2, 3},
		Tags:        []string{"tag"},
		OnlyVisible: false,
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't filter articles: artificial storage error"), err)
}

func TestArticleServer_Filter_MapArticleError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Filter(gomock.Eq([]internal.Category{2, 3}), gomock.Eq([]string{"tag1", "tag2"}), gomock.Eq(false)).
		Times(1).
		Return([]*internal.Article{{
			Id:        9,
			Title:     "title1",
			Content:   "content2",
			Category:  2,
			Tags:      []string{"tag1", "tag2"},
			IsVisible: false,
		}, {
			Id:        11,
			Title:     "title3",
			Content:   "content4",
			Category:  5,
			Tags:      []string{"tag2"},
			IsVisible: true,
		}}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Filter(context.Background(), &articlev12.FilterRequest{
		Categories:  []articlev12.Category{2, 3},
		Tags:        []string{"tag1", "tag2"},
		OnlyVisible: false,
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't map article: can't map category id: category id is invalid"), err)
}

func TestArticleServer_Filter_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Filter(gomock.Eq([]internal.Category{2, 3}), gomock.Eq([]string{"tag1", "tag2"}), gomock.Eq(false)).
		Times(1).
		Return([]*internal.Article{{
			Id:        9,
			Title:     "title1",
			Content:   "content2",
			Category:  2,
			Tags:      []string{"tag1", "tag2"},
			IsVisible: false,
		}, {
			Id:        11,
			Title:     "title3",
			Content:   "content4",
			Category:  3,
			Tags:      []string{"tag2"},
			IsVisible: true,
		}}, nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Filter(context.Background(), &articlev12.FilterRequest{
		Categories:  []articlev12.Category{2, 3},
		Tags:        []string{"tag1", "tag2"},
		OnlyVisible: false,
	})
	assert.Equal(t, &articlev12.FilterResponse{Articles: []*articlev12.Article{{
		Id:        9,
		Title:     "title1",
		Content:   "content2",
		Category:  2,
		Tags:      []string{"tag1", "tag2"},
		IsVisible: false,
	}, {
		Id:        11,
		Title:     "title3",
		Content:   "content4",
		Category:  3,
		Tags:      []string{"tag2"},
		IsVisible: true,
	}}}, resp)
	assert.NoError(t, err)
}

func TestArticleServer_Delete_StorageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Delete(gomock.Eq(internal.ArticleId(2))).
		Times(1).
		Return(nil)
	storage.EXPECT().
		Delete(gomock.Eq(internal.ArticleId(3))).
		Times(1).
		Return(errors.New("artificial storage error"))
	storage.EXPECT().
		Delete(gomock.Eq(internal.ArticleId(5))).
		Times(0)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Delete(context.Background(), &articlev12.DeleteRequest{
		Ids: []int32{2, 3, 5},
	})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Internal, "can't delete article: artificial storage error"), err)
}

func TestArticleServer_Delete_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockArticleStorage(ctrl)
	storage.EXPECT().
		Delete(gomock.Eq(internal.ArticleId(7))).
		Times(1).
		Return(nil)
	storage.EXPECT().
		Delete(gomock.Eq(internal.ArticleId(8))).
		Times(1).
		Return(nil)
	storage.EXPECT().
		Delete(gomock.Eq(internal.ArticleId(11))).
		Times(1).
		Return(nil)

	server := NewArticleServer(storage, mock.NewMockStoreAPIClient(ctrl), citizap_factory_ctx.New(zap.NewNop()))
	resp, err := server.Delete(context.Background(), &articlev12.DeleteRequest{
		Ids: []int32{7, 8, 11},
	})
	assert.Equal(t, &articlev12.DeleteResponse{}, resp)
	assert.NoError(t, err)
}
