package account

import (
	"net/http"

	"golang.org/x/net/context"

	"net/url"

	"github.com/almighty/almighty-core/account/tenant"
	"github.com/almighty/almighty-core/goasupport"
	goaclient "github.com/goadesign/goa/client"
)

type tenantConfig interface {
	GetTenantServiceURL() string
}

// NewInitTenant creates a new tenant service in oso
func NewInitTenant(config tenantConfig) func(context.Context) error {
	return func(ctx context.Context) error {
		return InitTenant(ctx, config)
	}
}

// InitTenant creates a new tenant service in oso
func InitTenant(ctx context.Context, config tenantConfig) error {

	u, err := url.Parse(config.GetTenantServiceURL())
	if err != nil {
		return err
	}

	c := tenant.New(goaclient.HTTPClientDoer(http.DefaultClient))
	c.Host = u.Host
	c.Scheme = u.Scheme
	c.SetJWTSigner(goasupport.NewForwardSigner(ctx))

	// Ignore response for now
	_, err = c.SetupTenant(ctx, tenant.SetupTenantPath())
	if err != nil {
		return err
	}
	return nil
}
