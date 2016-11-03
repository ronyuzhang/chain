package main

import "time"

/*

Theory of Operation

Perfdash stores numeric values (counters, gauges, and histograms).
measured from physical systems.
The main challenge is how to organize and present these values
so they're useful.

When asking questions of perfdash, the user might want to divide
the results many ways: Which sensor was sampled? Which city was it
in? Which process is being observed? Which action was the user taking?
Which user was taking the action? Which experimental cohort did the
user belong to? When did it happen? These questions (and others like
them) force the data to be organized along many dimensions.

Therefore, perfdash explicitly records for each sample its position
in an arbitrary number of user-defined dimensions.
The various dimensions
may or may not be orthogonal: for example, a "datacenter" dimension
might be determined entirely by a "server" dimension, or
a "procid" dimension might be independent of a "userid" dimension.
Perfdash doesn't care whether dimensions are orthogonal,
or even whether any exist, with two exceptions: the "metric" name,
and time. These two dimensions are required, and they are special.
The "metric" dimension is never orthogonal with respect to any
dimension except time, it contains them all.
Time is always orthogonal to all other dimensions.
(This orthogonality disctinction rarely matters, but when it
matters, it really matters. For example, sometimes perfdash needs
to answer the internal question "what are all the possible relevant
dimensions for this value" -- such a quesion is tractable only within
the context of a metric, where the answer is the list of dimensions
that have appeared in recorded samples. Across multiple metrics,
the question is nonsense -- the answer would be the universe of
all concievable dimensions for all conceivable metrics. Also,
dimensions are named with simple text labels. The same label might
mean different things in different metrics, but within a single
metric it is expected always to mean the same thing -- to refer
to the same conceptual dimension.)
The only requirement is that the recording process be consistent:
if two samples occupy the same true position in the
many-dimensional conceptual space, they should be recorded with
the same position.
(Unfortunately, this makes it difficult to invent and record
new dimentions for a metric after measurement has begun.
TODO(kr): this might be unnecessarily conservative; see if we can
relax this restriction/assumption.)

Values are _sampled_ on some fixed period (e.g. 1 minute).
The individual samples are then aggregated into buckets of a fixed size,
and buckets are further aggregated into other buckets of some
larger fixed size. Each bucket size is an integer multiple of
the next-largest size. Each granularity is subject to a
retention policy, for example: raw samples are stored for
one hour, one-hour buckets are stored for one week, and one-day
buckets are stored for one year.
Note that what perfdash considers an individual sample may already
be an aggregate value collected and aggregated by some other
process.

All samples and buckets are collectively referred to as records.

Glossary

Axis: (see Graph; see Dimension)

Dimension: A form of classification for sampled values.
In a nontrivial system, values can be organized along many
dimensions: which OS process they came from, which cohort
they relate to, which datacenter they came from, etc.
(The various dimensions may or may not actually be
independent or dependent, perfdash makes no assumption
either way.)
Dimensions are represented in perfdash as strings;
they are the keys in key-value labels on stored records.
Example dimensions might be "proc", "dc", "app", etc.

Position: A set of dimensions along with their values.
For example: "proc=abc123", or "app=cored dc=sjc".
A record's position is all its dimension labels and values
put together.

Graph: a 2D visual representation of data stored in perfdash.

Metric: records are organized into directions along many
dimensions; metrics are used to organize the dimensions.
A metric is identified by a string, its name.

Plot: (see Graph)

Values

We provide three types:

  gauge      float64 (with weight uint64)
  counter    uint64
  histogram  HDR-histogram

A gauge records weighted samples of a time-varying scalar value.
Its values are assumed to be locally normally distributed.
(That is, on any single position, inside a sufficiently
small time window, the values are normally distributed.)
For individual samples, the weight is often 1; weights are
summed during aggregation.
Examples: connection count, heap size, sensor temperature.
Its default aggregate function is weighted arithmetic mean.

A counter records samples of an integral value that is
expected to form a monotone function on each position.
Counters can be sampled as absolute counter value or as
deltas. Internally they are stored as deltas.
Note that a single position can be very specific,
often confined to a single OS process. This allows meaningful
measurement of process-local values such as allocations
and CPU time.
Examples: user signups, bytes allocated, cumulative GC
pause time.
Its default aggregate function is sum (of deltas), that is,
the max (of absolute values).

A histogram records an HDR histogram[1].
It is an aggregate of samples with non-normal distribution.
The most common use for this is to measure the latency
of a process -- how long does function f take?
It allows recovering the approximate value at an arbitrary
quantile in the distribution with bounded error, using a fixed
amount of storage regardless of the number of measurements.
Its only aggregate function is the HDR histogram merge operation.

Aggregation

Perfdash stores individual samples as well as aggregated "buckets"
in exponentially increasing sizes. (The size of bucket is often
referred to as granularity.) This allows it to recover an
accurate aggregate value over an arbitrary time interval with
O(log(n)) reads (for any given position in any given metric).

It has different retention durations for different granularities;
larger buckets are kept for a longer time. This means that
time-resolution gets coarser as you go farther back in history,
but the storage cost is much smaller.

Visualization

(Just some vague notes for now.)

- single vertical column of graphs
- x-axes should be aligned vertically
- y-axes don't need to be consistent
- historgram graph could be time-series +
  hdrhistogram-style aggregate graph
  (stacked on top of each other?)
- search box for user-defined dimensions
- make common things convenient, esp time intervals
  (e.g. ff to live view; select timespan, move
  time interval keeping same timespan)

[1]: http://hdrhistogram.org/

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
