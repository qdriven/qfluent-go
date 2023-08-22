package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qdriven/qfluent-go/pkg/orm"
	handlers2 "github.com/qdriven/qfluent-go/pkg/qgin/handlers"
	"reflect"
)

// Crud add a group of CRUD routes for model T to the base router
// on relativePath. For example, if the base router handles route to
//
//	/base
//
// then the CRUD routes will be:
//
//	GET|POST|PUT|DELETE /base/relativePath
//
// where relativePath is recommended to be the plural form of the model name.
//
// Example:
//
//	r := qgin.Default()
//	Crud[User](r, "/users")
//
// adds the following routes:
//
//	   GET /users/
//	   GET /users/:UserId
//	  POST /users/
//	   PUT /users/:UserId
//	DELETE /users/:UserId
//
// and with options parameters, it's optional to add the following routes:
//   - GetNested()    =>    GET /users/:UserId/friends
//   - CreateNested() =>   POST /users/:UserId/friends
//   - DeleteNested() => DELETE /users/:UserId/friends/:FriendId
func Crud[T orm.Model](base gin.IRouter, relativePath string, options ...CrudOption) gin.IRouter {
	group := base.Group(relativePath)

	if !gin.IsDebugging() { // GIN_MODE == "release"
		logger.WithField("model", getTypeName[T]()).
			//WithField("basePath", base). // we cannot get the base path from qgin.IRouter
			WithField("relativePath", relativePath).
			Info("Crud: Adding CRUD routes for model")
	}

	options = append(options, crud[T]())

	for _, option := range options {
		group = option(group)
	}

	return group
}

// CrudOption is options to construct the router group.
//
// By adding GetNested, CreateNested, DeleteNested to Crud,
// you can add CRUD routes for a nested model (Parent.Child).
//
// Or use CrudNested to add all three options above.
type CrudOption func(group *gin.RouterGroup) *gin.RouterGroup

// crud add CRUD routes for model T to the group:
//
//	   GET /
//	   GET /:idParam
//	  POST /
//	   PUT /:idParam
//	DELETE /:idParam
func crud[T orm.Model]() CrudOption {
	idParam := getIdParam[T]()
	return func(group *gin.RouterGroup) *gin.RouterGroup {
		group.GET("", handlers2.GetListHandler[T]())
		group.GET(fmt.Sprintf("/:%s", idParam), handlers2.GetByIDHandler[T](idParam))

		group.POST("", handlers2.CreateHandler[T]())
		group.PUT(fmt.Sprintf("/:%s", idParam), handlers2.UpdateHandler[T](idParam))
		group.DELETE(fmt.Sprintf("/:%s", idParam), handlers2.DeleteHandler[T](idParam))

		return group
	}
}

// GetNested add a GET route to the group for querying a nested model:
//
//	GET /:parentIdParam/field
func GetNested[P orm.Model, N orm.Model](field string) CrudOption {
	parentIdParam := getIdParam[P]()
	return func(group *gin.RouterGroup) *gin.RouterGroup {
		relativePath := fmt.Sprintf("/:%s/%s", parentIdParam, field)

		if !gin.IsDebugging() { // GIN_MODE == "release"
			logger.WithField("parent", getTypeName[P]()).
				WithField("child", getTypeName[N]()).
				WithField("relativePath", relativePath).
				Info("Crud: Adding GET route for getting nested model")
		}

		group.GET(relativePath,
			handlers2.GetFieldHandler[P](parentIdParam, field),
		)
		// there is no GET /:parentIdParam/:field/:childIdParam,
		// because it is equivalent to GET /:childModel/:childIdParam.
		// So there is also no PUT /:parentIdParam/:field/:childIdParam.
		// It is verbose and unnecessary.
		return group
	}
}

// CreateNested add a POST route to the group for creating a nested model:
//
//	POST /:parentIdParam/field
func CreateNested[P orm.Model, N orm.Model](field string) CrudOption {
	parentIdParam := getIdParam[P]()
	return func(group *gin.RouterGroup) *gin.RouterGroup {
		relativePath := fmt.Sprintf("/:%s/%s", parentIdParam, field)

		if !gin.IsDebugging() { // GIN_MODE == "release"
			logger.WithField("parent", getTypeName[P]()).
				WithField("child", getTypeName[N]()).
				WithField("relativePath", relativePath).
				Info("Crud: Adding POST route for creating nested model")
		}

		group.POST(relativePath,
			handlers2.CreateNestedHandler[P, N](parentIdParam, field),
		)
		return group
	}
}

// DeleteNested add a DELETE route to the group for deleting a nested model:
//
//	DELETE /:parentIdParam/field/:childIdParam
func DeleteNested[P orm.Model, T orm.Model](field string) CrudOption {
	parentIdParam := getIdParam[P]()
	childIdParam := getIdParam[T]()
	return func(group *gin.RouterGroup) *gin.RouterGroup {
		relativePath := fmt.Sprintf("/:%s/%s/:%s", parentIdParam, field, childIdParam)

		if !gin.IsDebugging() { // GIN_MODE == "release"
			logger.WithField("parent", getTypeName[P]()).
				WithField("child", getTypeName[T]()).
				WithField("relativePath", relativePath).
				Info("Crud: Adding DELETE route for deleting nested model")
		}

		group.DELETE(relativePath,
			handlers2.DeleteNestedHandler[P, T](parentIdParam, field, childIdParam),
		)
		return group
	}
}

// CrudNested = GetNested + CreateNested + DeleteNested
func CrudNested[P orm.Model, T orm.Model](field string) CrudOption {
	return func(group *gin.RouterGroup) *gin.RouterGroup {
		group = GetNested[P, T](field)(group)
		group = CreateNested[P, T](field)(group)
		group = DeleteNested[P, T](field)(group)
		return group
	}
}

// getIdParam Model => "ModelID"
func getIdParam[T orm.Model]() string {
	model := *new(T)
	modelName := reflect.TypeOf(model).Name()
	idField, _ := model.Identity()
	idParam := modelName + idField

	return idParam
}

// getTypeName is a helper function to get the type name of a generic type T.
func getTypeName[T any]() string {
	model := *new(T)
	return reflect.TypeOf(model).Name()
}
