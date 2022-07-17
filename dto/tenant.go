package dto

type CreateTenantRequest struct {
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	ExpiresIn map[string]uint64
}
