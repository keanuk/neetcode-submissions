import (
	"cmp"
	"slices"
)

/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i - 1].end > intervals[i].start {
			return false
		}
	}
	return true
}
