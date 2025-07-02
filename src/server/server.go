package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kiriminaja/kaj-rest-engine-go/src/app/router"
	"github.com/kiriminaja/kaj-rest-engine-go/src/bootstrap"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
	"github.com/rs/cors"
	pubsubrouter "github.com/sofyan48/pubsub-router"
)

// NewHTTPServer creates http server instance
// returns: Server instance
func NewHTTPServer() Server {
	return &httpServer{
		router: router.NewRouter(),
	}
}

// httpServer as HTTP server implementation
type httpServer struct {
	router router.Router
}

// Run runs the http server gracefully
// returns:
//
//	err: error operation
func (h *httpServer) Run(ctx context.Context, port string) error {
	var err error
	h.router.TaskRoutes(ctx)
	h.router.KafkaProcessor(ctx)

	corsSetup := cors.New(cors.Options{
		AllowedOrigins: strings.Split(os.Getenv("CORS_CONTROL_ALLOW_ORIGIN"), ","),
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
			http.MethodPatch,
		},
		AllowedHeaders: []string{
			"*",
		},
		AllowCredentials: true,
	})

	server := http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		Handler:      corsSetup.Handler(h.router.Route(ctx)),
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(15) * time.Second,
	}

	go func() {
		err = server.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Error(logger.SetMessageFormat("http server got %#v", err))
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutDown); err != nil {
		logger.Fatal(logger.SetMessageFormat("server Shutdown Failed:%+s", err))
	}

	logger.Info("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}

// Done runs event wheen service stopped
func (h *httpServer) Done() {
	logger.Info("service http stopped")
}

func (h *httpServer) EventProcessor(ctx context.Context) {
	subscription := strings.Split(os.Getenv("SUBSCRIBER_TOPIC"), ",")
	if len(subscription) > 0 && os.Getenv("SUBSCRIBER_TOPIC") != "" {
		sv := pubsubrouter.NewServer(ctx, bootstrap.PubsubRouterCfg())
		eventRouter := h.router.EventRouter()
		for _, i := range subscription {
			logger.Info(logger.SetMessageFormat("Subscriber Topic Starting %s", i))
			go sv.Subscribe(i, eventRouter).Start()
		}
	}
}
