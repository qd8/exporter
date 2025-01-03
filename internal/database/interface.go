package database

func (s *service) Delete(model interface{}) error {
	return s.gormDB.Delete(model).Error
}

func (s *service) Save(model interface{}) error {
	return s.gormDB.Save(model).Error
}

func (s *service) Create(model interface{}) error {
	return s.gormDB.Create(model).Error
}

func (s *service) Find(model interface{}) {
	s.gormDB.Find(model)
}

func (s *service) FindById(id uint, model interface{}) {
	s.gormDB.Where("id = ?", id).Find(model)
}