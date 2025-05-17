package app

import (
	"go-unit-of-work-example/internal/config"
	"go-unit-of-work-example/internal/db"
	"go-unit-of-work-example/internal/handler"
	"go-unit-of-work-example/internal/producer"
	"go-unit-of-work-example/internal/repository"
	"go-unit-of-work-example/internal/repository/user"
	"go-unit-of-work-example/internal/service"
	"net/http"
)

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

type App struct {
	cfg    *config.Config
	server *http.ServeMux
}

func (a *App) Build() error {
	conn, err := db.NewConnection(a.cfg.DSN)
	if err != nil {
		return err
	}

	userRepo := user.NewRepo(conn)
	regRepo := repository.NewRegisterUOW(conn)
	queue := producer.New()

	s := service.New(userRepo, regRepo, queue)
	h := handler.New(s)

	a.server = http.NewServeMux()
	a.server.HandleFunc("/register", h.Register)

	return nil
}

func (a *App) Run() error {
	if err := http.ListenAndServe(a.cfg.Addr, a.server); err != nil {
		return err
	}

	return nil
}
