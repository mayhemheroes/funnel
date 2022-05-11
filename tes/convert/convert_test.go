package convert

import (
	"testing"

	"github.com/ohsu-comp-bio/funnel/tes"
	"github.com/ohsu-comp-bio/funnel/tes/openapi"
)

func TestTaskConvert(t *testing.T) {
	p := tes.Task{}

	o := openapi.TesTask{
		Id: "blaa",
		Resources: openapi.TesResources{
			CpuCores: 4,
		},
	}

	OpenApi2Proto(o, &p)

	if o.Id != p.Id {
		t.Errorf("id %s != %s", o.Id, p.Id)
	}
	if o.Resources.CpuCores != int64(p.Resources.CpuCores) {
		t.Errorf("id %d != %d", o.Resources.CpuCores, p.Resources.CpuCores)
	}

}
