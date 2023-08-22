package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/qdriven/qfluent-go/pkg/orm"
	service2 "github.com/qdriven/qfluent-go/pkg/service"
	"reflect"
)

// CreateHandler handles
//
//	POST /T
//
// creates a new model T, responds with the created model T if successful.
//
// Request body:
//   - {...}  // fields of the model T
//
// Response:
//   - 200 OK: { T: {...} }
//   - 400 Bad Request: { error: "request band failed" }
//   - 422 Unprocessable Entity: { error: "create process failed" }
func CreateHandler[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model T
		if err := c.ShouldBindJSON(&model); err != nil {
			logger.WithContext(c).WithError(err).
				Warn("CreateHandler: Bind failed")
			ResponseError(c, CodeBadRequest, err)
			return
		}
		logger.WithContext(c).Tracef("CreateHandler: Create %#v", model)
		err := service2.Create(c, &model, service2.IfNotExist())
		if err != nil {
			logger.WithContext(c).WithError(err).
				Warn("CreateHandler: Create failed")
			ResponseError(c, CodeProcessFailed, err)
			return
		}
		c.JSON(200, SuccessResponseBody(model))
	}
}

// CreateNestedHandler handles
//
//	POST /P/:parentIDRouteParam/T
//
// where:
//   - P is the parent model, T is the child model
//   - parentIDRouteParam is the route param name of the parent model P
//   - field is the field name of the child model T in the parent model P
//
// responds with the updated parent model P
//
// Request body:
//   - {...}  // fields of the child model T
//
// Response:
//   - 200 OK: { P: {...} }
//   - 400 Bad Request: { error: "request band failed" }
//   - 422 Unprocessable Entity: { error: "create process failed" }
func CreateNestedHandler[P orm.Model, T orm.Model](parentIDRouteParam string, field string) gin.HandlerFunc {
	return func(c *gin.Context) {
		parentID := c.Param(parentIDRouteParam)
		if parentID == "" {
			ResponseError(c, CodeBadRequest, ErrMissingParentID)
			return
		}

		var child T
		if err := c.ShouldBindJSON(&child); err != nil {
			logger.WithContext(c).WithError(err).
				Warn("CreateNestedHandler: Bind failed")
			ResponseError(c, CodeBadRequest, err)
			return
		}

		if _, childID := child.Identity(); !reflect.ValueOf(childID).IsZero() {
			// child id exists: add to join table, but do not update child's fields
			logger.WithField("childID", childID).Debug("CreateNestedHandler: child model has ID, add to join table, but do not update child's fields")
			if err := service2.GetByID[T](c, childID, &child); err != nil {
				logger.WithContext(c).WithError(err).
					WithField("note", "try to query it because child id exists in request").
					Warn("CreateNestedHandler: GetByID[Child] failed")
				ResponseError(c, CodeNotFound, err)
				return
			}
		}
		// else: id is not set: create new child

		var parent P
		if err := service2.GetByID[P](c, parentID, &parent); err != nil {
			logger.WithContext(c).WithError(err).
				Warn("CreateNestedHandler: GetByID[Parent] failed")
			ResponseError(c, CodeNotFound, err)
			return
		}

		logger.WithContext(c).
			Tracef("CreateNestedHandler: Create %#v, parent=%#v", child, parent)

		//field := strings.ToUpper(field)[:1] + field[1:]
		field := nameToField(field, parent)

		err := service2.Create(c, &child, service2.NestInto(&parent, field))
		if err != nil {
			logger.WithContext(c).WithError(err).
				Warn("CreateNestedHandler: CreateNest failed")
			ResponseError(c, CodeProcessFailed, err)
			return
		}
		ResponseSuccess(c, parent)
	}
}
