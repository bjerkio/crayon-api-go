package helpers

import (
	apiclient "github.com/bjerkio/crayon-api-go/client"
	"github.com/bjerkio/crayon-api-go/client/customer_token"
	httptransport "github.com/go-openapi/runtime/client"
)

// GetTokenRequest are request details needed to get customer tokens
type GetTokenRequest struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

// GetToken returns CustomerTokens
func GetToken(r GetTokenRequest) (*customer_token.CreateCustomerTokenOK, error) {
	t := httptransport.BasicAuth(r.ClientID, r.ClientSecret)

	req := customer_token.NewCreateCustomerTokenParams()
	req.SetUsername(r.Username)
	req.SetPassword(r.Password)

	return apiclient.Default.CustomerToken.CreateCustomerToken(req, t)
}
