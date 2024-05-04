// Package todo.
//
// Documentation of user management APIs.
//
// Schemes:
//   - http
//   - https
//
// BasePath: /
// Version: 1.0.0
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// swagger:meta
package todo

//go:generate swagger generate spec -o ./static/swagger.json --scan-models

// swagger:route POST /api/users user createUserRequest
// Create a new user.
// responses:
//   201: createUserResponse

// swagger:parameters createUserRequest
type createUserRequest struct {
	// in:body
	Body User
}

// swagger:response createUserResponse
type createUserResponse struct {
	// in:body
	Body User
}

// swagger:route GET /api/users user listUsersRequest
// Get a list of all users.
// responses:
//   200: listUsersResponse

// swagger:response listUsersResponse
type listUsersResponse struct {
	// in:body
	Body []User
}

// swagger:route GET /api/users/{id} user getUserRequest
// Get a user by ID.
// responses:
//   200: getUserResponse

// swagger:parameters getUserRequest
type getUserRequest struct {
	// in:path
	// required: true
	ID int `json:"id"`
}

// swagger:response getUserResponse
type getUserResponse struct {
	// in:body
	Body User
}

// swagger:route PUT /api/users/{id} user updateUserRequest
// Update a user by ID.
// responses:
//   200: updateUserResponse

// swagger:parameters updateUserRequest
type updateUserRequest struct {
	// in:path
	// required: true
	ID int `json:"id"`

	// in:body
	Body User
}

// swagger:response updateUserResponse
type updateUserResponse struct {
	// in:body
	Body User
}

// swagger:route DELETE /api/users/{id} user deleteUserRequest
// Delete a user by ID.
// responses:
//   204: deleteUserResponse

// swagger:parameters deleteUserRequest
type deleteUserRequest struct {
	// in:path
	// required: true
	ID int `json:"id"`
}

// swagger:response deleteUserResponse
type deleteUserResponse struct{}
