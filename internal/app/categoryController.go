package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"vegetableShop/internal/models"
	"vegetableShop/pkg/responder"
)

func (a *App) addCategoryController() {
	a.router.HandleFunc("/categories", a.showAllCategoryHandler)
	a.router.HandleFunc("/categories/{id:[0-9]+}", a.showCategoryHandler)
	a.router.HandleFunc("/categories/{id:[0-9]+}/edit", a.updateCategoryHandler).Methods("POST")
	a.router.HandleFunc("/categories/{id:[0-9]+}/delete", a.deleteCategoryHandler)
	a.router.HandleFunc("/categories/add", a.createCategoryHandler).Methods("POST")
}

func (a *App) showAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if categories, err := a.storage.Category().GetAll(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, map[string][]models.Category{"categories": categories})
	}
}

func (a *App) showCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if category, err := a.storage.Category().FindById(id); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, category)
	}
}

func (a *App) updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if err := r.ParseForm(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	}
	name := r.PostForm.Get("name")
	description := r.PostForm.Get("description")
	parentId, _ := strconv.Atoi(r.PostForm.Get("parent_id"))

	if err := a.storage.Category().Update(uint(id), name, description, uint(parentId)); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, map[string]bool{"success": true})
	}
}

func (a *App) deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if err := a.storage.Category().Delete(uint(id)); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, map[string]bool{"success": true})
	}
}

func (a *App) createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	description := r.PostFormValue("description")
	parentId, _ := strconv.Atoi(r.PostFormValue("parentId"))
	fmt.Printf("d:%s %s %d", name, description, parentId)
	category := models.Category{
		Name:        name,
		Description: description,
		ParentID:    uint(parentId),
	}
	if category, err := a.storage.Category().Create(category); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		responder.RespondWithJson(w, http.StatusOK, category)
	}
}
