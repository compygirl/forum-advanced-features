package handlers

import (
	"errors"
	"forum/internal/models"
	"forum/internal/web/handlers/helpers"
	"net/http"
	"strconv"
)

func (h *Handler) AdminMainPageHandler(w http.ResponseWriter, r *http.Request) {
	adminPagePath := "internal/web/templates/adminPage.html"
	type templateData struct {
		AllRequests []*models.User
	}

	switch r.Method {
	case "GET":
		pendingUsers, err := h.service.UserServiceInterface.GetUsersByRole("pending")
		if err != nil {
			helpers.ErrorHandler(w, http.StatusBadRequest, errors.New("PEDNING USERS were not found"))
		}
		data := templateData{
			AllRequests: pendingUsers,
		}
		helpers.RenderTemplate(w, adminPagePath, data)
		return
	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("in Admin Page Handler"))
		return
	}
}

func (h *Handler) ApproveRejectModeratorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cookie := helpers.SessionCookieGet(r)
		if cookie == nil {
			helpers.ErrorHandler(w, http.StatusUnauthorized, errors.New("Cookie cannot be reseived in Admin Side for Approval Rejecting the Moderator Request"))
			return
		}

		expTime, err := h.service.UserServiceInterface.ExtendSessionTimeout(cookie.Value)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, errors.New("Cookie cannot be extended"))
			return
		}
		err = helpers.SessionCookieExtend(r, w, expTime)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		userID := r.FormValue("userId")
		typeOfButton := r.FormValue("action")

		// fmt.Println("1 USER ID: ", userID, "ACTUION: ", typeOfButton)

		intUserID, err := strconv.Atoi(userID)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}
		// changing role to pending from user
		if typeOfButton == "approve" {
			err = h.service.ChangeUserRole("moderator", intUserID)
			if err != nil {
				helpers.ErrorHandler(w, http.StatusInternalServerError, err)
				return
			}
		} else {
			// fmt.Println("WHEN REJECTED")
			err = h.service.ChangeUserRole("user", intUserID)
			if err != nil {
				helpers.ErrorHandler(w, http.StatusInternalServerError, err)
				return
			}
		}
		// fmt.Println("2 USER ID: ", userID, "ACTUION: ", typeOfButton)

		http.Redirect(w, r, "/admin_page", http.StatusSeeOther)
		return
		// r.ParseForm()
		// fmt.Println("Raw Form Data:", r.Form)

	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("in Admin Page Handler"))
		return
	}
}

func (h *Handler) ManageModeratorsHandler(w http.ResponseWriter, r *http.Request) {
	adminModeratorListPath := "internal/web/templates/moderatorListPage.html"
	type templateData struct {
		AllModerators []*models.User
	}

	switch r.Method {
	case "GET":
		moderatorUsers, err := h.service.UserServiceInterface.GetUsersByRole("moderator")
		if err != nil {
			helpers.ErrorHandler(w, http.StatusBadRequest, errors.New("PEDNING USERS were not found"))
		}
		data := templateData{
			AllModerators: moderatorUsers,
		}
		helpers.RenderTemplate(w, adminModeratorListPath, data)
		return
	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("in Admin Page Handler"))
		return
	}
}

func (h *Handler) DeleteModeratorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cookie := helpers.SessionCookieGet(r)
		if cookie == nil {
			helpers.ErrorHandler(w, http.StatusUnauthorized, errors.New("Cookie cannot be reseived in Admin Side for Approval Rejecting the Moderator Request"))
			return
		}

		expTime, err := h.service.UserServiceInterface.ExtendSessionTimeout(cookie.Value)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, errors.New("Cookie cannot be extended"))
			return
		}
		err = helpers.SessionCookieExtend(r, w, expTime)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}
		userID := r.FormValue("userId")
		// fmt.Println("1 USER ID: ", userID)
		intUserID, err := strconv.Atoi(userID)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}
		err = h.service.UserServiceInterface.ChangeUserRole("user", intUserID)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/moderator_list", http.StatusSeeOther)
		return
	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("in Admin Page Handler"))
		return
	}
}

func (h *Handler) AdminDisplayCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	adminCategoriesListPath := "internal/web/templates/categoriesList.html"

	type templateData struct {
		AllCategories []*models.Category
	}

	switch r.Method {
	case "GET":
		categories, err := h.service.PostServiceInterface.GetAllCategories()
		// fmt.Println(categories)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusBadRequest, errors.New("PEDNING USERS were not found"))
		}
		data := templateData{
			AllCategories: categories,
		}

		helpers.RenderTemplate(w, adminCategoriesListPath, data)
		return
	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("in Admin Page Handler"))
		return
	}
}

func (h *Handler) AdminDeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cookie := helpers.SessionCookieGet(r)
		if cookie == nil {
			helpers.ErrorHandler(w, http.StatusUnauthorized, errors.New("Cookie cannot be reseived in Admin Side for Approval Rejecting the Moderator Request"))
			return
		}

		expTime, err := h.service.UserServiceInterface.ExtendSessionTimeout(cookie.Value)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, errors.New("Cookie cannot be extended"))
			return
		}
		err = helpers.SessionCookieExtend(r, w, expTime)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		categoryID := r.FormValue("CategoryId")
		intCategoryID, err := strconv.Atoi(categoryID)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		err = h.service.PostServiceInterface.DeletePostCategory(intCategoryID)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/create_categories", http.StatusSeeOther)
		return
	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("in Admin Page Handler"))
		return
	}
}

func (h *Handler) AdminAddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cookie := helpers.SessionCookieGet(r)
		if cookie == nil {
			helpers.ErrorHandler(w, http.StatusUnauthorized, errors.New("Cookie cannot be reseived in Admin Side for Approval Rejecting the Moderator Request"))
			return
		}

		expTime, err := h.service.UserServiceInterface.ExtendSessionTimeout(cookie.Value)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, errors.New("Cookie cannot be extended"))
			return
		}
		err = helpers.SessionCookieExtend(r, w, expTime)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		categoryName := r.FormValue("category_name")
		// fmt.Println(categoryName)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}

		_, _, err = h.service.PostServiceInterface.CreateCategory(categoryName)
		if err != nil {
			helpers.ErrorHandler(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/create_categories", http.StatusSeeOther)
		return
	default:
		helpers.ErrorHandler(w, http.StatusMethodNotAllowed, errors.New("in Admin Page Handler"))
		return
	}
}
