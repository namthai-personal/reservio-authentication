package user

import (
	userDTO "reservio-be/app/user/dto"

	"github.com/gin-gonic/gin"
)

type UserPagination struct {
	Items []User `json:"items"`
	Page  int    `json:"page"`
}

var (
	userRepo UserRepo
)

// ListUser godoc
//
//	@Summary	Get all users
//	@Tags		Internal :: User
//	@Router		/user/ [get]
//
// @Success 200 {object} UserPagination
func ListUser(c *gin.Context) {
	users, _ := userRepo.List()

	c.JSON(200, UserPagination{users, 1})
}

// GetUser godoc
//
//	@Summary	Get single user by ID
//	@Tags		Internal :: User
//	@Param		id	path	string	true	"Account ID"
//	@Router		/user/{id} [get]
//
// @Success 200 {object} User
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, _ := userRepo.Get(id)
	c.JSON(200, user)
}

// CreateUser godoc
//
//	@Summary	Get single user by ID
//	@Tags		Internal :: User
//
// @Param data body user.CreateUserDTO true "Create user input"
//
//	@Router		/user [post]
//
// @Success 200 {object} User
func CreateUser(c *gin.Context) {
	// jsonData, err := io.ReadAll(c.Request.Body)
	userBody := userDTO.CreateUserDTO{}
	err := c.ShouldBindJSON(&userBody)

	if err != nil {
		c.JSON(500, "Server Error")
	}

	newUser, _ := userRepo.Create(userBody)

	c.JSON(200, newUser)
}

// Login godoc
//
//	@Summary	Get single user by ID
//	@Tags		Internal :: Auth
//
// @Param data body user.LoginDTO true "Login input"
//
//	@Router		/user/login [post]
//
// @Success 200 {object} bool
func Login(c *gin.Context) {
	// jsonData, err := io.ReadAll(c.Request.Body)
	userBody := userDTO.CreateUserDTO{}
	err := c.ShouldBindJSON(&userBody)

	if err != nil {
		c.JSON(500, "Server Error")
	}

	newUser, _ := userRepo.Login(userBody)

	c.JSON(200, newUser)
}
