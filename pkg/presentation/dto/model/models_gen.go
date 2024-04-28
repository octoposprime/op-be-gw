// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package presentation

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// Don't use
type AuthOps struct {
	// Don't use
	Auth bool `json:"Auth"`
}

// Dlr represents the basic values of the dlr
type Dlr struct {
	// Id is the id of the dlr
	ID string `json:"Id"`
	// DlrType is the type of the dlr
	DlrType DlrType `json:"DlrType"`
	// DlrStatus is the status of the dlr
	DlrStatus DlrStatus `json:"DlrStatus"`
	// DlrBase is the base values of the dlr
	DlrBase *DlrBase `json:"DlrBase"`
	// DlrCore is the core values of the dlr
	DlrCore *DlrCore `json:"DlrCore"`
	// CreatedAt is the create time of the dlr
	CreatedAt time.Time `json:"CreatedAt"`
	// UpdatedAt is the update time of the dlr
	UpdatedAt time.Time `json:"UpdatedAt"`
}

// DlrBase represents the base values of the dlr
type DlrBase struct {
	// Tags are the tag labels of the dlr
	Tags []string `json:"Tags"`
}

// DlrBase represents the base values of the dlr
type DlrBaseInput struct {
	// Tags are the tag labels of the dlr
	Tags []string `json:"Tags,omitempty"`
}

// DlrCore represents the core values of the dlr
type DlrCore struct {
	// DlrData is the dlr data of the dlr
	DlrData string `json:"DlrData"`
}

// DlrCore represents the core values of the dlr
type DlrCoreInput struct {
	// DlrData is the dlr data of the dlr
	DlrData *string `json:"DlrData,omitempty"`
}

// DlrFilterInput is used for filtering the Dlrs
type DlrFilterInput struct {
	// Id is the id of the dlr
	ID *string `json:"Id,omitempty"`
	// DlrType is the type of the dlr
	DlrType *DlrType `json:"DlrType,omitempty"`
	// DlrStatus is the status of the dlr
	DlrStatus *DlrStatus `json:"DlrStatus,omitempty"`
	// Tags are the tag labels of the dlr
	Tags []string `json:"Tags,omitempty"`
	// CreatedAtFrom is the start value of the create time of the dlr
	CreatedAtFrom *time.Time `json:"CreatedAtFrom,omitempty"`
	// CreatedAtTo is the end value of the create time of the dlr
	CreatedAtTo *time.Time `json:"CreatedAtTo,omitempty"`
	// UpdatedAtFrom is the start value of the update time of the dlr
	UpdatedAtFrom *time.Time `json:"UpdatedAtFrom,omitempty"`
	// UpdatedAtTo is the end value of the update time of the dlr
	UpdatedAtTo *time.Time `json:"UpdatedAtTo,omitempty"`
	// SearchText is the value of the full text search
	SearchText *string `json:"SearchText,omitempty"`
	// SortType is the sorting type. It can be only ASC or DESC
	SortType *string `json:"SortType,omitempty"`
	// SortField is the sortable field of the dlr
	SortField *DlrSortField `json:"SortField,omitempty"`
	// PageInput is used for pagination
	Pagination *PageInput `json:"Pagination,omitempty"`
}

// Dlr represents the basic values of the dlr
type DlrInput struct {
	// Id is the id of the dlr
	ID *string `json:"Id,omitempty"`
	// DlrStatus is the status of the dlr
	DlrType *DlrType `json:"DlrType,omitempty"`
	// DlrStatus is the status of the dlr
	DlrStatus *DlrStatus `json:"DlrStatus,omitempty"`
	// DlrBase is the base values of the dlr
	DlrBase *DlrBaseInput `json:"DlrBase,omitempty"`
	// DlrCore is the core values of the dlr
	DlrCore *DlrCoreInput `json:"DlrCore,omitempty"`
}

type Dlrs struct {
	// The total number of Dlrs that match the filter
	Total int32 `json:"Total"`
	// The Dlrs that match the filter
	Dlrs []*Dlr `json:"Dlrs"`
}

