package asana

import (
	"testing"

	helper "github.com/davidji99/terraform-provider-asana/helper/test"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testAccConfig *helper.TestConfig

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"asana": testAccProvider,
	}
	testAccConfig = helper.NewTestConfig()
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	testAccConfig.GetOrAbort(t, helper.TestConfigAsanaAccessToken)
}

func testAccProviderFactories(providers *[]*schema.Provider) map[string]func() (*schema.Provider, error) {
	return testAccProviderFactoriesInit(providers, []string{
		"asana",
	})
}

// testAccProviderFactoriesInit creates ProviderFactories for the provider under testing.
func testAccProviderFactoriesInit(providers *[]*schema.Provider, providerNames []string) map[string]func() (*schema.Provider, error) {
	var factories = make(map[string]func() (*schema.Provider, error), len(providerNames))

	for _, name := range providerNames {
		p := Provider()

		factories[name] = func() (*schema.Provider, error) { //nolint:unparam
			return p, nil
		}

		if providers != nil {
			*providers = append(*providers, p)
		}
	}

	return factories
}
