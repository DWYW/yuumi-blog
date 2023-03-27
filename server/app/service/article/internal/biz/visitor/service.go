package visitor

import (
	"context"
	"errors"
	"math"
	v1 "yuumi/api/service/article/v1"
	"yuumi/app/service/article/internal/data/mysql"
	"yuumi/app/service/article/internal/data/mysql/schema"
	"yuumi/app/service/article/internal/data/mysql/transform"
	"yuumi/app/service/article/internal/data/mysql/visitor"
	"yuumi/pkg/errorcode"
	"yuumi/pkg/logger"

	"gorm.io/gorm"
)

type ServiceInterface interface {
	CreateWithGithubUser(ctx context.Context, in *v1.CreateWithGithubUserRequest) (*v1.CreateWithGithubUserReply, error)
	UpdateWithGithubUser(ctx context.Context, in *v1.UpdateWithGithubUserRequest) (*v1.UpdateWithGithubUserReply, error)
	Delete(ctx context.Context, in *v1.DeleteVisitorRequest) (*v1.DeleteVisitorReply, error)
	GetInfo(ctx context.Context, in *v1.GetVisitorInfoRequest) (*v1.GetVisitorInfoReply, error)
	GetList(ctx context.Context, in *v1.GetVisitorListRequest) (*v1.GetVisitorListReply, error)
	GetVisitorWithGithubID(ctx context.Context, in *v1.GetVisitorWithGithubIDRequest) (*v1.GetVisitorWithGithubIDReply, error)
}

type Service struct {
	Logger *logger.Logger
}

func (s Service) CreateWithGithubUser(ctx context.Context, in *v1.CreateWithGithubUserRequest) (*v1.CreateWithGithubUserReply, error) {
	res, err := mysql.GetVisitor().CreateOne(ctx, &schema.Visitor{
		Name:      in.GithubUser.Name,
		AvatarUrl: in.GithubUser.AvatarUrl,
		GithubUser: schema.GithubUser{
			GithubID:  uint64(in.GithubUser.GithubId),
			Name:      in.GithubUser.Name,
			AvatarUrl: in.GithubUser.AvatarUrl,
			Blog:      in.GithubUser.Blog,
			HtmlUrl:   in.GithubUser.HtmlUrl,
		},
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCreationFailed, err)
	}

	return &v1.CreateWithGithubUserReply{Visitor: transform.TransformVisitorData(res)}, nil
}

func (s Service) UpdateWithGithubUser(ctx context.Context, in *v1.UpdateWithGithubUserRequest) (*v1.UpdateWithGithubUserReply, error) {
	visitorFields := schema.Visitor{}.GetFields()
	userFields := schema.GithubUser{}.GetFields()
	res, err := mysql.GetVisitor().UpdateWithIDAndUpdateGithubUser(ctx, in.Id, map[string]interface{}{
		visitorFields.Name:      in.GithubUser.Name,
		visitorFields.AvatarUrl: in.GithubUser.AvatarUrl,
	}, map[string]interface{}{
		userFields.Name:      in.GithubUser.Name,
		userFields.AvatarUrl: in.GithubUser.AvatarUrl,
		userFields.Blog:      in.GithubUser.Blog,
		userFields.HtmlUrl:   in.GithubUser.HtmlUrl,
	})

	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCreationFailed, err)
	}

	return &v1.UpdateWithGithubUserReply{Visitor: transform.TransformVisitorData(res)}, nil
}

func (s Service) Delete(ctx context.Context, in *v1.DeleteVisitorRequest) (*v1.DeleteVisitorReply, error) {
	_, err := mysql.GetVisitor().DeleteWithID(ctx, in.Id)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordDeletionFailed, err)
	}
	return &v1.DeleteVisitorReply{Message: "ok"}, nil
}

func (s Service) GetInfo(ctx context.Context, in *v1.GetVisitorInfoRequest) (*v1.GetVisitorInfoReply, error) {
	res, err := mysql.GetVisitor().FindByID(ctx, *in.Id)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordNotFound, err)
	}
	return &v1.GetVisitorInfoReply{Visitor: transform.TransformVisitorData(res)}, nil
}

func (s Service) GetList(ctx context.Context, in *v1.GetVisitorListRequest) (*v1.GetVisitorListReply, error) {
	findCondition := &visitor.FindCondition{
		Where: &visitor.FindConditionWhere{},
		Sort:  &visitor.FindConditionSort{ID: -1},
	}
	if in.Keywrod != nil {
		findCondition.Where.Keyword = *in.Keywrod
	}

	total, err := mysql.GetVisitor().Count(ctx, findCondition.Where, true)
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordCountFailed, err)
	}

	result, err := mysql.GetVisitor().FindList(ctx, &visitor.FindListCondition{
		Page:          in.Page,
		PageSize:      in.PageSize,
		FindCondition: findCondition,
	})
	if err != nil {
		return nil, errorcode.NewWithDetail(errorcode.RecordsNotFound, err)
	}

	return &v1.GetVisitorListReply{
		Visitors: transform.TransformVisitorsData(result),
		Pagination: &v1.PaginationData{
			Page:      in.Page,
			PageSize:  in.PageSize,
			Total:     total,
			PageTotal: int64(math.Ceil(float64(total) / float64(in.PageSize))),
		},
	}, nil
}

func (s Service) GetVisitorWithGithubID(ctx context.Context, in *v1.GetVisitorWithGithubIDRequest) (*v1.GetVisitorWithGithubIDReply, error) {
	res, err := mysql.GetVisitor().FindOneWithGithubID(ctx, in.GithubId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &v1.GetVisitorWithGithubIDReply{}, nil
		}

		return nil, err
	}

	return &v1.GetVisitorWithGithubIDReply{Visitor: transform.TransformVisitorData(res)}, nil
}
