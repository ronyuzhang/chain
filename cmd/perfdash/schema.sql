CREATE TABLE hdrhistogram (
	cluster text NOT NULL,
	process text NOT NULL,
	t timestamptz NOT NULL,
	metric text NOT NULL,

	min bigint NOT NULL,
	max bigint NOT NULL,
	sig int NOT NULL, -- in [1,5]
	data bytea NOT NULL
);
