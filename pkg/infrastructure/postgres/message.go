package postgres

type Service struct {
	client *PGPool
}

func NewService(pool *PGPool) *Service {
	return &Service{client: pool}
}
