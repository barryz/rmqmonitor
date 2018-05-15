package system

// StatsDBCtl is the RabbitMQ statsdb control system
type StatsDBCtl struct {
	name string
}

// NewStatsDBCtl creates a new StatsDBCtl instance
func NewStatsDBCtl() *StatsDBCtl {
	return &StatsDBCtl{
		name: "/sbin/rabbitmqctl",
	}
}

// Reset reset the RabbitMQ statsdb
func (s *StatsDBCtl) Reset() (bool, string, error) {
	output, err := ExecCommand(s.name, []string{"eval", "application:stop(rabbitmq_management), application:start(rabbitmq_management)."})
	return (err == nil), output, err
}

// Terminate terminate the RabbitMQ statsdb
func (s *StatsDBCtl) Terminate() (bool, string, error) {
	output, err := ExecCommand(s.name, []string{"eval", "exit(erlang:whereis(rabbit_mgmt_db), please_terminate)."})
	return (err == nil), output, err
}

// Crash crash the RabbitMQ statsdb
func (s *StatsDBCtl) Crash() (bool, string, error) {
	output, err := ExecCommand(s.name, []string{"eval", "exit(erlang:whereis(rabbit_mgmt_db), please_crash)."})
	return (err == nil), output, err
}
