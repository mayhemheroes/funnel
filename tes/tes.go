package tes

import (
	"fmt"

	"github.com/getlantern/deepcopy"
)

func (lr *ListTasksRequest) GetTags() map[string]string {
	out := map[string]string{}
	for i := range lr.TagKey {
		v := ""
		if len(lr.TagValue) > i {
			v = lr.TagValue[i]
		}
		out[lr.TagKey[i]] = v
	}
	return out
}

func GetTaskView(s string) (View, error) {
	if x, ok := View_value[s]; ok {
		return View(x), nil
	}
	return View_MINIMAL, fmt.Errorf("Not found")
}

// GetBasicView returns the basic view of a task.
func (task *Task) GetBasicView() *Task {
	view := &Task{}
	deepcopy.Copy(view, task)

	// remove contents from inputs
	for _, v := range view.Inputs {
		v.Content = ""
	}

	// remove stdout and stderr from Task.Logs.Logs
	for _, tl := range view.Logs {
		tl.SystemLogs = nil
		for _, el := range tl.Logs {
			el.Stdout = ""
			el.Stderr = ""
		}
	}
	return view
}

// GetMinimalView returns the minimal view of a task.
func (task *Task) GetMinimalView() *Task {
	id := task.Id
	state := task.State
	return &Task{
		Id:    id,
		State: state,
	}
}
