package confluentcloud

import (
	"fmt"
	"net/url"
  "time"
  "log"
)

type ServiceAccount struct {
	ID          int    `json:"id"`
	Name        string `json:"service_name"`
	Description string `json:"service_description"`
}

type ServiceAccountsResponse struct {
	ServiceAccounts []ServiceAccount `json:"users"`
}

type ServiceAccountResponse struct {
	ServiceAccount ServiceAccount `json:"user"`
}
type ServiceAccountCreateRequestW struct {
	ServiceAccount *ServiceAccountCreateRequest `json:"user"`
}
type ServiceAccountCreateRequest struct {
	Name        string `json:"service_name"`
	Description string `json:"service_description"`
}
type ServiceAccountDeleteRequestW struct {
	ServiceAccount ServiceAccountDeleteRequest `json:"user"`
}
type ServiceAccountDeleteRequest struct {
	ID int `json:"id"`
}

func (c *Client) CreateServiceAccount(request *ServiceAccountCreateRequest) (*ServiceAccount, error) {
	rel, err := url.Parse("service_accounts")
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	response, err := c.NewRequest().
		SetBody(&ServiceAccountCreateRequestW{ServiceAccount: request}).
		SetResult(&ServiceAccountResponse{}).
		SetError(&ErrorResponse{}).
		Post(u.String())

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, fmt.Errorf("service_accounts: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	return &response.Result().(*ServiceAccountResponse).ServiceAccount, nil
}

func healthCheck() {
    for i := 0; i < 3000; i++ {
        log.Printf("beep")
        time.Sleep(time.Second)
    }
}

func (c *Client) ListServiceAccounts() ([]ServiceAccount, error) {
  log.Printf("client lib LOGGGGGGGGGGGGGGGGGGGGGGG")

	rel, err := url.Parse("service_accounts")
	if err != nil {
		return []ServiceAccount{}, err
	}

  go healthCheck()

  log.Printf("set timeout")
  c.client.SetTimeout(10 * time.Second)

	u := c.BaseURL.ResolveReference(rel)

  log.Printf("request")
	response, err := c.NewRequest().
		SetResult(&ServiceAccountsResponse{}).
		SetError(&ErrorResponse{}).
		Get(u.String())

  log.Printf("after request")

  log.Printf("request again")

  r := c.NewRequest().
		SetResult(&ServiceAccountsResponse{}).
		SetError(&ErrorResponse{})

  log.Printf("before GET")

  response, err = r.Get(u.String())
  //log.Printf("put instead")
	//response, err = r.Put(u.String())

  log.Printf("after request again")

	if err != nil {
		return []ServiceAccount{}, err
	}

	if response.IsError() {
		return []ServiceAccount{}, fmt.Errorf("service_accounts: %s", response.Error().(*ErrorResponse).Error.Message)
	}
	return response.Result().(*ServiceAccountsResponse).ServiceAccounts, nil
}

func (c *Client) DeleteServiceAccount(id int) error {
	rel, err := url.Parse("service_accounts")
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	request := ServiceAccountDeleteRequest{
		ID: id,
	}

	response, err := c.NewRequest().
		SetError(&ErrorResponse{}).
		SetBody(&ServiceAccountDeleteRequestW{ServiceAccount: request}).
		Delete(u.String())

	if err != nil {
		return err
	}

	if response.IsError() {
		return fmt.Errorf("delete service account: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	return nil
}

// OKAY I THINK THERE IS NO SUCH ROUTE
//
//
// read service account
//func (c *Client) GetServiceAccount(id, account_id string) (*ServiceAccount, error) {
func (c *Client) GetServiceAccount(id string) (*ServiceAccount, error) {
	rel, err := url.Parse(fmt.Sprintf("service_accounts/%s", id))
  log.Printf("[ERROR] TESTTTTT: %s", err)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	fmt.Println(rel.String())

	response, err := c.NewRequest().
		SetResult(&ServiceAccountResponse{}).
		//SetQueryParam("account_id", account_id).
		SetError(&ErrorResponse{}).
		Get(u.String())

  log.Printf("[ERROR] TESTTTTT: %s", err)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
    log.Printf("[ERROR] its an error: %s", response.Error().(*ErrorResponse).Error.Message)
		return nil, fmt.Errorf("get service account: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	return &response.Result().(*ServiceAccountResponse).ServiceAccount, nil
}
