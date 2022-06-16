package freight

type containerType int

const (
	TWENTY_STANDARD containerType = iota
	FORTY_STANDARD
	FORTY_HIGH_CUBE
	FORTY_FIVE_HIGH_CUBE
)

func (s containerType) ConvertToString() string {
	switch s {
	case 0:
		return "20DC"
	case 1:
		return "40DC"
	case 2:
		return "40HC"
	case 3:
		return "45HC"
	default:
		return ""

	}
}

type Contact struct {
	Url         string
	PhoneNumber string
	AgentName   string
}

type BaseFreight struct {
	FromCity      string
	ToCity        string
	ContainerType string
	UsdPrice      int
	Line          string
	LineImage     string
	Contacts      Contact
}

type GetFreight struct {
	FromCity      string
	ToCity        string
	ContainerType containerType
	Limit         int
}

func NewGetFreight(fromCity string, toCity string, containerType containerType, limit int) GetFreight {
	return GetFreight{FromCity: fromCity, ToCity: toCity, ContainerType: containerType, Limit: limit}
}

type IRepository interface {
	GetFrieght(freight GetFreight) ([]BaseFreight, error)
}
