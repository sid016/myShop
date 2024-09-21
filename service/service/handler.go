package service

type handler struct {
	Service *Service
}

func newHandler(s *Service) *handler {
	return &handler{Service: s}
}
