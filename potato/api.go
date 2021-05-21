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
	"context"
	"net/http"
)

// BackgroundsApiRouter defines the required methods for binding the api requests to a responses for the BackgroundsApi
// The BackgroundsApiRouter implementation should parse necessary information from the http request,
// pass the data to a BackgroundsApiServicer to perform the required actions, then write the service results to the http response.
type BackgroundsApiRouter interface {
	AddBackground(http.ResponseWriter, *http.Request)
	EditBackground(http.ResponseWriter, *http.Request)
	GetBackground(http.ResponseWriter, *http.Request)
	GetBackgroundList(http.ResponseWriter, *http.Request)
}

// EffectsApiRouter defines the required methods for binding the api requests to a responses for the EffectsApi
// The EffectsApiRouter implementation should parse necessary information from the http request,
// pass the data to a EffectsApiServicer to perform the required actions, then write the service results to the http response.
type EffectsApiRouter interface {
	AddEffect(http.ResponseWriter, *http.Request)
	EditEffect(http.ResponseWriter, *http.Request)
	GetEffect(http.ResponseWriter, *http.Request)
	GetEffectList(http.ResponseWriter, *http.Request)
}

// EnginesApiRouter defines the required methods for binding the api requests to a responses for the EnginesApi
// The EnginesApiRouter implementation should parse necessary information from the http request,
// pass the data to a EnginesApiServicer to perform the required actions, then write the service results to the http response.
type EnginesApiRouter interface {
	AddEngine(http.ResponseWriter, *http.Request)
	EditEngine(http.ResponseWriter, *http.Request)
	GetEngine(http.ResponseWriter, *http.Request)
	GetEngineList(http.ResponseWriter, *http.Request)
}

// InfoApiRouter defines the required methods for binding the api requests to a responses for the InfoApi
// The InfoApiRouter implementation should parse necessary information from the http request,
// pass the data to a InfoApiServicer to perform the required actions, then write the service results to the http response.
type InfoApiRouter interface {
	EditInfo(http.ResponseWriter, *http.Request)
	GetServerInfo(http.ResponseWriter, *http.Request)
}

// LevelsApiRouter defines the required methods for binding the api requests to a responses for the LevelsApi
// The LevelsApiRouter implementation should parse necessary information from the http request,
// pass the data to a LevelsApiServicer to perform the required actions, then write the service results to the http response.
type LevelsApiRouter interface {
	AddLevel(http.ResponseWriter, *http.Request)
	EditLevel(http.ResponseWriter, *http.Request)
	GetLevel(http.ResponseWriter, *http.Request)
	GetLevelList(http.ResponseWriter, *http.Request)
}

// ParticlesApiRouter defines the required methods for binding the api requests to a responses for the ParticlesApi
// The ParticlesApiRouter implementation should parse necessary information from the http request,
// pass the data to a ParticlesApiServicer to perform the required actions, then write the service results to the http response.
type ParticlesApiRouter interface {
	AddParticle(http.ResponseWriter, *http.Request)
	GetParticle(http.ResponseWriter, *http.Request)
	GetParticleList(http.ResponseWriter, *http.Request)
	PatchParticlesParticleName(http.ResponseWriter, *http.Request)
}

// SkinsApiRouter defines the required methods for binding the api requests to a responses for the SkinsApi
// The SkinsApiRouter implementation should parse necessary information from the http request,
// pass the data to a SkinsApiServicer to perform the required actions, then write the service results to the http response.
type SkinsApiRouter interface {
	AddSkin(http.ResponseWriter, *http.Request)
	EditSkin(http.ResponseWriter, *http.Request)
	GetSkin(http.ResponseWriter, *http.Request)
	GetSkinList(http.ResponseWriter, *http.Request)
}

// TestsApiRouter defines the required methods for binding the api requests to a responses for the TestsApi
// The TestsApiRouter implementation should parse necessary information from the http request,
// pass the data to a TestsApiServicer to perform the required actions, then write the service results to the http response.
type TestsApiRouter interface {
	GetTestServerInfo(http.ResponseWriter, *http.Request)
	GetTestsBackgrounds(http.ResponseWriter, *http.Request)
	GetTestsEffects(http.ResponseWriter, *http.Request)
	GetTestsEngines(http.ResponseWriter, *http.Request)
	GetTestsLevels(http.ResponseWriter, *http.Request)
	GetTestsParticles(http.ResponseWriter, *http.Request)
	GetTestsSkins(http.ResponseWriter, *http.Request)
}

