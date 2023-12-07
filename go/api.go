/*
 * AppRepo - Get apps
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)

// AppsApiRouter defines the required methods for binding the api requests to a responses for the AppsApi
// The AppsApiRouter implementation should parse necessary information from the http request,
// pass the data to a AppsApiServicer to perform the required actions, then write the service results to the http response.
type AppsApiRouter interface {
	AppCreatePost(http.ResponseWriter, *http.Request)
	AppDeleteDelete(http.ResponseWriter, *http.Request)
	AppGet(http.ResponseWriter, *http.Request)
	AppUpdatePatch(http.ResponseWriter, *http.Request)
	AppsGet(http.ResponseWriter, *http.Request)
}

// AppsApiServicer defines the api actions for the AppsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AppsApiServicer interface {
	AppCreatePost(context.Context, AppCreatePostRequest) (ImplResponse, error)
	AppDeleteDelete(context.Context, int32) (ImplResponse, error)
	AppGet(context.Context, int32) (ImplResponse, error)
	AppUpdatePatch(context.Context, AppUpdatePatchRequest) (ImplResponse, error)
	AppsGet(context.Context) (ImplResponse, error)
}
