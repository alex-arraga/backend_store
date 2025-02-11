package services

import "github.com/alex-arraga/backend_store/repositories"

type ServicesContainer struct {
	UserSrv UserService
}

func LoadServices(repo *repositories.RepositoryContainer) *ServicesContainer {
	return &ServicesContainer{
		UserSrv: NewUserService(repo.UserRepo),
	}
}