// 该方案存在两个问题：
// 1.存储了它见过的所有Event
// 2.MinuteCount和HourCount的时间复杂度都是O(n)

package simple_solve

import "time"

type Event struct {
	Count int
	Ts    int64 //时间戳s
}

// MinuteHourCounter Track the cumulative counts over the past minute and over the past hour.
// Useful, for example, to track recent bandwidth usage.
type MinuteHourCounter struct {
	Events []Event
}

func (c *MinuteHourCounter) Add(count int) {
	c.Events = append(c.Events, Event{
		Count: count,
		Ts:    time.Now().Unix(),
	})
}

func (c *MinuteHourCounter) countSince(cutoffTs int64) int {
	if len(c.Events) == 0 {
		return 0
	}
	var count int
	for i := len(c.Events) - 1; i >= 0; i-- {
		e := c.Events[i]
		if e.Ts <= cutoffTs {
			break
		}
		count += e.Count
	}
	return count
}

// MinuteCount Return the accumulated count over the past 60 seconds
func (c *MinuteHourCounter) MinuteCount() int {
	return c.countSince(time.Now().Add(-60 * time.Second).Unix())
}

// HourCount Return the accumulated count over the past 1 hour
func (c *MinuteHourCounter) HourCount() int {
	return c.countSince(time.Now().Add(-time.Hour).Unix())
}
