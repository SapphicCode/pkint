package pluralkit

import "time"

// Member describes a PluralKit member object.
// https://pluralkit.me/api/#member-model
type Member struct {
	ID      string     `json:"id"`
	Created *time.Time `json:"created"`

	Name        *string     `json:"name,omitempty"`
	DisplayName *string     `json:"display_name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Pronouns    *string     `json:"pronouns,omitempty"`
	AvatarURL   *string     `json:"avatar_url,omitempty"`
	Birthday    *time.Time  `json:"birthday,omitempty"`
	ProxyTags   []*ProxyTag `json:"proxy_tags,omitempty"`
	KeepProxy   bool        `json:"keep_proxy"`

	// TODO: Privacy options
}

// ProxyTag describes PluralKit proxy tags.
type ProxyTag struct {
	Prefix *string `json:"prefix"`
	Suffix *string `json:"suffix"`
}
