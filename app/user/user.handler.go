package user

import (
	userDTO "reservio-be/app/user/dto"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
	TotalPage  int `json:"totalPage"`
}

type UserPagination struct {
	Pagination Pagination `json:"pagination"`
	Items      []User     `json:"items"`
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
	users, e := userRepo.List()

	if e != nil {
		c.AbortWithError(500, e)
		return
	}
	c.JSON(200, UserPagination{Pagination{1, 1, 1, 1}, users})
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

	user, e := userRepo.Get(id)
	if e != nil {
		c.AbortWithError(500, e)
		return
	}
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
// @Failure 400 {object} middlewares.GeneralError
func CreateUser(c *gin.Context) {
	userBody := userDTO.CreateUserDTO{}
	err := c.ShouldBindJSON(&userBody)

	if err != nil {
		c.Error(err)
		return
	}

	newUser, err := userRepo.Create(userBody)

	if err != nil {
		c.Error(err)
		return
	}

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
