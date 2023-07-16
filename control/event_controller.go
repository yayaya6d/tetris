package control

type EventController interface {
	StartPollingControlEvent(pollingChan chan<- int)
}
