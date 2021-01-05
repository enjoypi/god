package net

//
//type Service struct {
//	Config
//	*zap.Logger
//	*application.Node
//	pb.UnimplementedSessionServer
//
//	childState sc.State
//	godSvc     *application.Service
//}
//
//func NewService(cfg Config, logger *zap.Logger, node *application.Node, initialState sc.State, childState sc.State) *application.Service {
//	svc := &Service{
//		Config:     cfg,
//		Logger:     logger,
//		Node:       node,
//		childState: childState,
//	}
//	svc.godSvc = application.NewService(logger, svc, initialState, svc)
//	core.Go(func(core.ExitChan, interface{}) (interface{}, error) {
//		err := svc.Serve()
//		return nil, err
//	}, nil, nil)
//	return svc.godSvc
//}
//
//func (svc *Service) Serve() error {
//	lis, err := net.Listen("tcp", svc.Config.Net.ListenAddress)
//	if err != nil {
//		return err
//	}
//
//	s := grpc.NewServer()
//	pb.RegisterSessionServer(s, svc)
//
//	svc.Info(lis.Addr().String())
//	return s.Serve(lis)
//}
//
//func (svc *Service) Flow(stream pb.Session_FlowServer) error {
//	actor, err := svc.godSvc.NewAgent(0, svc.childState,
//		&Session{Logger: svc.Logger, Node: svc.Node, Session_FlowServer: stream})
//	if err != nil {
//		return err
//	}
//
//	actor.Run(core.DefaultActors.ExitChan)
//	svc.godSvc.RemoveAgent(actor.ID)
//	return nil
//}
