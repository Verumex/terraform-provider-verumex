package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func verumexWarehouseMetadata() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Warehouse metadata for the current environment.",

		ReadContext: dataSourceVerumexWarehouseMetadataRead,

		Schema: map[string]*schema.Schema{
			"warehouse_id": {
				// This description is used by the documentation generator and the language server.
				Description: "The VmX Core identifier of the warehouse.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceVerumexWarehouseMetadataRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "warehouse-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}
