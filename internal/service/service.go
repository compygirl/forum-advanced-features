package service

import (
	"forum/internal/database"
	"forum/internal/models"
	"mime/multipart"
	"net/http"
	"time"
)

type UserServiceInterface interface {
	CreateUser(*models.User) (int, int, error)
	Login(string, string, bool) (*models.Session, error)
	IsUserLoggedIn(*http.Request) bool
	Logout(string) error
	IsTokenExist(string) bool
	GetUserByUserID(int) (*models.User, error)
	GetSession(string) (*models.Session, error)
	ExtendSessionTimeout(string) (time.Time, error)
	GoogleAuthorization(*models.GoogleLoginUserData) (*models.Session, error)
	GitHubAuthorization(*models.GitHubLoginUserData) (*models.Session, error)
	ChangeUserRole(string, int) error
	GetUsersByRole(string) ([]*models.User, error)
}

type PostServiceInterface interface {
	GetAllPosts() ([]*models.Post, error)
	GetPostByID(int) (*models.Post, error)
	CreatePost(*models.Post, string) (int, int, error)
	GetCategories(int) ([]string, error)
	GetPostsByUserId(int) ([]*models.Post, error)
	UpdateReaction(int, int, int) error
	Filter(string, int) ([]*models.Post, error)
	AddImagesToPost(*multipart.FileHeader) (string, error)
	DeletePost(int) error
	DeletePostCategoryByPostID(int) error
	DeleteAllPostVotesByPostID(int) error
	ApprovePost(int) error
	ChangeReportStatusOfPostbyPostID(int, int) error
	AddPostReportCategory(int, string) error
	GetAllCategories() ([]*models.Category, error)
	DeletePostCategory(int) error
	CreateCategory(string) (int, int, error)
	UpdatePostContentByPostID(int, string) error
	GetMyReactedPosts(int) (map[int]int, error)
	GetAllMyPostsLikedByOtherUsers(int) ([]*models.PostVotes, error)
	GetAllMyPostsCommentedByOtherUsers(int) ([]*models.PostVotes, error)
}

type CommentServiceInterface interface {
	CreateComment(*models.Comment, string) (int, int, error)
	GetAlCommentsForPost(int) ([]*models.Comment, error)
	UpdateReaction(int, int, int) error
	DeleteAllCommentsByPostID(int) error
	DeleteAllCommentVotesByPostID(int) error
	DeleteAllCommentVotesByCommentID(int) error
	DeleteCommentByCommentID(int) error
	ApproveComment(int) error
	UpdateCommentContentByPostID(int, string) error
	GetMyReactedComments(int) (map[int]int, error)
	GetCommentByID(int) (*models.Comment, error)
	GetCommentByUserID(int) ([]*models.Comment, error)
}

type Service struct {
	UserServiceInterface // interface
	PostServiceInterface
	CommentServiceInterface
}

func NewService(repo *database.Repository) *Service {
	serviceObj := Service{
		UserServiceInterface:    CreateNewUserService(repo.UserRepoInterface),
		PostServiceInterface:    CreateNewPostService(repo.PostRepoInterface),
		CommentServiceInterface: CreateNewCommentService(repo.CommentRepoInterface),
	}
	return &serviceObj
}
