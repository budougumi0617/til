package server

import "github.com/samber/lo"

type User struct {
	ID   string
	Name string
}

type UserRepo interface {
	List([]string) ([]User, error)
}

type Task struct {
	ID     string
	Title  string
	UserID string
}
type TaskRepo interface {
	All() ([]Task, error)
}

type TaskWithUserName struct {
	Title    string
	UserName string
}

type service struct {
	userRepo UserRepo
	taskRepo TaskRepo
}

func (s *service) ListTaskWithUserName() ([]TaskWithUserName, error) {
	tasks, err := s.taskRepo.All()
	if err != nil {
		return nil, err
	}
	userIDs := lo.Map(tasks, func(t Task, _ int) string { return t.UserID })
	users, err := s.userRepo.List(userIDs)
	if err != nil {
		return nil, err
	}
	userMap := lo.SliceToMap(users, func(u User) (string, User) { return u.ID, u })
	var results []TaskWithUserName
	for _, task := range tasks {
		result := TaskWithUserName{Title: task.Title}
		if u, ok := userMap[task.UserID]; ok {
			result.UserName = u.Name
		}
		results = append(results, result)
	}
	return results, nil
}
