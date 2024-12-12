package user

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type ItemRole int

const (
	Admin ItemRole = iota
	Member
	Viewer
)

var allItemRole = [3]string{"admin", "member", "viewer"}

func (i ItemRole) String() string {
	return allItemRole[i]
}

func ParseStr2ItemRole(s string) (ItemRole, error) {
	//chuyen het ve chu thuong
	s = strings.ToLower(s)
	for i := range allItemRole {
		if allItemRole[i] == s {
			return ItemRole(i), nil
		}
	}
	return 0, errors.New("invalid item status")
}

func (i *ItemRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal value: %v", value)
	}
	v, err := ParseStr2ItemRole(string(bytes))
	if err != nil {
		return fmt.Errorf("failed to unmarshal value: %v", value)
	}
	*i = v
	return nil
}

func (i ItemRole) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i ItemRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", i.String())), nil
}

func (i *ItemRole) UnmarshalJSON(b []byte) error {
	str := strings.ReplaceAll(string(b), "\"", "")
	itemValue, err := ParseStr2ItemRole(str)
	log.Printf("unmarshaljson: %v ", str)
	log.Printf("itemValue:  %v ", itemValue)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value: %v", str)
	}
	*i = itemValue
	return nil
}

//	role ow day ci co quyen la role cua system  ,
//
// con su ly role trong nhom thi ta su dung ow
// ben kia de kien tra xem ai o trong nhom do
type User struct {
	UserId     string    `json:"userid" bson:"userid"`
	Username   string    `json:"username" bson:"username"`
	Email      string    `json:"email" bson:"email"`
	Password   string    `json:"password" bson:"password"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
	Role       *ItemRole `json:"role-system" bson:"role-system"`
}
type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
type InFormation struct {
	Username string   `json:"username" bson:"username"`
	Role     ItemRole `json:"role-system" bson:"role-system"`
}
type UpdateUser struct {
	UserName   *string    `json:"username,omitempty" bson:"username,omitempty"`
	Email      *string    `json:"email,omitempty" bson:"email,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty" bson:"update_time,omitempty"`
}
type ChangePassword struct {
	OldPassword string `json:"old_password,omitempty" bson:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty" bson:"new_password,omitempty"`
}

type JsonRole struct {
	Role string `json:"role-system" bson:"role-system"`
}
