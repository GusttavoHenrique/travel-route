package route

type Service struct {
	Repo *Repository
}

func (s *Service) Find(origin string, destination string, price float64) (*[]route, error) {
	routes, err := s.Repo.Find(origin, destination, price)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func (s *Service) FindAll() []*Route {
	return s.Repo.FindAll()
}

func (s *Service) Save(origin string, destination string, price float64) error {
	_, err := s.Repo.Save(origin, destination, price)
	if err != nil {
		return err
	}
	return nil
}
