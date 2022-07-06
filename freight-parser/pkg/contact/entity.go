package contact

type BaseContact struct {
	Url         string
	Email       string
	AgentName   string
	PhoneNumber string
}

type Contact struct {
	Id int
	BaseContact
}
