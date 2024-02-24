package app

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/2twin/L0/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	app.serviceProvider.natsStreaming.Subscribe(ctx, app.serviceProvider.orderRepository)

	return app, nil
}

func (a *App) Run() error {
	return a.runHttpServer()
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.connectNatsStreaming,
		a.initHttpServer,
	}

	for _, fn := range deps {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) connectNatsStreaming(_ context.Context) error {
	err := a.serviceProvider.NatsStreaming().Connect()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) newRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))
	return router
}

func (a *App) handlerGetOrder(w http.ResponseWriter, r *http.Request) {
	orderUUID := chi.URLParam(r, "orderUUID")

	order, err := a.serviceProvider.OrderRepository().Get(context.Background(), orderUUID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "couldn't get order")
	}

	respondWithJSON(w, http.StatusOK, order)
}

func (a *App) initHttpServer(ctx context.Context) error {
	router := a.newRouter()
	router.Get("/order/{orderUUID}", a.handlerGetOrder)

	server := &http.Server{
		Handler: router,
		Addr:    a.serviceProvider.HttpConfig().Address(),
	}
	a.httpServer = server
	return nil
}

func (a *App) runHttpServer() error {
	log.Printf("Server is running on address %s\n", a.serviceProvider.HttpConfig().Address())
	log.Fatal(a.httpServer.ListenAndServe())
	return nil
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Responding with 5XX error: %v", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, status, errorResponse{
		Error: msg,
	})
}
