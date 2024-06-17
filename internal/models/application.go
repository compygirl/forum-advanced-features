package models

import (
	"database/sql"
	"time"
)

type GitHubResponseToken struct {
	AccessToken string `json:"access_token"`
	TokenID     string `json:"id_token"`
	Scope       string `json:"scope"`
}

type GoogleResponseToken struct {
	AccessToken string `json:"access_token"`
	TokenID     string `json:"id_token"`
}

type GoogleUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
	Password       string
}

type GitHubUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
	Password       string
}

type GoogleLoginUserData struct {
	ID         int
	Name       string
	Email      string
	Password   string
	FirstName  string
	SecondName string
	Provider   string
}

type GitHubLoginUserData struct {
	ID         int
	Name       string
	Email      string
	Password   string
	FirstName  string
	SecondName string
	Login      string
	Provider   string
}

type User struct {
	UserUserID int
	FirstName  string
	SecondName string
	Username   string
	Email      string
	Password   string
	Role       string
}

type Session struct {
	UserID  int
	Token   string
	ExpTime time.Time
}

type Post struct {
	PostID            int
	UserID            int
	Username          string
	Title             string
	Content           string
	CreatedTime       time.Time
	CreatedTimeString string
	LikesCounter      int
	DislikeCounter    int
	Categories        []string
	Comments          []*Comment
	ImagePath         string
	UserRole          string
	IsApproved        int
	ReportStatus      int
	ReportCategories  string
	Reaction          string
}

type Comment struct {
	CommentID         int
	PostID            int
	UserID            int
	Username          string
	Content           string
	CreatedTime       time.Time
	CreatedTimeString string
	LikesCounter      int
	DislikeCounter    int
	UserRole          string
	IsApproved        int
	ReportStatus      int
	Reaction          string
}

type Database struct {
	DB *sql.DB
}

type Category struct {
	CategoryID int
	Category   string
}

type CommentsWithPosts struct {
	PostID            int
	PostTitle         string
	PostContent       string
	PostTimeString    string
	CommentID         int
	CommentContent    string
	CommentTimeString string
}

type PostVotes struct {
	PostVotesID int
	PostID      int
	UserID      int
	Reaction    int
	ReactionStr string
}
