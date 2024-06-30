// 传送带式的解决方案
// 我们在这里维护两个list：
// list1、存储最近一分钟的事件，如果超出一分钟，进入下面的list
// list2、如果超过最近1h，直接丢弃

package main

import "time"

type Event struct {
	Count int
	Ts    int64 // 时间戳s
}

type MinuteHourCounter struct {
	MinuteEvents []Event
	HourEvents   []Event // only contains elements Not in minute events

	minuteCount, hourCount int
}

func (c *MinuteHourCounter) Add(count int) {
	now := time.Now()
	c.shiftOldEvents(now)

	c.MinuteEvents = append(c.MinuteEvents, Event{
		Count: count,
		Ts:    now.Unix(),
	})
	c.minuteCount += count
	c.hourCount += count
}

func (c *MinuteHourCounter) shiftOldEvents(now time.Time) {
	minuteAgo := now.Add(-time.Minute).Unix()
	hourAgo := now.Add(-time.Hour).Unix()

	for len(c.MinuteEvents) != 0 && c.MinuteEvents[0].Ts < minuteAgo {
		c.HourEvents = append(c.HourEvents, c.MinuteEvents[0])
		c.minuteCount -= c.MinuteEvents[0].Count
		if len(c.MinuteEvents) > 1 {
			c.MinuteEvents = c.MinuteEvents[1:]
		} else {
			c.MinuteEvents = nil
		}
	}

	for len(c.HourEvents) != 0 && c.HourEvents[0].Ts < hourAgo {
		c.minuteCount -= c.MinuteEvents[0].Count
		if len(c.HourEvents) > 1 {
			c.HourEvents = c.HourEvents[1:]
		} else {
			c.HourEvents = nil
		}
	}
}

// MinuteCount Return the accumulated count over the past 60 seconds
func (c *MinuteHourCounter) MinuteCount() int {
	c.shiftOldEvents(time.Now())
	return c.minuteCount
}

// HourCount Return the accumulated count over the past 1 hour
func (c *MinuteHourCounter) HourCount() int {
	c.shiftOldEvents(time.Now())
	return c.hourCount
}
