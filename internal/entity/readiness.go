package entity

type Readiness struct {
	errors []error
}

func NewReadiness(errors []error) *Readiness {
	return &Readiness{
		errors: errors,
	}
}

func (r Readiness) IsReady() bool {
	return len(r.errors) == 0
}

func (r Readiness) Reasons() []error {
	return r.errors
}
