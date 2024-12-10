package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
	"video-call-project/pkg/utils"
)

type userService struct {
	userRepo UserRepositoryInterface
}

func NewUserService(userRepo UserRepositoryInterface) *userService {
	return &userService{userRepo: userRepo}
}
func (u *userService) registerService(c context.Context, user User) (*User, error) {
	//  dau vao chi co ten va mat khau , email
	// kiem tra xem email co ton tai hay khong
	// neu ton tai thi  ra ve error

	user.UserId = uuid.New().String()
	user.Role = Member
	user.CreateTime = time.Now()

	_, ok, _ := u.userRepo.GetUserRepoByEmail(c, user.Email)
	if ok == true {
		return nil, errors.New("user exists!")
	}
	// ma hoa mat khau
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashPassWord)

	//  chua ton tai thi them vao csdl
	_, err = u.userRepo.CreateUserRepo(c, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *userService) Login(c context.Context, login Login) (*InFormation, *string, error) {

	//  kiem tra ton tai cua user
	info, ok, err := u.userRepo.GetUserRepoByEmail(c, login.Email)
	if err != nil || ok == false {
		return nil, nil, err
	}
	log.Printf(info.Password)
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(login.Password))
	if err != nil {
		return nil, nil, errors.New("account not exists or password error ")
	}

	// tao token cho account
	token, err := utils.GenerateJWTTokenAuth(info.Email, info.UserId, info.Role)
	if err != nil {
		return nil, nil, err
	}

	information := &InFormation{
		Username: info.Username,
		Role:     info.Role,
	}
	// tra ve token cho user
	return information, &token, err
}
func (u userService) GetUserByIdService(c context.Context, id interface{}) (*User, error) {
	user, _, err := u.userRepo.GetUserRepoById(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userService) UpdateUserByIdService(c context.Context, id interface{}, user UpdateUser) error {
	// kiem tra xem user co ton tai hay khong
	//  kiem tra dinh dang cua email nua
	check, err := CheckEmail(*user.Email)
	if check == false {
		return errors.New("Incorrect email format !")
	}
	//  kiem tra xem email co ton tai hay khong ,  neu ton tai thi bawt user pahi cap nhat lai email
	_, exists, err := u.userRepo.GetUserRepoByEmail(c, *user.Email)
	if exists == true {
		return errors.New("email was  exists!")
	}
	//  check email ton tai
	//  if ton tai thi cap nhat cho no
	err = u.userRepo.UpdateUserRepo(c, id, user)
	if err != nil {
		return err
	}
	return nil

}
func CheckEmail(email string) (bool, error) {
	if strings.Contains(email, "@") {
		return true, nil
	}
	return false, nil
}
func (u *userService) ChangePassword(c context.Context, uuid interface{}, password ChangePassword) error {
	//  so sanh pass cu xem co dung khogn sau do dung khong
	// neu dung thi cap nhat pass moi cho user la duoc
	user, err := u.GetUserByIdService(c, uuid)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password.OldPassword))
	if err != nil {
		return err
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	//  cap nhat password moi vao cho csdl
	err = u.userRepo.ChangePassword(c, uuid, user.Password)
	if err != nil {
		return err
	}
	return nil

}