// Error represents the built-in error message
type Error struct {
	Error string `json:"Error"`
}

// Login Request provides to authenticate in the system
type LoginRequest struct {
	// UserName is the user name of the user
	UserName string `json:"UserName"`
	// Email is the email address of the user
	Email string `json:"Email"`
	// Password is the password of the user
	Password string `json:"Password"`
}

type LoginRequestInput struct {
	// UserName is the user name of the user
	UserName *string `json:"UserName,omitempty"`
	// Email is the email address of the user
	Email *string `json:"Email,omitempty"`
	// Password is the password of the user
	Password *string `json:"Password,omitempty"`
}

type Mutation struct {
}

// PageInput is used for pagination
type PageInput struct {
	// Limit is the max count of return values
	Limit *int32 `json:"Limit,omitempty"`
	// Offset is the starting point of limitation
	Offset *int32 `json:"Offset,omitempty"`
}

// Permission represents the built-in policies
type Permission struct {
	// Policy represents the built-in policy type
	Policy string `json:"Policy"`
}

// PermissionGroup represents the built-in groups of permissions
type PermissionGroup struct {
	// GroupName is the name of the group of permission
	GroupName string `json:"GroupName"`
	// Permissions represents the built-in policies
	Permissions []*Permission `json:"Permissions"`
}

type Query struct {
}

// Role represents the built-in role types
type Role struct {
	// RoleName is the name of the role of permission
	RoleName string `json:"RoleName"`
	// PermissionGroup represents the built-in groups of permissions
	PermissionGroups []*PermissionGroup `json:"PermissionGroups"`
}

// RoleFilterInput is used for filtering the Roles
type RoleFilterInput struct {
	// RoleName is the name of the role of permission
	RoleName *string `json:"RoleName,omitempty"`
}

type Subscription struct {
}

// UserToken represents the object of a token for authentication
type Token struct {
	// AuthenticationToken is the authentication token
	AuthenticationToken string `json:"AuthenticationToken"`
	// RefreshToken is the refresh token
	RefreshToken string `json:"RefreshToken"`
}

type TokenInput struct {
	// AuthenticationToken is the authentication token
	AuthenticationToken string `json:"AuthenticationToken"`
	// RefreshToken is the refresh token
	RefreshToken string `json:"RefreshToken"`
}

// User represents the basic values of the user
type User struct {
	// Id is the id of the user
	ID string `json:"Id"`
	// UserName is the user name of the user
	UserName string `json:"UserName"`
	// Email is the email address of the user
	Email string `json:"Email"`
	// Role is the role of the user
	Role string `json:"Role"`
	// UserType is the type of the user
	UserType UserType `json:"UserType"`
	// UserStatus is the status of the user
	UserStatus UserStatus `json:"UserStatus"`
	// UserBase is the base values of the user
	UserBase *UserBase `json:"UserBase"`
	// CreatedAt is the create time of the user
	CreatedAt time.Time `json:"CreatedAt"`
	// UpdatedAt is the update time of the user
	UpdatedAt time.Time `json:"UpdatedAt"`
}

// UserBase represents the base values of the user
type UserBase struct {
	// Tags are the tag labels of the user
	Tags []string `json:"Tags"`
	// FirstName is the first name of the user
	FirstName string `json:"FirstName"`
	// LastName is the last name of the user
	LastName string `json:"LastName"`
}

// UserBase represents the base values of the user
type UserBaseInput struct {
	// Tags are the tag labels of the user
	Tags []string `json:"Tags,omitempty"`
	// FirstName is the first name of the user
	FirstName *string `json:"FirstName,omitempty"`
	// LastName is the last name of the user
	LastName *string `json:"LastName,omitempty"`
}

