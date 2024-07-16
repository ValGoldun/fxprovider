package healthcheck

type Checker interface {
	Source() string
	HealthCheck() error
}

type Checkers []Checker

func New() *Checkers {
	return new(Checkers)
}

func (checkers *Checkers) Add(checker Checker) {
	*checkers = append(*checkers, checker)
}

func (checkers *Checkers) HealthCheck(policy FailPolicy) Health {
	switch policy {
	case OneOf:
		return checkers.HealthCheckFailOneOf()
	case All:
		return checkers.HealthCheckFailAll()
	}

	return checkers.HealthCheckFailAll()
}

// HealthCheckFailOneOf returns error if one of checker returned error
func (checkers *Checkers) HealthCheckFailOneOf() Health {
	var checks = make([]Check, len(*checkers))
	var isOk = true

	for index, checker := range *checkers {
		err := checker.HealthCheck()

		if err != nil {
			isOk = false
		}

		checks[index] = Check{
			Source: checker.Source(),
			Error:  err,
		}

	}

	return Health{
		IsOK:       isOk,
		FailPolicy: OneOf,
		Checks:     checks,
	}
}

// HealthCheckFailAll returns error if all checkers returned error
func (checkers *Checkers) HealthCheckFailAll() Health {
	var checks = make([]Check, len(*checkers))

	if len(*checkers) == 0 {
		return Health{
			IsOK:       true,
			FailPolicy: All,
			Checks:     make([]Check, 0),
		}
	}

	var errCount int

	for index, checker := range *checkers {
		err := checker.HealthCheck()
		if err != nil {
			errCount++
		}

		checks[index] = Check{
			Source: checker.Source(),
			Error:  err,
		}
	}

	var isOk = true

	if len(*checkers) == errCount {
		isOk = false
	}

	return Health{
		IsOK:       isOk,
		FailPolicy: All,
		Checks:     checks,
	}
}
