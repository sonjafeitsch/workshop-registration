package client

type Client interface {
	GetCustomer(mail string) (*int, error)
	CreateCustomer(mail string, firstname string, lastname string) (int, error)
	CreateInvoice(customerId int, productKey string) error
}

func NewClient(baseUrl, token string) Client {
	return clientImpl{
		baseUrl: baseUrl,
		token:   token,
	}
}

type clientImpl struct {
	baseUrl string
	token   string
}

func (c clientImpl) GetCustomer(mail string) (*int, error) {
	panic("implement me")
}

func (c clientImpl) CreateCustomer(mail string, firstname string, lastname string) (int, error) {
	panic("implement me")
}

func (c clientImpl) CreateInvoice(customerId int, productKey string) error {
	panic("implement me")
}
