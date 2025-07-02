package router

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/gorilla/mux"
	pubsubrouter "github.com/sofyan48/pubsub-router"

	"github.com/kiriminaja/kaj-rest-engine-go/src/app/appctx"
	"github.com/kiriminaja/kaj-rest-engine-go/src/bootstrap"
	"github.com/kiriminaja/kaj-rest-engine-go/src/consts"
	"github.com/kiriminaja/kaj-rest-engine-go/src/handler"
	"github.com/kiriminaja/kaj-rest-engine-go/src/middleware"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/database"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/validates"

	"github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/contract"
	ucaseContract "github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/contract"
)

type router struct {
	router        *mux.Router
	validator     validates.Validator
	dbAdapter     database.Adapter
	pubsubrouters *pubsubrouter.Router
	// gcpClient     gcp.Contract
	// dbMongoAdapter mongodb.Adapter
	// elasticSession elastic.Client
	// kfConsumer    kafka.Consumer
	// kfProducer    kafka.Producer

}

func NewRouter() Router {
	bootstrap.RegistryMessage()
	bootstrap.RegistryLogger()

	return &router{
		router:        mux.NewRouter(),
		validator:     validates.New(),
		dbAdapter:     bootstrap.RegistryMariaMasterSlave(),
		pubsubrouters: bootstrap.PubSubRouter(),
		// gcpClient:     bootstrap.GCPClient(),
		// kfConsumer:    bootstrap.RegistryKafkaConsumer(),
		// kfProducer:    bootstrap.RegistryKafkaProducer(),
		// dbMongoAdapter: bootstrap.NewRegisterMongoDB(),
	}
}

func (rtr *router) task(ctx context.Context, sleep int64, task ucaseContract.TaskBackground) {
	handler.TaskHandler(sleep, task, ctx)
}

// func (rtr *router) kafkaProcessorHandle(ctx context.Context, topic []string, consumerGroup string, mp ucaseContract.KafkaProcessor) {
// 	handler.KafkaProcessorHandler(ctx, rtr.kfConsumer, rtr.kfProducer, topic, consumerGroup, mp)
// }

func (rtr *router) eventHandle(svc contract.EventProcessor) pubsubrouter.HandlerFunc {
	return svc.Serve
}

func (rtr *router) handle(hfn httpHandlerFunc, svc ucaseContract.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				res := appctx.Response{
					Name: consts.InternalFailure,
				}

				res.SetMessage()
				logger.Error(logger.SetMessageFormat("error %v", string(debug.Stack())))
				err = json.NewEncoder(w).Encode(res)
				if err != nil {
					logger.Error(logger.SetMessageFormat("error %v", string(debug.Stack())))
					return
				}
				return
			}
		}()
		if r.Header.Get("X-Real-IP") != "" {
			r.RemoteAddr = r.Header.Get("X-Real-IP")
		}

		var st time.Time
		var lt time.Duration

		st = time.Now()

		ctx := context.WithValue(r.Context(), "access", map[string]interface{}{
			"path":      r.URL.Path,
			"remote_ip": r.RemoteAddr,
			"method":    r.Method,
		})

		req := r.WithContext(ctx)

		if status := middleware.FilterFunc(req, mdws); status != consts.MiddlewarePassed {
			rtr.response(w, appctx.Response{
				Name: status,
			})

			return
		}

		resp := hfn(req, svc)

		resp.Lang = rtr.defaultLang(req.Header.Get("X-Lang"))

		rtr.response(w, resp)

		lt = time.Since(st)
		logger.AccessLog("request",
			logger.Any("tag", "http-request"),
			logger.Any("http.path", req.URL.Path),
			logger.Any("http.method", req.Method),
			logger.Any("http.agent", req.UserAgent()),
			logger.Any("http.referer", req.Referer()),
			logger.Any("http.status", resp.GetCode()),
			logger.Any("http.latency", lt.Seconds()),
		)
	}
}

// response prints as a json and formatted string for DGP legacy
func (rtr *router) response(w http.ResponseWriter, resp appctx.Response) {
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		resp.SetMessage()
		w.WriteHeader(resp.GetCode())
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			logger.Error(logger.SetMessageFormat("error %v", string(debug.Stack())))
		}
	}()

	return

}

func (rtr *router) defaultLang(l string) string {

	if len(l) == 0 {
		return os.Getenv("APP_DEFAULT_LANG")
	}

	return l
}
