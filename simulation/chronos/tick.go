package chrono

type TickFunc func(at Tick)

type Tick int

type Ticker interface {
	WaitNext() Tick
	ScheduleNext(f TickFunc)
	ScheduleFor(t Tick, f TickFunc)
}
