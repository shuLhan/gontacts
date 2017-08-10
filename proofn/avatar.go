package proofn

//
// Avatar define an URI to Proofn contact avatar image.
//
type Avatar struct {
	AvatarPathSmall  string `json:"avatarPathSmall,omitempty"`
	AvatarPathMedium string `json:"avatarPathMedium,omitempty"`
	AvatarPathLarge  string `json:"avatarPathLarge,omitempty"`
}
