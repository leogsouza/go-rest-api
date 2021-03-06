package entity

// Owner represents the User entity
type Owner struct {
	OwnerData `json:"User"`
}

// OwnerData holds data from the owner entity
type OwnerData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
