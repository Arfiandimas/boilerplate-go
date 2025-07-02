package healthchecks

type HealthchecksIO interface {
	Ping(uuid string) error
}
