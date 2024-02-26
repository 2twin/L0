package app

import (
	"context"
	"log"
	"net/http"
	"time"
	"html/template"

	"github.com/2twin/L0/internal/model"
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
	go func() {
		for {
			app.serviceProvider.natsStreaming.Publish(model.GenerageOrder())
			time.Sleep(3 * time.Second)
		}
	}()

	app.serviceProvider.natsStreaming.Subscribe(ctx, app.serviceProvider.OrderRepository())

	return app, nil
}

func (a *App) Run() error { 
	return a.runHttpServer()
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(context.Context) error{
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

func (a *App) handlerGetOrder(w http.ResponseWriter, r *http.Request) {
	orderUUID := r.URL.Query().Get("order_uid")
	
	var order *model.Order
	var err error

	if orderUUID != "" {
		order, err = a.serviceProvider.OrderRepository().Get(context.Background(), orderUUID)
	}

	if err != nil {
		log.Println(err)
	}

	log.Printf(`
		==================================
		Get Order:
		%v
		==================================
	`, order)

	tmpl_list := []string{"templates/index.html", "templates/order.tmpl"}

    tmpl, err := template.ParseFiles(tmpl_list...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, order); err != nil {
        log.Println(err)
    }
}

func (a *App) initHttpServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/order", a.handlerGetOrder)

	server := &http.Server{
		Handler: mux,
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