// UserFilterInput is used for filtering the Users
type UserFilterInput struct {
	// Id is the id of the user
	ID *string `json:"Id,omitempty"`
	// UserName is the user name of the user
	UserName *string `json:"UserName,omitempty"`
	// Email is the email address of the user
	Email *string `json:"Email,omitempty"`
	// UserType is the type of the user
	UserType *UserType `json:"UserType,omitempty"`
	// UserStatus is the status of the user
	UserStatus *UserStatus `json:"UserStatus,omitempty"`
	// Tags are the tag labels of the user
	Tags []string `json:"Tags,omitempty"`
	// FirstName is the first name of the user
	FirstName *string `json:"FirstName,omitempty"`
	// LastName is the last name of the user
	LastName *string `json:"LastName,omitempty"`
	// CreatedAtFrom is the start value of the create time of the user
	CreatedAtFrom *time.Time `json:"CreatedAtFrom,omitempty"`
	// CreatedAtTo is the end value of the create time of the user
	CreatedAtTo *time.Time `json:"CreatedAtTo,omitempty"`
	// UpdatedAtFrom is the start value of the update time of the user
	UpdatedAtFrom *time.Time `json:"UpdatedAtFrom,omitempty"`
	// UpdatedAtTo is the end value of the update time of the user
	UpdatedAtTo *time.Time `json:"UpdatedAtTo,omitempty"`
	// SearchText is the value of the full text search
	SearchText *string `json:"SearchText,omitempty"`
	// SortType is the sorting type. It can be only ASC or DESC
	SortType *string `json:"SortType,omitempty"`
	// SortField is the sortable field of the user
	SortField *UserSortField `json:"SortField,omitempty"`
	// PageInput is used for pagination
	Pagination *PageInput `json:"Pagination,omitempty"`
}

// User represents the basic values of the user
type UserInput struct {
	// Id is the id of the user
	ID *string `json:"Id,omitempty"`
	// UserName is the user name of the user
	UserName *string `json:"UserName,omitempty"`
	// Email is the email address of the user
	Email *string `json:"Email,omitempty"`
	// Role is the role of the user
	Role *string `json:"Role,omitempty"`
	// UserStatus is the status of the user
	UserType *UserType `json:"UserType,omitempty"`
	// UserStatus is the status of the user
	UserStatus *UserStatus `json:"UserStatus,omitempty"`
	// UserBase is the base values of the user
	UserBase *UserBaseInput `json:"UserBase,omitempty"`
}

// UserPassword represents the password values of the user
type UserPassword struct {
	// Id is the id of the user password
	ID string `json:"Id"`
	// UserId is the id of the user
	UserID string `json:"UserId"`
	// Password is the password of the user
	Password string `json:"Password"`
	// PasswordStatus is the status of password of the user
	PasswordStatus PasswordStatus `json:"PasswordStatus"`
}

// UserPassword represents the password values of the user
type UserPasswordInput struct {
	// Id is the id of the user password
	ID *string `json:"Id,omitempty"`
	// UserId is the id of the user
	UserID string `json:"UserId"`
	// Password is the password of the user
	Password string `json:"Password"`
}

type Users struct {
	// The total number of Users that match the filter
	Total int32 `json:"Total"`
	// The Users that match the filter
	Users []*User `json:"Users"`
}

type DlrSortField string

const (
	// No Type
	DlrSortFieldNone DlrSortField = "None"
	// Id
	DlrSortFieldID DlrSortField = "Id"
	// Created Time
	DlrSortFieldCreatedAt DlrSortField = "CreatedAt"
	// Updated Time
	DlrSortFieldUpdatedAt DlrSortField = "UpdatedAt"
)

var AllDlrSortField = []DlrSortField{
	DlrSortFieldNone,
	DlrSortFieldID,
	DlrSortFieldCreatedAt,
	DlrSortFieldUpdatedAt,
}

func (e DlrSortField) IsValid() bool {
	switch e {
	case DlrSortFieldNone, DlrSortFieldID, DlrSortFieldCreatedAt, DlrSortFieldUpdatedAt:
		return true
	}
	return false
}

func (e DlrSortField) String() string {
	return string(e)
}

func (e *DlrSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DlrSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DlrSortField", str)
	}
	return nil
}

