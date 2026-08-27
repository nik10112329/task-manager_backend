package user

type CreateUserRequest struct {
	Email            string   `json:"email"`
	DisplayName      string   `json:"display_name"`
	PhotoURL         string   `json:"photo_url"`
	PhoneNumber      string   `json:"phone_number"`
	SignInMethod     string   `json:"sign_in_method"`
	IDLinkedProvider string   `json:"id_linked_provider"`
	LinkedProviders  []string `json:"linked_providers"`
}