// UsersApiRouter defines the required methods for binding the api requests to a responses for the UsersApi
// The UsersApiRouter implementation should parse necessary information from the http request,
// pass the data to a UsersApiServicer to perform the required actions, then write the service results to the http response.
type UsersApiRouter interface {
	EditUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	GetUserList(http.ResponseWriter, *http.Request)
	GetUserServerInfo(http.ResponseWriter, *http.Request)
	GetUsersBackgrounds(http.ResponseWriter, *http.Request)
	GetUsersEffects(http.ResponseWriter, *http.Request)
	GetUsersEngines(http.ResponseWriter, *http.Request)
	GetUsersLevels(http.ResponseWriter, *http.Request)
	GetUsersParticles(http.ResponseWriter, *http.Request)
	GetUsersSkins(http.ResponseWriter, *http.Request)
}

// BackgroundsApiServicer defines the api actions for the BackgroundsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type BackgroundsApiServicer interface {
	AddBackground(context.Context, string, Background) (ImplResponse, error)
	EditBackground(context.Context, string, Background) (ImplResponse, error)
	GetBackground(context.Context, string) (ImplResponse, error)
	GetBackgroundList(context.Context, string, int32, string) (ImplResponse, error)
}

// EffectsApiServicer defines the api actions for the EffectsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type EffectsApiServicer interface {
	AddEffect(context.Context, string, Effect) (ImplResponse, error)
	EditEffect(context.Context, string, Effect) (ImplResponse, error)
	GetEffect(context.Context, string) (ImplResponse, error)
	GetEffectList(context.Context, string, int32, string) (ImplResponse, error)
}

// EnginesApiServicer defines the api actions for the EnginesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type EnginesApiServicer interface {
	AddEngine(context.Context, string, Engine) (ImplResponse, error)
	EditEngine(context.Context, string, Engine) (ImplResponse, error)
	GetEngine(context.Context, string) (ImplResponse, error)
	GetEngineList(context.Context, string, int32, string) (ImplResponse, error)
}

// InfoApiServicer defines the api actions for the InfoApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type InfoApiServicer interface {
	EditInfo(context.Context, ServerInfo) (ImplResponse, error)
	GetServerInfo(context.Context) (ImplResponse, error)
}

// LevelsApiServicer defines the api actions for the LevelsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type LevelsApiServicer interface {
	AddLevel(context.Context, string, Level) (ImplResponse, error)
	EditLevel(context.Context, string, Level) (ImplResponse, error)
	GetLevel(context.Context, string) (ImplResponse, error)
	GetLevelList(context.Context, string, int32, string) (ImplResponse, error)
}

// ParticlesApiServicer defines the api actions for the ParticlesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ParticlesApiServicer interface {
	AddParticle(context.Context, string, Particle) (ImplResponse, error)
	EditParticle(context.Context, string, Particle) (ImplResponse, error)
	GetParticle(context.Context, string) (ImplResponse, error)
	GetParticleList(context.Context, string, int32, string) (ImplResponse, error)
}

// SkinsApiServicer defines the api actions for the SkinsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type SkinsApiServicer interface {
	AddSkin(context.Context, string, Skin) (ImplResponse, error)
	EditSkin(context.Context, string, Skin) (ImplResponse, error)
	GetSkin(context.Context, string) (ImplResponse, error)
	GetSkinList(context.Context, string, int32, string) (ImplResponse, error)
}

// TestsApiServicer defines the api actions for the TestsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TestsApiServicer interface {
	GetTestServerInfo(context.Context, string) (ImplResponse, error)
	GetTestsBackgrounds(context.Context, string, string, int32, string) (ImplResponse, error)
	GetTestsEffects(context.Context, string, string, int32, string) (ImplResponse, error)
	GetTestsEngines(context.Context, string, string, int32, string) (ImplResponse, error)
	GetTestsLevels(context.Context, string, string, int32, string) (ImplResponse, error)
	GetTestsParticles(context.Context, string, string, int32, string) (ImplResponse, error)
	GetTestsSkins(context.Context, string, string, int32, string) (ImplResponse, error)
}

// UsersApiServicer defines the api actions for the UsersApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UsersApiServicer interface {
	EditUser(context.Context, string, User) (ImplResponse, error)
	GetUser(context.Context, string) (ImplResponse, error)
	GetUserList(context.Context) (ImplResponse, error)
	GetUserServerInfo(context.Context, string) (ImplResponse, error)
	GetUsersBackgrounds(context.Context, string, string, int32, string) (ImplResponse, error)
	GetUsersEffects(context.Context, string, string, int32, string) (ImplResponse, error)
	GetUsersEngines(context.Context, string, string, int32, string) (ImplResponse, error)
	GetUsersLevels(context.Context, string, string, int32, string) (ImplResponse, error)
	GetUsersParticles(context.Context, string, string, int32, string) (ImplResponse, error)
	GetUsersSkins(context.Context, string, string, int32, string) (ImplResponse, error)
}
