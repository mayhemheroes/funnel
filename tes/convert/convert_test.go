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
		Inputs: []openapi.TesInput{
			{
				Type: openapi.FILE,
				Path: "/data/file",
			}, {
				Type: openapi.DIRECTORY,
				Path: "/data/results",
			},
		},
		Executors: []openapi.TesExecutor{
			{
				Command: []string{"echo", "hello"},
				Image:   "ubuntu",
			}, {
				Command: []string{"md5sum", "/data/file"},
				Image:   "alpine",
			},
		},
		State: openapi.COMPLETE,
		Tags: map[string]string{
			"job": "12345",
		},
	}

	OpenApi2Proto(o, &p)

	if o.Id != p.Id {
		t.Errorf("id %s != %s", o.Id, p.Id)
	}
	if o.Resources.CpuCores != p.Resources.CpuCores {
		t.Errorf("cpuCores %d != %d", o.Resources.CpuCores, p.Resources.CpuCores)
	}
	if p.State != tes.State_COMPLETE {
		t.Errorf("Incorrect state: %s", p.State)
	}
	if len(p.Executors) != 2 {
		t.Errorf("Incorrect executor count: %d", len(p.Executors))
	}
	if x, ok := p.Tags["job"]; ok {
		if x != "12345" {
			t.Errorf("tag incorrect %s != %s", "12345", x)
		}
	} else {
		t.Errorf("Tag key job not found")
	}

	if len(p.Inputs) != 2 {
		t.Errorf("Incorrect input count: %d", len(p.Inputs))
	}
}
