package router

import (
	"github.com/kiriminaja/kaj-rest-engine-go/src/app/ucase/pubsub"
	pubsubrouter "github.com/sofyan48/pubsub-router"
)

func (rtr *router) EventRouter() *pubsubrouter.Router {
	// please dont delete this
	defaultRoutes := pubsub.NewDefaultRoutes()
	rtr.pubsubrouters.Handle("", rtr.eventHandle(defaultRoutes))
	// end dont delete

	exSubs := pubsub.NewSubscriber()
	rtr.pubsubrouters.Handle("/test", rtr.eventHandle(exSubs))
	return rtr.pubsubrouters
}
