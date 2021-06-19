package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/karanbirsingh7/go-greenlight/internal/data"
	"github.com/karanbirsingh7/go-greenlight/internal/validator"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.PrintInfo("POST /v1/users", nil)

	// temp struct to hold request data
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	// Use the Password.Set() method to generate and store the hashed and plaintext
	// passwords.
	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// VALIDATE inputs
	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// CREATE USER
	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		// If we get a ErrDuplicateEmail error, use the v.AddError() method to manually
		// add a message to the validator instance, and then call our
		// failedValidationResponse() helper.

		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// SEND registeration mail
	app.background(func() {
		err = app.mailer.Send(user.Email, "user_welcome.tmpl", user)
		if err != nil {
			app.logger.PrintError(err, nil)
		}
	})

	// USER CREATED
	err = app.writeJSON(w, http.StatusCreated, envelope{
		"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "show user with ID")

}

func (app *application) listUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list users")

}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update user with ID")

}
func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete user with ID")

}
