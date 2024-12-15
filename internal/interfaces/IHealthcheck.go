package interfaces

type IHealthcheckServices interface {
	HealthcheckServices() (string, error)
}

type IHealthcheckRepo interface {
}
