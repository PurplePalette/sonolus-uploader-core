/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

// GetEffectResponse - Response struct of getEffect
type GetEffectResponse struct {
	Item Effect `json:"item"`

	Description string `json:"description"`

	Recommended []Effect `json:"recommended"`
}
