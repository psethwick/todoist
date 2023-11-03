package todoist

type Collaborator struct {
	HaveID
	Email    string  `json:"email"`
	FullName string  `json:"full_name"`
	ImageId  *string `json:"image_id"`
	Timezone string  `json:"timezone"`
}

func (c Collaborator) Removable() bool {
	return false
}

type Collaborators []Collaborator

type CollaboratorState struct {
	ProjectID string `json:"project_id"`
	UserID    string `json:"user_id"`
	State     string `json:"state"`
	IsDeleted bool   `json:"is_deleted"`
}

type CollaboratorStates []CollaboratorState
