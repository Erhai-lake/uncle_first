package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

// User 用户模型
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository 用户数据仓库接口
type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id int64) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
}

// MySQLUserRepository MySQL实现
type MySQLUserRepository struct {
	db *sql.DB
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository(db *sql.DB) UserRepository {
	return &MySQLUserRepository{db: db}
}

// CreateUser 创建用户
func (r *MySQLUserRepository) CreateUser(user *User) error {
	query := `INSERT INTO users (name, email, created_at, updated_at) 
              VALUES (?, ?, ?, ?)`

	now := time.Now()
	result, err := r.db.Exec(query, user.Name, user.Email, now, now)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

// GetUserByID 根据ID获取用户
func (r *MySQLUserRepository) GetUserByID(id int64) (*User, error) {
	query := `SELECT id, name, email, created_at, updated_at 
              FROM users WHERE id = ?`

	user := &User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 用户不存在
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// GetAllUsers 获取所有用户
func (r *MySQLUserRepository) GetAllUsers() ([]*User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("❌ 数据库查询行关闭失败: ", err)
		}
	}(rows)

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser 更新用户
func (r *MySQLUserRepository) UpdateUser(user *User) error {
	query := `UPDATE users SET name = ?, email = ?, updated_at = ? 
              WHERE id = ?`

	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, user.Name, user.Email, user.UpdatedAt, user.ID)
	return err
}

// DeleteUser 删除用户
func (r *MySQLUserRepository) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
