package repository

type MemoryRepo struct {
	data map[string]int
}

func (m *MemoryRepo) GetAge(name string) (int, error) {
	return m.data[name], nil
}

func NewRepository() *MemoryRepo {
	return &MemoryRepo{
		data: map[string]int{
			"budougumi0617": 33,
			"john":          30,
		},
	}
}
