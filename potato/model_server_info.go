/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

// ServerInfo - Response / Request struct of getServerInfo
type ServerInfo struct {
	Levels []Level `json:"levels"`

	Skins []Skin `json:"skins"`

	Backgrounds []Background `json:"backgrounds"`

	Effects []Effect `json:"effects"`

	Particles []Particle `json:"particles"`

	Engines []Engine `json:"engines"`
}
