package project

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`                 // Уникальный идентификатор
	Email     string    `json:"email" db:"email"`           // Email пользователя
	Password  string    `json:"-" db:"password"`            // Хэш пароля
	Name      string    `json:"name" db:"name"`             // Имя пользователя
	Role      string    `json:"role" db:"role"`             // Роль (admin, employee)
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Время создания
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // Время последнего обновления
}
