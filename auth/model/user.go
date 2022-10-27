package model

type User struct {
	ID          string
	FirstName   string
	LastName    string
	Mobile      *string
	SessionCred string
	Password    string
}

const TERMINATED_SESSION = "TERMINATED"
const OPEN_SESSION = "OPEN"
