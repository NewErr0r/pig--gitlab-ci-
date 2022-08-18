package internal

type API struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type apiHandler struct {
	api *API
}
