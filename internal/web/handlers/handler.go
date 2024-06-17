package handlers

import (
	service "forum/internal/service"
	"net/http"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	handlerObj := Handler{service: service}
	return &handlerObj
}

func (handler *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	images := http.FileServer(http.Dir("./data/assets/images"))
	mux.Handle("/images/", http.StripPrefix("/images/", images))
	mux.HandleFunc("/", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.GetMainPage)))
	mux.HandleFunc("/registration", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.RegistrationHandler))))
	mux.HandleFunc("/login", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.OnlyUnauthMiddleware(handler.CheckCookieMiddleware(handler.LoginHandler))))
	mux.HandleFunc("/logout", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.LogoutHandler))))
	mux.HandleFunc("/submit-post", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.CreatePostHandler))))
	mux.HandleFunc("/post/react", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ReactOnPostHandler))))
	mux.HandleFunc("/comments/", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.DisplayCommentsHandler)))
	mux.HandleFunc("/submit-comment", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.CreateCommentsHandler))))
	mux.HandleFunc("/comment/react", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ReactOnCommentHandler))))
	mux.HandleFunc("/filter/", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.FilterHandler)))
	// authorisation
	mux.HandleFunc("/auth/google/in", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GoogleAuthHandler))))
	mux.HandleFunc("/auth/google/callback", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GoogleCallback))))
	mux.HandleFunc("/auth/github/in", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GithubAuthHandler))))
	mux.HandleFunc("/auth/github/callback", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.OnlyUnauthMiddleware(handler.GithubCallback))))
	// moderation
	mux.HandleFunc("/moderator", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ModeratorRequestHandler))))
	mux.HandleFunc("/admin_page", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.AdminMainPageHandler))))
	mux.HandleFunc("/approve-reject", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ApproveRejectModeratorHandler))))
	mux.HandleFunc("/moderator_list", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ManageModeratorsHandler))))
	mux.HandleFunc("/delete_moderator", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.DeleteModeratorHandler))))
	mux.HandleFunc("/delete_post", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.DeletePostHandler))))
	mux.HandleFunc("/delete_comment", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.DeleteCommentHandler))))
	mux.HandleFunc("/approve_post", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ApprovePostHandler))))
	mux.HandleFunc("/approve_comment", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ApproveCommentHandler))))
	mux.HandleFunc("/report_post", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ReportPostHandler))))
	mux.HandleFunc("/answer_report", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.AnswerPostReportHandler))))
	mux.HandleFunc("/create_categories", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.AdminDisplayCategoriesHandler))))
	mux.HandleFunc("/delete_category", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.AdminDeleteCategoryHandler))))
	mux.HandleFunc("/add_category", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.AdminAddCategoryHandler))))
	// advanced-features
	mux.HandleFunc("/edit_post", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.EditPostHandler))))
	mux.HandleFunc("/edit_comment", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.EditCommentHandler))))
	mux.HandleFunc("/created_my_posts", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ShowMyPostsHandler))))
	mux.HandleFunc("/reacted_posts", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ShowMyReactedPostsHandler))))
	mux.HandleFunc("/reacted_comments", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ShowMyReactedCommentsHandler))))
	mux.HandleFunc("/commented_posts", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ShowMyCommentsWithPostsHandler))))
	mux.HandleFunc("/notifications", NewRateLimiter(10, time.Minute).LimitMiddleware(handler.CheckCookieMiddleware(handler.NeedAuthMiddleware(handler.ShowMyNotificationsHandler))))
	return mux
}
