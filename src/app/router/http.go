package router

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/app/ucase"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/handler"
	_ "github.com/Arfiandimas/kaj-rest-engine-go/src/swagger"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Route preparing http router and will return mux router object
func (rtr *router) Route(ctx context.Context) *mux.Router {
	root := rtr.router.PathPrefix("/").Subrouter()
	in := root.PathPrefix("/in/").Subrouter()
	// //internal version sub router
	//inV1 := in.PathPrefix("/v1/").Subrouter()

	// swagger routes
	in.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// healthy
	healthy := ucase.NewHealthCheck()
	in.HandleFunc("/health", rtr.handle(
		handler.HttpRequest,
		healthy,
	)).Methods(http.MethodGet)

	// // // external subrouter
	// ex := root.PathPrefix("/ex/").Subrouter()
	// // please dont changes this prefix and /api dont standard for any other model
	// exV1 := ex.PathPrefix("/v1/").Subrouter()

	// // // this is use boostrap for example purpose, please delete

	// // // this is use repo for example purpose, please delete
	// repoExample := repositories.NewExample(rtr.dbAdapter, rtr.dbMongoAdapter)

	// // // this is use case for example purpose, please delete
	// el := example.NewExampleList(repoExample)
	// ec := example.NewExampleCreate(repoExample)
	// eld := example.NewExampleDetail(repoExample)
	// ed := example.NewExampleDelete(repoExample)
	// // producerKafka := kafkaCase.NewPublish(rtr.kfProducer)

	// // // TODO: create your route here
	// // // this route for example rest, please delete
	// // example list
	// exV1.HandleFunc("/example", rtr.handle(
	// 	handler.HttpRequest,
	// 	el,
	// )).Methods(http.MethodGet)

	// exV1.HandleFunc("/example", rtr.handle(
	// 	handler.HttpRequest,
	// 	ec,
	// )).Methods(http.MethodPost)

	// exV1.HandleFunc("/example/{id:[0-9]+}", rtr.handle(
	// 	handler.HttpRequest,
	// 	eld,
	// )).Methods(http.MethodGet)

	// exV1.HandleFunc("/example/{id:[0-9]+}", rtr.handle(
	// 	handler.HttpRequest,
	// 	ed,
	// )).Methods(http.MethodDelete)

	// // uncoment this if you need pubsub client
	// // gcpPubsub := bootstrap.PubSub(ctx, rtr.gcpClient)
	// // exV1.HandleFunc("/example/publisher", rtr.handle(
	// // 	handler.HttpRequest,
	// // 	pubsub.NewPublisher(gcpPubsub),
	// // )).Methods(http.MethodPost)

	return rtr.router

}
