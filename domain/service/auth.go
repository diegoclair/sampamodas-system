package service

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/twinj/uuid"
)

type authService struct {
	svc *Service
}

func newAuthService(svc *Service) AuthService {
	return &authService{
		svc: svc,
	}
}

func (s *authService) Signup(signup entity.User) (err error) {
	logger.Info("Signup: Process Started")
	defer logger.Info("Signup: Process Finished")

	signup.UUID = uuid.NewV4().String()

	hasher := md5.New()
	hasher.Write([]byte(signup.Password))
	signup.Password = hex.EncodeToString(hasher.Sum(nil))

	err = s.svc.dm.MySQL().User().CreateUser(signup)
	if err != nil {
		logger.Error("Signup.CreateUser: ", err)
		return err
	}

	return nil
}
