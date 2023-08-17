package service

import (
	"gorm.io/gorm"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

type SessionService interface {
	VerifyPassword(username string, password string) (bool, *model.User)
	GetUser(id int) (*model.User, error)
}

func NewSessionService(db *gorm.DB) (*SessionServiceImpl, error) {
	return &SessionServiceImpl{
		ServiceBase: NewServiceBase(),
		db:          db,
	}, nil
}

type SessionServiceImpl struct {
	ServiceBase
	db *gorm.DB
}

func (s *SessionServiceImpl) VerifyPassword(username string, password string) (ok bool, user *model.User) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		user, err = s.userDao.GetUserByUserName(tx, username)
		if err != nil {
			return err
		}
		ok = user.VerifyPassword(password)
		return nil
	})
	if err != nil || !ok {
		return false, nil
	}
	return true, user
}

func (s *SessionServiceImpl) GetUser(currentUserID int) (res *model.User, err error) {
	s.db.Transaction(func(tx *gorm.DB) error {
		res, err = s.userDao.GetUserByID(tx, currentUserID)
		return err
	})
	return
}
