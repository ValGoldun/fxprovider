package healthcheck

type FailPolicy string

const (
	OneOf FailPolicy = "ONE_OF"
	All              = "ALL"
)
