package impl

func (h *handler) RegisterApi() {

	h.route.POST("/login", h.Login)
}
