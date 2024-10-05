package services

type Service struct {
	client Client
	db     BDFilm
}

func New(client Client, db BDFilm) *Service {
	return &Service{
		client: client,
		db:     db,
	}
}
