package expensify

import(
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"terraform-provider-expensify/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"partner_user_id": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("PARTNER_USER_ID", ""),
			},
			"partner_user_secret": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("PARTNER_USER_SECRET", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"expensify_user": resourceUser(),
			"expensify_policy": resourcePolicy(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"expensify_user": dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	partnerUserId := d.Get("partner_user_id").(string)
	partnerUserSecret := d.Get("partner_user_secret").(string)
	var diags diag.Diagnostics
	return client.NewClient(partnerUserId, partnerUserSecret), diags
}