func (e DlrSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DlrStatus string

const (
	// No Type
	DlrStatusNone DlrStatus = "NONE"
	// Active
	DlrStatusActive DlrStatus = "ACTIVE"
	// Inactive
	DlrStatusInactive DlrStatus = "INACTIVE"
)

var AllDlrStatus = []DlrStatus{
	DlrStatusNone,
	DlrStatusActive,
	DlrStatusInactive,
}

func (e DlrStatus) IsValid() bool {
	switch e {
	case DlrStatusNone, DlrStatusActive, DlrStatusInactive:
		return true
	}
	return false
}

func (e DlrStatus) String() string {
	return string(e)
}

func (e *DlrStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DlrStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DlrStatus", str)
	}
	return nil
}

func (e DlrStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DlrType string

const (
	// No Type
	DlrTypeNone DlrType = "NONE"
)

var AllDlrType = []DlrType{
	DlrTypeNone,
}

func (e DlrType) IsValid() bool {
	switch e {
	case DlrTypeNone:
		return true
	}
	return false
}

func (e DlrType) String() string {
	return string(e)
}

func (e *DlrType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DlrType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DlrType", str)
	}
	return nil
}

func (e DlrType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PasswordStatus string

const (
	// No Type
	PasswordStatusNone PasswordStatus = "NONE"
	// Active
	PasswordStatusActive PasswordStatus = "ACTIVE"
	// Inactive
	PasswordStatusInactive PasswordStatus = "INACTIVE"
	// Auto Generated
	PasswordStatusAutoGenerated PasswordStatus = "AUTO_GENERATED"
	// Change Required
	PasswordStatusChangeRequired PasswordStatus = "CHANGE_REQUIRED"
	// Expired
	PasswordStatusExpired PasswordStatus = "EXPIRED"
)

var AllPasswordStatus = []PasswordStatus{
	PasswordStatusNone,
	PasswordStatusActive,
	PasswordStatusInactive,
	PasswordStatusAutoGenerated,
	PasswordStatusChangeRequired,
	PasswordStatusExpired,
}

func (e PasswordStatus) IsValid() bool {
	switch e {
	case PasswordStatusNone, PasswordStatusActive, PasswordStatusInactive, PasswordStatusAutoGenerated, PasswordStatusChangeRequired, PasswordStatusExpired:
		return true
	}
	return false
}

func (e PasswordStatus) String() string {
	return string(e)
}

func (e *PasswordStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PasswordStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PasswordStatus", str)
	}
	return nil
}

func (e PasswordStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserSortField string

const (
	// No Type
	UserSortFieldNone UserSortField = "None"
	// Id
	UserSortFieldID UserSortField = "Id"
	// Name
	UserSortFieldName UserSortField = "Name"
	// Created Time
	UserSortFieldCreatedAt UserSortField = "CreatedAt"
	// Updated Time
	UserSortFieldUpdatedAt UserSortField = "UpdatedAt"
)

var AllUserSortField = []UserSortField{
	UserSortFieldNone,
	UserSortFieldID,
	UserSortFieldName,
	UserSortFieldCreatedAt,
	UserSortFieldUpdatedAt,
}

func (e UserSortField) IsValid() bool {
	switch e {
	case UserSortFieldNone, UserSortFieldID, UserSortFieldName, UserSortFieldCreatedAt, UserSortFieldUpdatedAt:
		return true
	}
	return false
}

func (e UserSortField) String() string {
	return string(e)
}

func (e *UserSortField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSortField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserSortField", str)
	}
	return nil
}

func (e UserSortField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	// No Type
	UserStatusNone UserStatus = "NONE"
	// Active
	UserStatusActive UserStatus = "ACTIVE"
	// Inactive
	UserStatusInactive UserStatus = "INACTIVE"
)

var AllUserStatus = []UserStatus{
	UserStatusNone,
	UserStatusActive,
	UserStatusInactive,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusNone, UserStatusActive, UserStatusInactive:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserType string

const (
	// No Type
	UserTypeNone UserType = "NONE"
	// Admin
	UserTypeAdmin UserType = "ADMIN"
	// User
	UserTypeUser UserType = "USER"
)

var AllUserType = []UserType{
	UserTypeNone,
	UserTypeAdmin,
	UserTypeUser,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeNone, UserTypeAdmin, UserTypeUser:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
