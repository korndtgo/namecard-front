package util

import (
	"time"
)

// func InTimeSpan(start, end, check time.Time) bool {
// 	_end := end
// 	_check := check
// 	if end.Before(start) {
// 		_end = end.Add(24 * time.Hour)
// 		if check.Before(start) {
// 			_check = check.Add(24 * time.Hour)
// 		}
// 	}
// 	return _check.After(start) && _check.Before(_end)
// }

func InTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
