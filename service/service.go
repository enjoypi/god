package service

// 1 service 1 supervisor 0-N actor

type Service struct {
	Type uint16
	*supervisor
}

func NewService(sup *supervisor)  *Service{
	return &Service{
		supervisor: sup,
	}
}