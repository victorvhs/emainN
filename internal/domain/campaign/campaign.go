package campaign

import "time"

type Contatc struct {
	Email string
}
type Campaign struct {
	ID               string
	Name             string
	CreatedOn        time.Time
	ShortDescription string
	Content          string
	Contacts         []Contatc
}
