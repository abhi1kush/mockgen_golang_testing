package bird

type Bird interface {
	//weight unit is kg, speed unit kilometer per hour
	FlyingSpeed() (speed int)
}

type Swan struct {
	Color  string
	Weight int
}

func (s *Swan) FlyingSpeed() (speed int) {
	return 100 / s.Weight
}
