/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

// Level - A level provides a list of entities and powered by scripted behavior in engine, with skin, effect, background and particle, to create gameplay experience for players / It defines level for actual user play. It includes all data to play one level. https://github.com/NonSpicyBurrito/sonolus-wiki/wiki/Level
type Level struct {

	// english and number only name for searching
	Name string `json:"name"`

	// Reserved for future update. current default is 1.
	Version int32 `json:"version"`

	// Difficulty of the level
	Rating int32 `json:"rating"`

	Engine Engine `json:"engine"`

	UseSkin LevelUseSkin `json:"useSkin"`

	UseBackground LevelUseBackground `json:"useBackground"`

	UseEffect LevelUseEffect `json:"useEffect"`

	UseParticle LevelUseParticle `json:"useParticle"`

	// base title of this content
	Title string `json:"title"`

	// artist names of original music
	Artists string `json:"artists"`

	// author of this content
	Author string `json:"author"`

	Cover SonolusResourceLocator `json:"cover"`

	Bgm SonolusResourceLocator `json:"bgm"`

	Data SonolusResourceLocator `json:"data"`

	// 独自要素: 楽曲のジャンル
	Genre string `json:"genre,omitempty"`

	// 独自要素: 楽曲が全体公開かどうか
	Public bool `json:"public,omitempty"`

	// 独自要素: 譜面作成者のユーザーID
	UserId string `json:"userId,omitempty"`

	// 独自要素: 譜面内のノーツ数
	Notes int32 `json:"notes,omitempty"`
}
