package user

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type userRepository struct {
	DB *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *userRepository {
	return &userRepository{
		DB: db,
	}

}

// chuc nag dang ki tai khoan
func (u *userRepository) CreateUserRepo(c context.Context, user User) (*User, error) {
	_, err := u.DB.InsertOne(c, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *userRepository) GetUserRepoByEmail(c context.Context, email string) (*User, bool, error) {
	var information User
	err := u.DB.FindOne(c, bson.M{"email": email}).Decode(&information)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, err
		}

	}
	// if user exists then return true
	return &information, true, err
}
func (u userRepository) GetUserRepoById(c context.Context, uuid interface{}) (*User, bool, error) {
	var information User
	err := u.DB.FindOne(c, bson.M{"userid": uuid}).Decode(&information)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, err
		}
	}
	return &information, true, err
}

func (u *userRepository) UpdateUserRepo(c context.Context, id interface{}, user UpdateUser) error {
	// cho phep cap nhat ten va email cua nguoi dung
	updateFields := bson.M{}
	if user.UserName != nil {
		updateFields["username"] = *user.UserName
	}
	if user.Email != nil {
		updateFields["email"] = *user.Email
	}
	if len(updateFields) == 0 {
		return errors.New("update field is empty")
	}
	updateFields["update_time"] = time.Now()
	filter := bson.M{"userid": id}
	update := bson.M{"$set": updateFields}
	_, err := u.DB.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) ChangePassword(c context.Context, uuid interface{}, password string) error {
	// kiem tra a
	filter := bson.M{"userid": uuid}
	update := bson.M{"$set": bson.M{"password": password}}
	_, err := u.DB.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (u userRepository) GetAllUserRepo(c context.Context) ([]*User, error) {
	cursor, err := u.DB.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*User
	err = cursor.All(c, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepository) RoleRightsRepo(c context.Context, uuid, role interface{}) error {
	filter := bson.M{"userid": uuid}
	update := bson.M{"$set": bson.M{"role-system": role}}
	_, err := u.DB.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil
}

type UserRepositoryInterface interface {
	CreateUserRepo(c context.Context, user User) (*User, error)
	GetUserRepoByEmail(c context.Context, email string) (*User, bool, error)
	GetUserRepoById(c context.Context, uuid interface{}) (*User, bool, error)
	UpdateUserRepo(c context.Context, id interface{}, user UpdateUser) error
	ChangePassword(c context.Context, uuid interface{}, password string) error
	GetAllUserRepo(c context.Context) ([]*User, error)
	RoleRightsRepo(c context.Context, uuid, role interface{}) error
}
