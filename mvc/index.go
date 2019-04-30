package mvc

import (
	ginmvc "dc-xm/dc/lite/ginex/mvc"
	actionhandler "dc-xm/dc/lite/ginex/mvc/action-handler"
	actionresult "dc-xm/dc/lite/ginex/mvc/action-result"
	routehandler "dc-xm/dc/lite/ginex/mvc/route-handler"
	"dc-xm/dc/lite/mvc"
	invokehandler "dc-xm/dc/lite/mvc/invoke-handler"
	servicehandler "dc-xm/dc/lite/mvc/service-handler"
	validationhandler "dc-xm/dc/lite/validation/mvc/validation-handler"
)

func Run(port int) {
	ginmvc.Listen(
		port,
		createRelease(),
	)
}

func createRelease() mvc.IHandler {
	handleErr := func(err error, ctx mvc.IContext) {
		actionresult.Panic(err).Exec(ctx)
	}
	m := routehandler.New(
		handleErr,
	)
	m.SetNext(
		servicehandler.New(handleErr),
	).SetNext(
		actionhandler.New(handleErr),
	).SetNext(
		validationhandler.New(),
	).SetNext(
		invokehandler.New(handleErr),
	)
	return m
}
