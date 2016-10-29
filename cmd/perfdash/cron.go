package main

import "time"

/*

Theory of Operation

We put collected values into buckets of various sizes
and store a certain number of buckets of each size.
We add new values into the newest smallest bucket,
then rotate the oldest few smallest buckets into the newest
next-smallest bucket.

When it is time to rotate, we have a current span layout
and a new span layout. For example, to rotate a span x into
a span y:

                      t0
  ... yyy   yyy   yyy |x x x x x x x x x ...
  ... yyy   yyy   yyy   yyy |x x x x x x ...
                            t1


*/

var sizes = []time.Duration{
	time.Minute, // first size must == bucket size
	time.Hour,   // each size must be an integral multiple of previous
	24 * time.Hour,
}

// len(spans) == len(sizes)
var spans = []int{
	60 * 24, // one day (of minutes)
	24 * 30, // one month (of hours)
	3 * 365, // three years (of days)
}

func cron() {
	for t := range time.Tick(sizes[0]) {
		rotateSchema(t)
		fetchMetrics()
	}
}

func rotateSchema(t time.Time) {
	for i, size := range sizes[1:] {
		rotateSizeIn(t, size, sizes[i])
	}
}

func rotateSizeIn(t time.Time, toSize, fromSize time.Duration) {
	var zero time.Time
	if t.Sub(zero)%toSize != 0 {
		return
	}

	t0:= t
}

func fetchMetrics() {
}
