package city

type BaseCity struct {
	FullName string
	Unlocode string
}

type City struct {
	BaseCity
	Id int
}
