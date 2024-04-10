// Package dbapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package dbapi

import (
	"encoding/json"

	"github.com/oapi-codegen/runtime"
)

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Message string `json:"message"`
}

// GenreResponse defines model for GenreResponse.
type GenreResponse struct {
	GenreId int    `json:"genre_id"`
	Name    string `json:"name"`
}

// Group defines model for Group.
type Group struct {
	GroupId   int    `json:"group_id"`
	GroupName string `json:"group_name"`
}

// GroupPostRecommendationRequest defines model for GroupPostRecommendationRequest.
type GroupPostRecommendationRequest struct {
	Accuracy float32 `json:"accuracy"`
	MovieId  int     `json:"movie_id"`
}

// GroupRequest defines model for GroupRequest.
type GroupRequest struct {
	Name string `json:"name"`
}

// GroupResponse defines model for GroupResponse.
type GroupResponse struct {
	GroupId int    `json:"group_id"`
	Name    string `json:"name"`
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// MovieReccomendationResponse defines model for MovieReccomendationResponse.
type MovieReccomendationResponse struct {
	Accuracy float32 `json:"accuracy"`
	MovieId  int     `json:"movie_id"`
}

// UserPostCreateRequest defines model for UserPostCreateRequest.
type UserPostCreateRequest struct {
	Mail string `json:"mail"`
	Name string `json:"name"`
}

// UserPostMovieRequest defines model for UserPostMovieRequest.
type UserPostMovieRequest struct {
	Rating int `json:"rating"`
}

// UserPostRecommendationRequest defines model for UserPostRecommendationRequest.
type UserPostRecommendationRequest struct {
	Accuracy float32 `json:"accuracy"`
	MovieId  int     `json:"movie_id"`
}

// UserPostUpdateRequest defines model for UserPostUpdateRequest.
type UserPostUpdateRequest struct {
	Name string `json:"name"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Mail   string `json:"mail"`
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
}

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []ValidationError_Loc_Item `json:"loc"`
	Msg  string                     `json:"msg"`
	Type string                     `json:"type"`
}

// ValidationErrorLoc0 defines model for .
type ValidationErrorLoc0 = string

// ValidationErrorLoc1 defines model for .
type ValidationErrorLoc1 = int

// ValidationError_Loc_Item defines model for ValidationError.loc.Item.
type ValidationError_Loc_Item struct {
	union json.RawMessage
}

// CreateGroupGroupsPostJSONRequestBody defines body for CreateGroupGroupsPost for application/json ContentType.
type CreateGroupGroupsPostJSONRequestBody = GroupRequest

// AddGroupRecommendationsGroupsGroupIdRecommendationsPostJSONRequestBody defines body for AddGroupRecommendationsGroupsGroupIdRecommendationsPost for application/json ContentType.
type AddGroupRecommendationsGroupsGroupIdRecommendationsPostJSONRequestBody = GroupPostRecommendationRequest

// CreateUserUsersPostJSONRequestBody defines body for CreateUserUsersPost for application/json ContentType.
type CreateUserUsersPostJSONRequestBody = UserPostCreateRequest

// AddUserInfosUsersUserIdPostJSONRequestBody defines body for AddUserInfosUsersUserIdPost for application/json ContentType.
type AddUserInfosUsersUserIdPostJSONRequestBody = UserPostUpdateRequest

// AddMovieToWatchedUsersUserIdMoviesMovieIdPostJSONRequestBody defines body for AddMovieToWatchedUsersUserIdMoviesMovieIdPost for application/json ContentType.
type AddMovieToWatchedUsersUserIdMoviesMovieIdPostJSONRequestBody = UserPostMovieRequest

// AddRecommendationsUsersUserIdRecommendationsPostJSONRequestBody defines body for AddRecommendationsUsersUserIdRecommendationsPost for application/json ContentType.
type AddRecommendationsUsersUserIdRecommendationsPostJSONRequestBody = UserPostRecommendationRequest

// AsValidationErrorLoc0 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc0
func (t ValidationError_Loc_Item) AsValidationErrorLoc0() (ValidationErrorLoc0, error) {
	var body ValidationErrorLoc0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc0 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) FromValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc0 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsValidationErrorLoc1 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc1
func (t ValidationError_Loc_Item) AsValidationErrorLoc1() (ValidationErrorLoc1, error) {
	var body ValidationErrorLoc1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc1 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) FromValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc1 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ValidationError_Loc_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ValidationError_Loc_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
