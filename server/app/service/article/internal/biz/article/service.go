package article

import (
	"context"
	"math"
	v1 "yuumi/api/service/article/v1"
	"yuumi/app/service/article/internal/data/mysql"
	"yuumi/app/service/article/internal/data/mysql/article"
	"yuumi/app/service/article/internal/data/mysql/schema"
	"yuumi/app/service/article/internal/data/mysql/transform"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"
)

type ServiceInterface interface {
	Create(ctx context.Context, in *v1.CreateArticleRequest) (*v1.CreateArticleReply, error)
	Delete(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error)
	Update(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error)
	GetInfo(ctx context.Context, in *v1.GetArticleInfoRequest) (*v1.GetArticleInfoReply, error)
	GetList(ctx context.Context, in *v1.GetArticleListRequest) (*v1.GetArticleListReply, error)
}

type Service struct {
	Logger *logger.Logger
}

func (s Service) Create(ctx context.Context, in *v1.CreateArticleRequest) (*v1.CreateArticleReply, error) {
	res, err := mysql.GetArticle().CreateOne(ctx, &schema.Article{
		Title:       in.Title,
		Description: in.Description,
		Keyword:     in.Keyword,
		Content:     in.Content,
		CoverUrl:    in.CoverUrl,
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCreationFailed, err)
	}

	return &v1.CreateArticleReply{Article: transform.TransformArticleData(res)}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	_, err := mysql.GetArticle().DeleteWithID(ctx, in.Id)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordDeletionFailed, err)
	}
	return &v1.DeleteArticleReply{Message: "ok"}, nil
}

func (s Service) Update(ctx context.Context, in *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error) {
	fields := schema.Article{}.GetFields()
	target := map[string]interface{}{}
	if in.Title != nil {
		target[fields.Title] = in.Title
	}
	if in.Description != nil {
		target[fields.Description] = *in.Description
	}
	if in.CoverUrl != nil {
		target[fields.CoverUrl] = *in.CoverUrl
	}
	if in.Keyword != nil {
		target[fields.Keyword] = *in.Keyword
	}
	if in.Content != nil {
		target[fields.Content] = *in.Content
	}
	if in.Status != nil {
		target[fields.Status] = *in.Status
	}

	res, err := mysql.GetArticle().UpdateWithID(ctx, in.Id, target)
	if err != nil {
		return nil, errorcode.New(errorcode.RecordUpdateFailed)
	}

	return &v1.UpdateArticleReply{Article: transform.TransformArticleData(res)}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetArticleInfoRequest) (*v1.GetArticleInfoReply, error) {
	res, err := mysql.GetArticle().FindOne(ctx, &article.FindCondition{
		Where: &article.FindConditionWhere{
			ID:     *in.Id,
			Status: in.Status,
		},
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}
	return &v1.GetArticleInfoReply{Article: transform.TransformArticleData(res)}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetArticleListRequest) (*v1.GetArticleListReply, error) {
	findCondition := &article.FindCondition{
		Where: &article.FindConditionWhere{
			Keyword: in.Keywrod,
			Status:  in.Status,
		},
		Sort: &article.FindConditionSort{ID: -1},
	}

	total, err := mysql.GetArticle().Count(ctx, findCondition.Where, true)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCountFailed, err)
	}

	result, err := mysql.GetArticle().FindList(ctx, &article.FindListCondition{
		Page:          in.Page,
		PageSize:      in.PageSize,
		FindCondition: findCondition,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetArticleListReply{
		Articles: transform.TransformArticlesData(result),
		Pagination: &v1.PaginationData{
			Page:      in.Page,
			PageSize:  in.PageSize,
			Total:     total,
			PageTotal: int64(math.Ceil(float64(total) / float64(in.PageSize))),
		},
	}, nil
}
