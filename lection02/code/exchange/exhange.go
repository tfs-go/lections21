package exchange

import (
	"errors"
	"time"
)

type Exchange struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time

	duration time.Duration
}

func New(name string, sTime, eTime time.Time) (Exchange, error) {
	if eTime.Before(sTime) {
		return Exchange{}, errors.New("end time can't be before start time")
	}
	return Exchange{
		Name:      name,
		StartTime: sTime,
		EndTime:   eTime,
		duration:  eTime.Sub(sTime),
	}, nil
}

func (e Exchange) Duration() time.Duration {
	return e.duration
}

func Duration(e Exchange) time.Duration {
	return e.duration
}

// так не надо :(
func (e Exchange) GetDuration() time.Duration {
	return e.duration
}

func (e *Exchange) UpdateEndTime(t time.Time) {
	e.EndTime = t
	e.duration = e.EndTime.Sub(e.StartTime)
}

//nolint: staticcheck
func (e Exchange) UpdateEndTimeValue(t time.Time) {
	e.EndTime = t
	e.duration = e.EndTime.Sub(e.StartTime)
}
