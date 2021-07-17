package models

import (
	"dbsite/security"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// User
//
// swagger:model
type User struct {

	// Base model
	Model `gorm:"embedded"`

	// Nom de l'utilisateur
	// required: true
	FirstName string `json:"firstName" validate:"required"`

	// Prénom de l'utilisateur
	// required: true
	LastName string `json:"lastName" validate:"required"`

	// Mot de passe de l'utilisateur
	// required: true
	Password string `json:"password" validate:"required"`

	// Age de l'utilisateur
	// required: false
	Age uint8 `json:"age" validate:"gte=0,lte=130"`

	// Email de l'utilisateur
	// required: true
	Email string `json:"email" validate:"required,email" gorm:"primaryKey"`

	// Address de l'utilisateur
	// require: false
	Address string `json:"address"`

	// Numéros de téléphone de l'utilisateur
	// require: false
	Phone string `json:"phone"`

	// Role de l'utilisateur
	// pattern: " customer | admin"
	// required: true
	Role string `json:"role" validate:"required,eq=admin|eq=customer"`
}

func (s *User) BeforeCreate(tx *gorm.DB) (err error) {
	res, err := security.HashPwd(s.Password)
	if err != nil {
		return err
	}
	s.Password = res
	return
}

func (s *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") {
		res, err := security.HashPwd(s.Password)
		if err != nil {
			return err
		}
		s.Password = res
	}
	return
}

func (s *User) Retrieve(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	userData := security.GetUserData(c, "user")
	if userData["email"] == s.Email || userData["admin"].(bool) {
		return db.Model(s).First(s), nil
	}
	return nil, fmt.Errorf("access restricted")
}

func (s *User) Update(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	userData := security.GetUserData(c, "user")
	if userData["email"].(string) == s.Email || userData["admin"].(bool) {
		return db.Model(s).Updates(s), nil
	}
	return nil, fmt.Errorf("access restricted")
}

func (s *User) Create(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	return db.FirstOrCreate(s, s), nil
}

func (s *User) Delete(c *fiber.Ctx, db *gorm.DB) (*gorm.DB, error) {
	userData := security.GetUserData(c, "user")
	if userData["email"].(string) == s.Email || userData["admin"].(bool) {
		return db.Where("Email = ?", s.Email).Delete(s), nil
	}
	return nil, fmt.Errorf("access restricted")
}

func (s *User) DeleteListQuery() Query {
	return &UserDeleteQuery{}
}

func (s *User) ListQuery() Query {
	return &UserListQuery{}
}

type UserListQuery struct {
	ToFind  string `query:"tofind"`
	Role    string `query:"role" validate:"omitempty,eq=admin|eq=customer"`
	OrderBy string `query:"orderBy" validate:"omitempty,eq=created_at|eq=updated_at|eq=firstName|eq=lastName|eq=age|eq=address"`
	Limit   int    `query:"limit" validate:"omitempty,gte=0"`
	Offset  int    `query:"offset" validate:"omitempty,gte=0"`
}

func (s *UserListQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {

	users := new([]User)
	tmp := db

	if s.Limit > 0 {
		tmp = tmp.Limit(s.Limit)
	}
	if s.Offset > 0 {
		tmp = tmp.Offset(s.Offset)
	}
	if len(s.ToFind) > 0 {
		tmp = tmp.Where("FirstName LIKE ?", "%"+s.ToFind+"%").Or("LastName LIKE ?", "%"+s.ToFind+"%").Or("Email LIKE ?", "%"+s.ToFind+"%")
	}
	if len(s.Role) > 0 {
		tmp = tmp.Where("Role = ?", s.Role)
	}
	if len(s.OrderBy) > 0 {
		tmp = tmp.Order(s.OrderBy)
	}
	result := tmp.Find(users)
	return users, result
}

type UserDeleteQuery struct {
	Emails []string `query:"emails"`
}

func (s *UserDeleteQuery) Run(db *gorm.DB) (interface{}, *gorm.DB) {
	var users []User

	if result := db.Where("Email IN ?", s.Emails).Find(&users); result.Error != nil {
		return result, nil
	}
	return nil, db.Delete(&users, s.Emails)
}
