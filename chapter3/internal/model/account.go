package model

const (
	UnKnow int8 = 0
	Female int8 = 1
	Male int8 = 2
	
)

type Account struct {
	Id int64
	Name string 
	Gender int8
	Email string
	Address string
}

func (a *Account) SetGender(g int8) {
	switch g {
	case Female:
		a.Gender = Female
	case Male:
		a.Gender = Male
	default:
		a.Gender = UnKnow
	}
}