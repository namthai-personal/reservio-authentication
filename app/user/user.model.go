package user

import (
	userDTO "reservio-be/app/user/dto"
	hasher "reservio-be/common/hasher"
	"reservio-be/infrastructure/database"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	QUERY_LIMIT = 1000
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string    `json:"updatedBy"`
}

func (User) TableName() string {
	return "user"
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (repo *UserRepo) List() ([]User, error) {
	users := []User{}

	result := database.DB.Limit(QUERY_LIMIT).Order("created_at DESC").Find(&users)

	return users, result.Error
}

func (repo *UserRepo) Get(id string) (User, error) {
	user := User{}
	result := database.DB.Where("id = ?", id).First(&user)
	color.Yellow("%d", result.RowsAffected)
	return user, result.Error
}

func (repo *UserRepo) Create(props userDTO.CreateUserDTO) (User, error) {
	createTime := time.Now()
	user := User{
		ID:       uuid.New().String(),
		Name:     props.Name,
		Username: props.Username,
		Password: hasher.HashSha256(props.Password),

		CreatedAt: createTime,
		UpdatedAt: createTime,
		CreatedBy: "admin",
		UpdatedBy: "admin",
	}

	result := database.DB.Create(user)
	color.Yellow("%d", result.RowsAffected)

	return user, result.Error
}

func (repo *UserRepo) Login(props userDTO.CreateUserDTO) (bool, error) {
	user := User{}
	hashPw := hasher.HashSha256(props.Password)
	result := database.DB.Where("username = ? AND password = ?", props.Username, hashPw).First(&user)
	if result.Error != nil {
		return false, result.Error
	}

	return true, result.Error
}
