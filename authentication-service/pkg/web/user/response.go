package user

import (
	"github.com/mysa-fika/go/authentication-service/pkg/domain"
	"github.com/mysa-fika/go/authentication-service/pkg/service"
)

// FetchResponse is the shape of data for a loaded demo record
type FetchResponse struct {
	domain.User
}

// CreateResponse contains the ID post demo creation
type CreateResponse struct {
	ID string
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d domain.User, _ *service.UserService) *FetchResponse {
	return &FetchResponse{d}
}

// NewCreateResponse instantiates a new response when demo is created
func NewCreateResponse(d domain.User, _ *service.UserService) *CreateResponse {
	return &CreateResponse{ID: "demo"}
}
