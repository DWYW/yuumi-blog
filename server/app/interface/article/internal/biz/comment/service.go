package comment

import (
	"context"
	v1 "yuumi/api/interface/article/v1"
	articleServiceV1 "yuumi/api/service/article/v1"
	"yuumi/pkg/client"
	"yuumi/pkg/logger"

	"google.golang.org/grpc"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentReply, error)
	GetList(ctx context.Context, in *v1.GetCommentListRequest) (*v1.GetCommentListReply, error)
}

type Service struct {
	Log           *logger.Logger
	ArticleClient *client.GrpcClient
}

func (s Service) Create(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentReply, error) {
	client, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *articleServiceV1.CreateCommentReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		client := articleServiceV1.NewCommentServiceClient(cc)
		reply, err = client.Create(ctx, &articleServiceV1.CreateCommentRequest{
			Content:   in.Content,
			CreatorId: in.CreatorId,
			ArticleId: in.ArticleId,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.CreateCommentReply{Comment: reply.Comment}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetCommentListRequest) (*v1.GetCommentListReply, error) {
	client, err := s.ArticleClient.NewClient()
	if err != nil {
		return nil, err
	}

	var reply *articleServiceV1.GetCommentListReply
	err = client.Handle(func(ctx context.Context, cc grpc.ClientConnInterface) error {
		c := articleServiceV1.NewCommentServiceClient(cc)
		reply, err = c.GetList(ctx, &articleServiceV1.GetCommentListRequest{
			Page:           in.Page,
			PageSize:       in.PageSize,
			ArticleId:      in.ArticleId,
			CreatorId:      in.CreatorId,
			PreloadCreator: in.PreloadCreator,
			PreloadArticle: in.PreloadArticle,
		})
		return err
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &v1.GetCommentListReply{Comments: reply.Comments, Pagination: reply.Pagination}, nil
}
