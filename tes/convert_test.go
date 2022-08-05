package tes

import (
	"testing"

	"github.com/ohsu-comp-bio/funnel/tes/openapi"
)

func TestOpenApi2ProtoConvert(t *testing.T) {
	p := Task{}

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
	if p.State != State_COMPLETE {
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

func TestProto2OpenAPIConvert(t *testing.T) {
	o := openapi.TesTask{}

	p := &Task{
		Id: "blaa",
		Resources: &Resources{
			CpuCores: 4,
		},
		Inputs: []*Input{
			{
				Type: FileType_FILE,
				Path: "/data/file",
			}, {
				Type: FileType_DIRECTORY,
				Path: "/data/results",
			},
		},
		Executors: []*Executor{
			{
				Command: []string{"echo", "hello"},
				Image:   "ubuntu",
			}, {
				Command: []string{"md5sum", "/data/file"},
				Image:   "alpine",
			},
		},
		State: State_COMPLETE,
		Tags: map[string]string{
			"job": "12345",
		},
	}

	Proto2OpenApi(p, &o)

	if o.Id != p.Id {
		t.Errorf("id %s != %s", o.Id, p.Id)
	}
	if o.Resources.CpuCores != p.Resources.CpuCores {
		t.Errorf("cpuCores %d != %d", o.Resources.CpuCores, p.Resources.CpuCores)
	}
	if o.State != openapi.COMPLETE {
		t.Errorf("Incorrect state: %s", p.State)
	}
	if len(o.Executors) != 2 {
		t.Errorf("Incorrect executor count: %d", len(o.Executors))
	}
	if x, ok := o.Tags["job"]; ok {
		if x != "12345" {
			t.Errorf("tag incorrect %s != %s", "12345", x)
		}
	} else {
		t.Errorf("Tag key job not found")
	}

	if len(o.Inputs) != 2 {
		t.Errorf("Incorrect input count: %d", len(o.Inputs))
	}
}
