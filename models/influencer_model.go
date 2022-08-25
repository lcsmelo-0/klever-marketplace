package models

type Influencer struct {
	Id                 string `json:"id,omitempty" bson:"_id,omitempty"`
	Name               string `json:"name,omitempty" validate:"required"`
	Email              string `json:"email,omitempty" validate:"required"`
	InstagramProfile   string `json:"instagramProfile,omitempty" validate:"required"`
	InstagramFollowers string `json:"instagramfollowers,omitempty" validate:"required"`
	ProfileDescription string `json:"profileDescription,omitempty" validate:"required"`
}
