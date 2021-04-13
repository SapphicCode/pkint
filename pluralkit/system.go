package pluralkit

import "time"

// System describes a PluralKit system.
// https://pluralkit.me/api/#system-model
type System struct {
	ID      string     `json:"id"`
	Created *time.Time `json:"created"`

	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Tag         *string `json:"tag,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	Timezone    *string `json:"tz,omitempty"`

	// TODO: Privacy options
}
