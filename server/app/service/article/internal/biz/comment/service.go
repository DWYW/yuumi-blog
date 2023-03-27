package comment

import (
	"context"
	"math"
	v1 "yuumi/api/service/article/v1"
	"yuumi/app/service/article/internal/data/mysql"
	"yuumi/app/service/article/internal/data/mysql/comment"
	"yuumi/app/service/article/internal/data/mysql/schema"
	"yuumi/app/service/article/internal/data/mysql/transform"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"

	"gorm.io/gorm"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentReply, error)
	Delete(ctx context.Context, in *v1.DeleteCommentRequest) (*v1.DeleteCommentReply, error)
	GetInfo(ctx context.Context, in *v1.GetCommentInfoRequest) (*v1.GetCommentInfoReply, error)
	GetList(ctx context.Context, in *v1.GetCommentListRequest) (*v1.GetCommentListReply, error)
}

type Service struct {
	Logger *logger.Logger
}

func (s Service) Create(ctx context.Context, in *v1.CreateCommentRequest) (*v1.CreateCommentReply, error) {
	res, err := mysql.GetComment().CreateOne(ctx, &schema.Comment{
		Content: in.Content,
		Creator: &schema.Visitor{Model: gorm.Model{ID: uint(in.CreatorId)}},
		Article: &schema.Article{Model: gorm.Model{ID: uint(in.ArticleId)}},
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCreationFailed, err)
	}

	return &v1.CreateCommentReply{Comment: transform.TransformCommentData(res)}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteCommentRequest) (*v1.DeleteCommentReply, error) {
	_, err := mysql.GetComment().DeleteWithID(ctx, in.Id)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordDeletionFailed, err)
	}
	return &v1.DeleteCommentReply{Message: "ok"}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetCommentInfoRequest) (*v1.GetCommentInfoReply, error) {
	res, err := mysql.GetComment().FindByID(ctx, *in.Id)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}
	return &v1.GetCommentInfoReply{Comment: transform.TransformCommentData(res)}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetCommentListRequest) (*v1.GetCommentListReply, error) {
	findCondition := &comment.FindCondition{
		Where: &comment.FindConditionWhere{
			ArticleID: in.ArticleId,
			CreatorID: in.CreatorId,
		},
		Sort: &comment.FindConditionSort{ID: -1},
		Preload: &comment.FindConditionPreload{
			Creator: in.PreloadCreator,
			Article: in.PreloadArticle,
		},
	}

	total, err := mysql.GetComment().Count(ctx, findCondition.Where, true)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCountFailed, err)
	}

	result, err := mysql.GetComment().FindList(ctx, &comment.FindListCondition{
		Page:          in.Page,
		PageSize:      in.PageSize,
		FindCondition: findCondition,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetCommentListReply{
		Comments: transform.TransformCommentsData(result),
		Pagination: &v1.PaginationData{
			Page:      in.Page,
			PageSize:  in.PageSize,
			Total:     total,
			PageTotal: int64(math.Ceil(float64(total) / float64(in.PageSize))),
		},
	}, nil
}
