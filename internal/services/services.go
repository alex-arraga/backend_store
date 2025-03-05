package services

import "github.com/alex-arraga/backend_store/internal/repositories"

type ServicesContainer struct {
	AuthSrv AuthServices
	UserSrv UserService
}

func LoadServices(repo *repositories.RepositoryContainer) *ServicesContainer {
	return &ServicesContainer{
		AuthSrv: newAuthService(repo.AuthRepo),
		UserSrv: NewUserService(repo.UserRepo),
	}
}
