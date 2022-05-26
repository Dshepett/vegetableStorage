package app

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
	"vegetableShop/internal/config"
	"vegetableShop/internal/storage"
	"vegetableShop/pkg/responder"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	router  *mux.Router
	storage *storage.Storage
	config  *config.Config
	logger  *logrus.Logger
}

func Init(config *config.Config) *App {
	a := &App{
		router: mux.NewRouter(),
		config: config,
		logger: logrus.New(),
	}
	a.logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})
	a.storage = storage.New(a.config)
	a.addRouters()
	a.logger.Info("server initialized")
	return a
}

func (a *App) Run() {
	a.storage.Open()
	s := &http.Server{
		Handler:      a.router,
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	a.logger.Info("starting server...")
	a.logger.Fatal(s.ListenAndServe())
}

func (a *App) addRouters() {
	a.addMiddlewares()
	a.addUserController()
	a.addCategoryController()
	a.router.HandleFunc("/login/{id:[0-9]+}", a.LoginHandler)
}

func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	}
	password := r.PostForm.Get("password")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if user, err := a.storage.User().FindById(uint(id)); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, err.Error())
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
			responder.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			responder.RespondWithJson(w, http.StatusOK, map[string]bool{"success": true})
		}
	}
}
