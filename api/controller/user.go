package controller

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userRepo models.UserRepository
}

func NewUserController(userRepo models.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

// GetUsers 获取所有用户
func (c *UserController) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.userRepo.GetAllUsers()
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "获取用户列表失败")
			return
		}

		utils.Success(w, users)
	}
}

// GetUser 获取单个用户
func (c *UserController) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.Error(w, http.StatusBadRequest, "无效的用户ID")
			return
		}

		user, err := c.userRepo.GetUserByID(id)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "获取用户失败")
			return
		}

		if user == nil {
			utils.Error(w, http.StatusNotFound, "用户不存在")
			return
		}

		utils.Success(w, user)
	}
}

// CreateUser 创建用户
func (c *UserController) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := utils.ParseJSON(r, &user); err != nil {
			utils.Error(w, http.StatusBadRequest, "无效的请求数据")
			return
		}

		// 简单验证
		if user.Name == "" || user.Email == "" {
			utils.Error(w, http.StatusBadRequest, "姓名和邮箱不能为空")
			return
		}

		if err := c.userRepo.CreateUser(&user); err != nil {
			utils.Error(w, http.StatusInternalServerError, "创建用户失败")
			return
		}

		utils.Success(w, user)
	}
}
