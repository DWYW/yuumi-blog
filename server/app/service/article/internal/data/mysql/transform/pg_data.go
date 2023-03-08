package transform

import (
	v1 "yuumi/api/service/article/v1"
	"yuumi/app/service/article/internal/data/mysql/schema"
)

func TransformArticleData(v *schema.Article) *v1.ArticleData {
	return &v1.ArticleData{
		Id:          int64(v.ID),
		CreatedAt:   v.CreatedAt.String(),
		UpdatedAt:   v.UpdatedAt.String(),
		Title:       v.Title,
		Description: v.Description,
		Keyword:     v.Keyword,
		Content:     v.Content,
		CoverUrl:    v.CoverUrl,
	}
}

func TransformArticlesData(v []*schema.Article) []*v1.ArticleData {
	var articles = make([]*v1.ArticleData, len(v))
	for i, article := range v {
		articles[i] = TransformArticleData(article)
	}
	return articles
}

func TransformGithubUserData(v *schema.GithubUser) *v1.GithubUserData {
	return &v1.GithubUserData{
		Id:        int64(v.ID),
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
		Name:      v.Name,
		AvatarUrl: v.AvatarUrl,
		Blog:      v.Blog,
		HtmlUrl:   v.HtmlUrl,
	}
}

func TransformVisitorData(v *schema.Visitor) *v1.VisitorData {
	target := v1.VisitorData{
		Id:        int64(v.ID),
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
		Name:      v.Name,
		AvatarUrl: v.AvatarUrl,
	}

	if v.GithubUser.ID != 0 {
		target.GithubUser = TransformGithubUserData(&v.GithubUser)
	}

	return &target
}

func TransformVisitorsData(v []*schema.Visitor) []*v1.VisitorData {
	var visitors = make([]*v1.VisitorData, len(v))
	for i, visitor := range v {
		visitors[i] = TransformVisitorData(visitor)
	}
	return visitors
}

func TransformCommentData(v *schema.Comment) *v1.CommentData {
	target := v1.CommentData{
		Id:        int64(v.ID),
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
		Content:   v.Content,
		CreatorId: v.CreatorID,
		ArticleId: v.ArticleID,
	}

	if v.Creator != nil {
		target.Creator = TransformVisitorData(v.Creator)
	}

	if v.Article != nil {
		target.Article = TransformArticleData(v.Article)
	}

	return &target
}

func TransformCommentsData(v []*schema.Comment) []*v1.CommentData {
	var comments = make([]*v1.CommentData, len(v))
	for i, comment := range v {
		comments[i] = TransformCommentData(comment)
	}
	return comments
}
