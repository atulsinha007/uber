package driver

//
//type Handler struct {
//	ctrl Ctrl
//}
//
//func NewHandler(ctrl Ctrl) *Handler {
//	return &Handler{ctrl: ctrl}
//}
//
//func (h *Handler) CreateDriver(req *http.Request) handler.Response {
//	var payload Create
//	err := json.NewDecoder(req.Body).Decode(&payload)
//	if err != nil {
//		return handler.BadRequest("invalid payload")
//	}
//}
