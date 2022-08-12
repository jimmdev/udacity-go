package main

// Customer json representation
type Customer struct {

	// Id is a unique identifier of the person, preferably a UUID
	Id string `json:"id,omitempty"`

	// Name is the full name of the person
	Name string `json:"name,omitempty"`

	// Role is the role this person has
	Role string `json:"role,omitempty"`

	// Email is the e-mail address of the customer
	Email string `json:"email,omitempty"`

	// Phone is the phone number represented only in digital
	Phone uint64 `json:"phone,omitempty"`

	// Contacted indicates if this customer has already been contacted
	Contacted bool `json:"contacted,omitempty"`
}
