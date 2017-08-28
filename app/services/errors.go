package services

import (
	"fmt"
)

type UserNotFound struct {
	UserID int
}

type GroupNotFound struct {
	GroupID int
}

type PolicyNotFound struct {
	PolicyID int
}

type PermissionNotFound struct {
	PermissionID int
}

func NewUserNotFoundError(UserID int) *UserNotFound {
	return &UserNotFound{
		UserID,
	}
}

func (err *UserNotFound) Error() string {
	return fmt.Sprintf("User with indetifier '%s' not found", err.UserID)
}

func NewGroupNotFoundError(GroupID int) *GroupNotFound {
	return &GroupNotFound{
		GroupID,
	}
}

func (err *GroupNotFound) Error() string {
	return fmt.Sprintf("Group with indetifier '%s' not found", err.GroupID)
}

func NewPolicyNotFoundError(PolicyID int) *PolicyNotFound {
	return &PolicyNotFound{
		PolicyID,
	}
}

func (err *PolicyNotFound) Error() string {
	return fmt.Sprintf("Policy with indetifier '%s' not found", err.PolicyID)
}

func NewPermissionNotFoundError(PermissionID int) *PermissionNotFound {
	return &PermissionNotFound{
		PermissionID,
	}
}

func (err *PermissionNotFound) Error() string {
	return fmt.Sprintf("Permission with indetifier '%s' not found", err.PermissionID)
}

//type UserValidationError struct {
//	Message string
//}

//func (err *UserValidationError) Error() string {
//
//}
