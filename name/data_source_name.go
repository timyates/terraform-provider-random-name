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
		ReadContext: dataSourceNameRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
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

	if err := data.Set("name", fmt.Sprintf("%s-%s", adjectives.Choose(), nouns.Choose())); err != nil {
		return diag.FromErr(err)
	}

	return diags
}