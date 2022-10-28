package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVerumexWarehouseMetadata(t *testing.T) {
	t.Skip("data source not yet implemented, remove this when implementing...")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVerumexWarehouseMetadata,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.verumex_warehouse_metadata.foo", "warehouse_id", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourceVerumexWarehouseMetadata = `
data "verumex_warehouse_metadata" "foo" {
  warehouse_id = "bar"
}
`
