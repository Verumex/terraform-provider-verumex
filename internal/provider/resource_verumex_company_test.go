package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceVerumexCompany(t *testing.T) {
	t.Skip("resource test not yet implemented, remove this once you implement")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVerumexCompany,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"verumex_company.foo", "name", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccResourceVerumexCompany = `
resource "verumex_company" "foo" {
  name = "bar"
}
`
