package database

import (
	"database/sql"
	"forum/internal/models"
)

type CommentRepoImpl struct {
	db *sql.DB
}

func CreateNewCommentDB(db *sql.DB) *CommentRepoImpl {
	return &CommentRepoImpl{db}
}

func (cmnt *CommentRepoImpl) CreateCommentRepo(comment *models.Comment) (int64, error) {
	result, err := cmnt.db.Exec(`
	INSERT INTO comments (user_id, post_id, content, created_time, likes_counter, dislikes_counter, is_approved, reports) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`,
		comment.UserID, comment.PostID, comment.Content, comment.CreatedTime, comment.LikesCounter, comment.DislikeCounter, comment.IsApproved, comment.ReportStatus)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (cmnt *CommentRepoImpl) GetAlCommentsForPost(postID int) ([]*models.Comment, error) {
	comments := []*models.Comment{}

	rows, err := cmnt.db.Query("SELECT * FROM comments ORDER BY created_time DESC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedTime, &comment.LikesCounter, &comment.DislikeCounter, &comment.IsApproved, &comment.ReportStatus)
		if err != nil {
			return nil, err
		}

		if postID == comment.PostID {
			comments = append(comments, &comment)
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (cmnt *CommentRepoImpl) UpdateLikesCounter(commentID, valueToAdd int) error {
	_, err := cmnt.db.Exec("UPDATE comments SET likes_counter = likes_counter + ? WHERE id = ?", valueToAdd, commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) UpdateDislikesCounter(commentID, valueToAdd int) error {
	_, err := cmnt.db.Exec("UPDATE comments SET dislikes_counter = dislikes_counter + ? WHERE id = ?", valueToAdd, commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) GetCommentReaction(commentID, userID int) int {
	var reaction int
	if err := cmnt.db.QueryRow(
		`SELECT reaction FROM comment_votes WHERE comment_id = ? AND user_id = ?`,
		commentID, userID).Scan(&reaction); err != nil {
		return 0
	}
	return reaction
}

func (cmnt *CommentRepoImpl) AddReactionToCommentVotes(commentID, userID, reaction int) error {
	_, err := cmnt.db.Exec(`
		INSERT INTO comment_votes (comment_id, user_id,reaction) VALUES (?, ?, ?);`,
		commentID, userID, reaction)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) DeleteReactionFromCommentVotes(commentID, userID int) error {
	_, err := cmnt.db.Exec("DELETE FROM comment_votes WHERE comment_id = ? AND user_id = ?", commentID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) UpdateReactionInCommentVotes(commentID, userID, newReaction int) error {
	_, err := cmnt.db.Exec("UPDATE comment_votes SET reaction = ? WHERE comment_id = ? AND user_id = ?", newReaction, commentID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) DeleteAllCommentsByPostID(postID int) error {
	_, err := cmnt.db.Exec("DELETE FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) DeleteAllCommentVotesByPostID(postID int) error {
	_, err := cmnt.db.Exec("DELETE FROM comment_votes WHERE comment_id IN (SELECT id FROM comments WHERE post_id = ?)", postID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) DeleteAllCommentVotesByCommentID(commentID int) error {
	_, err := cmnt.db.Exec("DELETE FROM comment_votes WHERE comment_id = ?", commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) DeleteCommentByCommentID(commentID int) error {
	_, err := cmnt.db.Exec("DELETE FROM comments WHERE id = ?;", commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) UpdateIsApproveCommentStatus(commentID int) error {
	_, err := cmnt.db.Exec("UPDATE comments SET is_approved = 1 WHERE id = ?", commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) UpdateCommentContentByPostID(intCommentID int, content string) error {
	_, err := cmnt.db.Exec("UPDATE comments SET content = ? WHERE id = ?", content, intCommentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmnt *CommentRepoImpl) GetMyReactedComments(userID int) (map[int]int, error) {
	commentToReaction := make(map[int]int)
	rows, err := cmnt.db.Query(`
	SELECT comment_id,reaction FROM comment_votes WHERE user_id=?`, userID)
	if err != nil {
		// fmt.Println("FILTER:  1 error")
		return nil, err
	}

	for rows.Next() {
		var commentID int
		var reaction int
		err = rows.Scan(&commentID, &reaction)
		if err != nil {
			return nil, err
		}
		commentToReaction[commentID] = reaction
	}
	if err = rows.Err(); err != nil {
		// fmt.Println("FILTER:  2 error")
		return nil, err
	}

	return commentToReaction, nil
}

func (cmnt *CommentRepoImpl) GetCommentByID(commentID int) (*models.Comment, error) {
	comment := &models.Comment{}

	if err := cmnt.db.QueryRow(
		`SELECT * FROM comments WHERE id = ?`,
		commentID).Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedTime, &comment.LikesCounter, &comment.DislikeCounter, &comment.IsApproved, &comment.ReportStatus); err != nil {
		return nil, err
	}
	return comment, nil
}

func(cmnt *CommentRepoImpl) GetCommentByUserID(userID int) ([]*models.Comment, error) {
	comments := []*models.Comment{}
	rows, err := cmnt.db.Query(`
	SELECT * FROM comments WHERE user_id = ? ORDER BY created_time DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedTime, &comment.LikesCounter, &comment.DislikeCounter, &comment.IsApproved, &comment.ReportStatus)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
