package def

type OptionListen struct {
	ActorType
	ActorID

	Network string
	Address string
	Handler ActorType
}
