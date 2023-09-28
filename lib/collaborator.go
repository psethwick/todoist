package todoist

type Collaborator struct {
	Email    string  `json:"email"`
	FullName string  `json:"full_name"`
	Id       string  `json:"id"`
	ImageId  *string `json:"image_id"`
	Timezone string  `json:"timezone"`
}

type Collaborators []Collaborator
