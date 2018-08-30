package bindings

// PersonPair binding
type PersonPair struct {
	PersonPair []string `json:"person_pair"`
}

// Request binding
type Request struct {
	Requestor string `json:"requestor"`
	Target 	  string `json:"target"`
}

// GetFriends binding
type GetFriends struct {
	Email  string `json:"email"`
}

// GetFollowers binding
type GetFollowers struct {
	Email string `json:"email"`
	Post  string `json:"update"`
}
