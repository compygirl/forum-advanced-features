package service

import (
	"errors"
	"forum/internal/database"
	"forum/internal/models"
	"net/http"
	"time"
)

type CommentServiceImpl struct {
	repo database.CommentRepoInterface
}

func CreateNewCommentService(repo database.CommentRepoInterface) *CommentServiceImpl {
	commentService := CommentServiceImpl{repo: repo}
	return &commentService
}

func (cmtObj *CommentServiceImpl) CreateComment(comment *models.Comment, userRole string) (int, int, error) {
	if err := cmtObj.isCommentParamsValid(comment); err != nil {
		return http.StatusBadRequest, -1, err
	}
	comment.CreatedTime = time.Now()
	comment.LikesCounter = 0
	comment.DislikeCounter = 0
	comment.ReportStatus = 0

	if userRole == "admin" || userRole == "moderator" {
		comment.IsApproved = 1
	} else {
		comment.IsApproved = 0
	}

	// comment.IsApproved = 0

	id, err := cmtObj.repo.CreateCommentRepo(comment)
	if err != nil {
		return http.StatusInternalServerError, -1, err
	}
	comment.CommentID = int(id)
	return http.StatusOK, int(id), nil
}

func (cmtObj *CommentServiceImpl) isCommentParamsValid(comment *models.Comment) error {
	if len(comment.Content) < 2 {
		return errors.New("The content must be at least 2 characters")
	}
	return nil
}

func (cmtObj *CommentServiceImpl) GetAlCommentsForPost(postID int) ([]*models.Comment, error) {
	comments, err := cmtObj.repo.GetAlCommentsForPost(postID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (cmtObj *CommentServiceImpl) UpdateReaction(currReaction, commentID, userID int) error {
	var err error
	prevReaction := cmtObj.repo.GetCommentReaction(commentID, userID)

	if prevReaction == 0 {
		cmtObj.repo.AddReactionToCommentVotes(commentID, userID, currReaction)
		if currReaction == 1 {
			err = cmtObj.repo.UpdateLikesCounter(commentID, 1)
		} else {
			err = cmtObj.repo.UpdateDislikesCounter(commentID, 1)
		}
	} else if prevReaction == currReaction {

		cmtObj.repo.DeleteReactionFromCommentVotes(commentID, userID) // second like/dislike will cancel the reaction
		if currReaction == 1 {
			cmtObj.repo.UpdateLikesCounter(commentID, -1)
		} else {
			cmtObj.repo.UpdateDislikesCounter(commentID, -1)
		}
	} else if prevReaction != currReaction {
		cmtObj.repo.UpdateReactionInCommentVotes(commentID, userID, currReaction)
		if currReaction == 1 {
			cmtObj.repo.UpdateLikesCounter(commentID, 1)
			cmtObj.repo.UpdateDislikesCounter(commentID, -1)
		} else {
			cmtObj.repo.UpdateLikesCounter(commentID, -1)
			cmtObj.repo.UpdateDislikesCounter(commentID, 1)
		}
	}
	return err
}

func (cmtObj *CommentServiceImpl) DeleteAllCommentsByPostID(postID int) error {
	err := cmtObj.repo.DeleteAllCommentsByPostID(postID)
	if err != nil {
		return err
	}
	return nil
}

func (cmtObj *CommentServiceImpl) DeleteAllCommentVotesByPostID(postID int) error {
	err := cmtObj.repo.DeleteAllCommentVotesByPostID(postID)
	if err != nil {
		return err
	}
	return nil
}

func (cmtObj *CommentServiceImpl) DeleteAllCommentVotesByCommentID(commentID int) error {
	err := cmtObj.repo.DeleteAllCommentVotesByCommentID(commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmtObj *CommentServiceImpl) DeleteCommentByCommentID(commentID int) error {
	err := cmtObj.repo.DeleteCommentByCommentID(commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmtObj *CommentServiceImpl) ApproveComment(commentID int) error {
	err := cmtObj.repo.UpdateIsApproveCommentStatus(commentID)
	if err != nil {
		return err
	}
	return nil
}

func (cmtObj *CommentServiceImpl) UpdateCommentContentByPostID(intCommentID int, content string) error {
	err := cmtObj.repo.UpdateCommentContentByPostID(intCommentID, content)
	if err != nil {
		return err
	}
	return nil
}

func (cmtObj *CommentServiceImpl) GetMyReactedComments(userID int) (map[int]int,error) {
	mapa, err := cmtObj.repo.GetMyReactedComments(userID)
	if err != nil {
		return nil,err
	}
	return mapa,nil
}

func (cmtObj *CommentServiceImpl) GetCommentByID(commentID int) (*models.Comment, error) {
	comment, err := cmtObj.repo.GetCommentByID(commentID)
	if err != nil {
		return nil,err
	}
	return comment,nil
}

func(cmtObj *CommentServiceImpl) GetCommentByUserID(userID int) ([]*models.Comment, error) {
	comments, err := cmtObj.repo.GetCommentByUserID(userID)
	if err != nil {
		return nil,err
	}
	return comments,nil
}
