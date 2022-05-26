package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"vegetableShop/internal/models"
	"vegetableShop/pkg/responder"
)

func (a *App) addUserController() {
	a.router.HandleFunc("/users", a.showAllUserHandler)
	a.router.HandleFunc("/users/{id:[0-9]+}", a.showUserHandler)
	a.router.HandleFunc("/users/{id:[0-9]+}/edit", a.updateUserHandler).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}/delete", a.deleteUserHandler)
	a.router.HandleFunc("/users/add", a.createUserHandler).Methods("POST")
}

func (a *App) showAllUserHandler(w http.ResponseWriter, r *http.Request) {
	if users, err := a.storage.User().GetAll(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, map[string][]models.User{"users": users})
	}
}

func (a *App) showUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if user, err := a.storage.User().FindById(uint(id)); err == nil {
		responder.RespondWithJson(w, http.StatusOK, user)
	} else {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	}
}

func (a *App) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if err := r.ParseForm(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	}
	username := r.PostForm.Get("name")
	password := r.PostForm.Get("password")
	email := r.PostForm.Get("email")
	userRole := r.PostForm.Get("role")
	role, _ := strconv.Atoi(userRole)
	user := models.User{
		Id:             uint(id),
		Name:           username,
		Email:          email,
		HashedPassword: password,
		Role:           role,
	}

	if err := user.CheckEmail(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := user.EncryptPassword(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := a.storage.User().Update(user); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, map[string]bool{"success": true})
	}
}

func (a *App) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if err := a.storage.User().Delete(uint(id)); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, map[string]bool{"success": true})
	}
}

func (a *App) createUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	userRole := r.PostForm.Get("role")
	role, _ := strconv.Atoi(userRole)
	user := models.User{
		Name:           username,
		HashedPassword: password,
		Email:          email,
		Role:           role,
	}
	if err := user.CheckEmail(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := user.EncryptPassword(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if user, err := a.storage.User().Create(user); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, user)
	}
}
