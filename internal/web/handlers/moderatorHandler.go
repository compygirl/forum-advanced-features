package handlers

import (
	"errors"
	helpers "forum/internal/web/handlers/helpers"
	"net/http"
)

func (h *Handler) ModeratorRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cookie := helpers.SessionCookieGet(r)
		if cookie == nil {
			helpers.ErrorHandler(w, http.StatusUnauthorized, errors.New("Cookie failed in the Moderator Request Handler"))
			return
		}
		expTime, err := h.service.UserServiceInterface.ExtendSessionTimeout(cookie.Value)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, errors.New("Cookie cannot be extended in Approval of the Post"))
			return
		}
		err = helpers.SessionCookieExtend(r, w, expTime)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		session, err := h.service.UserServiceInterface.GetSession(cookie.Value)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, errors.New("Session failed in the Moderator Request Handler"))
			return
		}

		user, err := h.service.UserServiceInterface.GetUserByUserID(session.UserID)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		// changing role to pending from user
		err = h.service.UserServiceInterface.ChangeUserRole("pending", user.UserUserID)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("Error in Moderator Request Handler"))
		return
	}
}
