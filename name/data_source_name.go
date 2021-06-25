package name

import (
	"context"
	"embed"
	"fmt"
	"hash/fnv"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//go:embed adjectives.txt nouns.txt
var embeddedFs embed.FS
var adjectives = NewWordGenerator(embeddedFs, "adjectives.txt")
var nouns = NewWordGenerator(embeddedFs, "nouns.txt")

func dataSourceName() *schema.Resource {
	return &schema.Resource{
		CreateContext: dataSourceNameRead,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,

		Schema: map[string]*schema.Schema{

			"seed": {
				Description: "The optional seed to use",
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"id": {
				Description: "The random name",
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func dataSourceNameRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if data.HasChange("seed") {
		seed := data.Get("seed").(string)
		rand.Seed(int64(hash(seed)))
	} else {
		rand.Seed(time.Now().Unix())
	}

	data.SetId(fmt.Sprintf("%s-%s", adjectives.Choose(), nouns.Choose()))

	return diags
}
