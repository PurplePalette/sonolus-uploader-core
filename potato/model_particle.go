/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

// Particle - A particle provides particle effect elements to levels / It defines particle effect used for specific level https://github.com/NonSpicyBurrito/sonolus-wiki/wiki/Particle
type Particle struct {

	// english and number only name for searching
	Name string `json:"name"`

	// Reserved for future update. current default is 1.
	Version int32 `json:"version"`

	// base title of this content
	Title string `json:"title"`

	// something footer(ex. featuring xyz) for this content
	Subtitle string `json:"subtitle"`

	// author of this content
	Author string `json:"author"`

	Thumbnail SonolusResourceLocator `json:"thumbnail"`

	Data SonolusResourceLocator `json:"data"`

	Texture SonolusResourceLocator `json:"texture"`

	// 独自要素: データを作成したエポックミリ秒(ソート用)
	CreatedTime int32 `json:"createdTime,omitempty"`

	// 独自要素: データを更新したエポックミリ秒(ソート用)
	UpdatedTime int32 `json:"updatedTime,omitempty"`

	// 独自要素: 譜面作成者のユーザーID
	UserId string `json:"userId,omitempty"`
}
