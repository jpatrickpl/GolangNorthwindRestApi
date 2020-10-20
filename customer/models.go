package customer

type Customer struct {
	ID             int    `json:"id"`
	FirstName      string `json:"firstame"`
	LastName       string `json:"lastname"`
	Address        string `json:"address"`
	BussinessPhone string `json:"bussinessphone"`
	City           string `json:"city"`
	Company        string `json:"company"`
}

type CustomerList struct {
	Data         []*Customer `json:"data"`
	TotalRecords int64
}
