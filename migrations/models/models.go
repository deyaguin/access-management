package models

import "time"

type Service struct {
	ID      int `gorm:"AUTO_INCREMENT;not null;unique"`
	Name    string
	Actions []Action
}

type Action struct {
	ID          int `gorm:"AUTO_INCREMENT;not null;unique"`
	Name        string
	Permissions []Permission
}

type User struct {
	ID        int `gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Groups    []Group  `gorm:"many2many:user_groups;save_associations:false" json:"-"`
	Policies  []Policy `gorm:"many2many:user_policies;save_associations:false" json:"-"`
}

type Group struct {
	ID        int `gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Users     []User   `gorm:"many2many:user_groups;save_associations:false" json:"-"`
	Policies  []Policy `gorm:"many2many:group_policies;save_associations:false" json:"-"`
}

type Permission struct {
	ID        int `gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Resourse  string
	Access    bool
	ActionID  int
	PolicyID  int
}

type Policy struct {
	ID          int `gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	Groups      []Group `gorm:"many2many:group_policies;save_associations:false" json:"-"`
	Users       []User  `gorm:"many2many:user_policies;save_associations:false" json:"-"`
	Permissions []Permission
}
