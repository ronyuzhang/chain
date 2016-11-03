CREATE TABLE hdrhistogram (
	t timestamptz NOT NULL,
	metric text NOT NULL,
	labels json NOT NULL, -- constrained to map[string]string

	min bigint NOT NULL,
	max bigint NOT NULL,
	sig int NOT NULL, -- in [1,5]
	data bytea NOT NULL
);
