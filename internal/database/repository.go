package database

import (
	"database/sql"
	"forum/internal/models"
)

type UserRepoInterface interface {
	CreateUserRepo(*models.User) (int64, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserByUsername(string) (*models.User, error)
	GetUserByUserID(int) (*models.User, error)
	CreateSession(*models.Session) error
	UpdateSession(*models.Session) error
	GetSessionByUserID(int) (*models.Session, error)
	GetSessionByToken(string) (*models.Session, error)
	DeleteSessionByToken(string) error
	DeleteSessionByUserID(int) error
	ChangeUserRole(string, int) error
	GetUserRole(int) (string, error)
	GetUserByRole(string) ([]*models.User, error)
}

type PostRepoInterface interface {
	CreatePostRepo(*models.Post) (int64, error)
	GetAllPosts() ([]*models.Post, error)
	GetCategoriesByPostID(int) ([]string, error)
	GetPostByID(int) (*models.Post, error)
	GetPostsByUserId(int) ([]*models.Post, error)
	GetPostsByLikes(int) ([]*models.Post, error)
	CreatePostCategory([]string, int) (int64, error)
	UpdateLikesCounter(int, int) error
	UpdateDislikesCounter(int, int) error
	GetReaction(int, int) (int, error)
	AddReactionToPostVotes(int, int, int) error
	DeleteFromPostVotes(int, int) error
	UpdateReactionInPostVotes(int, int, int) error
	GetPostsByCategory(string) ([]*models.Post, error)
	DeletePostByID(int) error
	DeletePostCategoryByPostID(int) error
	DeleteAllPostVotesByPostID(int) error
	UpdateIsApprovePostStatus(int) error
	ChangeReportStatusOfPostbyPostID(int, int) error
	AddPostReportCategory(int, string) error
	GetAllCategories() ([]*models.Category, error)
	DeletePostCategoryByCategoryID(int) error
	CreateCategory(string) (int64, error)
	UpdatePostContentByPostID(int, string) error
	GetMyReactedPosts(int) (map[int]int, error)
	GetAllMyPostsLikedByOtherUsers(int) ([]*models.PostVotes, error)
	GetAllMyPostsCommentedByOtherUsers(int) ([]*models.PostVotes, error)
}

type CommentRepoInterface interface {
	CreateCommentRepo(*models.Comment) (int64, error)
	GetAlCommentsForPost(int) ([]*models.Comment, error)
	UpdateLikesCounter(int, int) error
	UpdateDislikesCounter(int, int) error
	GetCommentReaction(int, int) int
	AddReactionToCommentVotes(int, int, int) error
	DeleteReactionFromCommentVotes(int, int) error
	UpdateReactionInCommentVotes(int, int, int) error
	DeleteAllCommentsByPostID(int) error
	DeleteAllCommentVotesByPostID(int) error
	DeleteAllCommentVotesByCommentID(int) error
	DeleteCommentByCommentID(int) error
	UpdateIsApproveCommentStatus(int) error
	UpdateCommentContentByPostID(int, string) error
	GetMyReactedComments(int) (map[int]int, error)
	GetCommentByID(int) (*models.Comment, error)
	GetCommentByUserID(int) ([]*models.Comment, error)
}

type Repository struct {
	UserRepoInterface
	PostRepoInterface
	CommentRepoInterface
}

func NewRepository(db *sql.DB) *Repository {
	repositoryObj := Repository{
		UserRepoInterface:    CreateNewUserDB(db),
		PostRepoInterface:    CreateNewPostDB(db),
		CommentRepoInterface: CreateNewCommentDB(db),
	}
	return &repositoryObj
}
