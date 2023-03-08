package article

import (
	"context"
	v1 "yuumi/api/interface/article/v1"
	articleServiceV1 "yuumi/api/service/article/v1"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateArticleRequest) (*v1.CreateArticleReply, error)
	Delete(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error)
	Update(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error)
	GetInfo(ctx context.Context, in *v1.GetArticleInfoRequest) (*v1.GetArticleInfoReply, error)
	GetList(ctx context.Context, in *v1.GetArticleListRequest) (*v1.GetArticleListReply, error)
}

type Service struct {
	Log           *logger.Logger
	ArticleClient *client.GrpcClient
}

func (s Service) Create(ctx context.Context, in *v1.CreateArticleRequest) (*v1.CreateArticleReply, error) {
	client, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *articleServiceV1.CreateArticleReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := articleServiceV1.NewArticleServiceClient(cc)
		reply, err = client.Create(ctx, &articleServiceV1.CreateArticleRequest{
			Title:       in.Title,
			Description: in.Description,
			CoverUrl:    in.CoverUrl,
			Content:     in.Content,
			Keyword:     in.Keyword,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.CreateArticleReply{Article: reply.Article}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	client, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *articleServiceV1.DeleteArticleReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := articleServiceV1.NewArticleServiceClient(cc)
		reply, err = client.Delete(ctx, &articleServiceV1.DeleteArticleRequest{Id: in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteArticleReply{Message: reply.Message}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error) {
	client, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *articleServiceV1.UpdateArticleReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := articleServiceV1.NewArticleServiceClient(cc)
		reply, err = client.Update(ctx, &articleServiceV1.UpdateArticleRequest{
			Id:          in.Id,
			Title:       in.Title,
			Description: in.Description,
			CoverUrl:    in.CoverUrl,
			Content:     in.Content,
			Keyword:     in.Keyword,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.UpdateArticleReply{Article: reply.Article}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetArticleInfoRequest) (*v1.GetArticleInfoReply, error) {
	client, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *articleServiceV1.GetArticleInfoReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		c := articleServiceV1.NewArticleServiceClient(cc)
		reply, err = c.GetInfo(ctx, &articleServiceV1.GetArticleInfoRequest{Id: &in.Id})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetArticleInfoReply{Article: reply.Article}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetArticleListRequest) (*v1.GetArticleListReply, error) {
	client, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *articleServiceV1.GetArticleListReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		c := articleServiceV1.NewArticleServiceClient(cc)
		reply, err = c.GetList(ctx, &articleServiceV1.GetArticleListRequest{
			Page:     in.Page,
			PageSize: in.PageSize,
			Keywrod:  in.Keyword,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetArticleListReply{Articles: reply.Articles, Pagination: reply.Pagination}, nil
}
