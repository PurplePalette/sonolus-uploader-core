/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A SkinsAPIController binds http requests to an api service and writes the service results to the http response
type SkinsAPIController struct {
	service SkinsAPIServicer
}

// NewSkinsAPIController creates a default api controller
func NewSkinsAPIController(s SkinsAPIServicer) Router {
	return &SkinsAPIController{service: s}
}

// Routes returns all of the api route for the SkinsAPIController
func (c *SkinsAPIController) Routes() Routes {
	return Routes{
		{
			"AddSkin",
			strings.ToUpper("Post"),
			"/skins/{skinName}",
			c.AddSkin,
		},
		{
			"EditSkin",
			strings.ToUpper("Patch"),
			"/skins/{skinName}",
			c.EditSkin,
		},
		{
			"GetSkinList",
			strings.ToUpper("Get"),
			"/skins/list",
			c.GetSkinList,
		},
		{
			"GetSkin",
			strings.ToUpper("Get"),
			"/skins/{skinName}",
			c.GetSkin,
		},
	}
}

// AddSkin - Add skin
func (c *SkinsAPIController) AddSkin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	skinName := params["skinName"]

	skin := &Skin{}
	if err := json.NewDecoder(r.Body).Decode(&skin); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.AddSkin(r.Context(), skinName, *skin)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// EditSkin - Edit skin
func (c *SkinsAPIController) EditSkin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	skinName := params["skinName"]

	skin := &Skin{}
	if err := json.NewDecoder(r.Body).Decode(&skin); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.EditSkin(r.Context(), skinName, *skin)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSkin - Get skin
func (c *SkinsAPIController) GetSkin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	skinName := params["skinName"]

	result, err := c.service.GetSkin(r.Context(), skinName)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSkinList - Get skin list
func (c *SkinsAPIController) GetSkinList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	localization := query.Get("localization")
	page, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keywords := query.Get("keywords")
	result, err := c.service.GetSkinList(r.Context(), localization, page, keywords)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
