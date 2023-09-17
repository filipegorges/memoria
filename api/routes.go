package api

func (a *API) BuildRoutes() {
	v1 := a.srv.Group("/v1")
	a.Handlers(v1)
}
