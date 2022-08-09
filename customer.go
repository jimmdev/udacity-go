package main

type Customer struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Role string `json:"role,omitempty"`

	Email string `json:"email,omitempty"`

	Phone uint64 `json:"phone,omitempty"`

	Contacted bool `json:"contacted,omitempty"`
}
