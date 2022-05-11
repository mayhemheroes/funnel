package events

import (
	"context"

	"github.com/ohsu-comp-bio/funnel/tes"
)

// TaskBuilder aggregates events into an in-memory Task object.
type TaskBuilder struct {
	*tes.Task
	UnimplementedEventServiceServer
}

// WriteEvent updates the Task object.
func (tb TaskBuilder) WriteEvent(ctx context.Context, ev *Event) (*WriteEventResponse, error) {
	t := tb.Task
	t.Id = ev.Id
	attempt := int(ev.Attempt)
	index := int(ev.Index)

	switch ev.Type {
	case Type_TASK_STATE:
		to := ev.GetState()
		if err := tes.ValidateTransition(t.GetState(), ev.GetState()); err != nil {
			return nil, err
		}
		t.State = to

	case Type_SYSTEM_LOG:
		t.GetLogs()[attempt].SystemLogs = append(t.GetLogs()[attempt].SystemLogs, ev.SysLogString())

	case Type_TASK_START_TIME:
		t.GetLogs()[attempt].StartTime = ev.GetStartTime()

	case Type_TASK_END_TIME:
		t.GetLogs()[attempt].EndTime = ev.GetEndTime()

	case Type_TASK_OUTPUTS:
		t.GetLogs()[attempt].Outputs = ev.GetOutputs().Value

	case Type_TASK_METADATA:
		if t.GetLogs()[attempt].Metadata == nil {
			t.GetLogs()[attempt].Metadata = map[string]string{}
		}
		for k, v := range ev.GetMetadata().Value {
			t.GetLogs()[attempt].Metadata[k] = v
		}

	case Type_EXECUTOR_START_TIME:
		t.GetLogs()[attempt].Logs[index].StartTime = ev.GetStartTime()

	case Type_EXECUTOR_END_TIME:
		t.GetLogs()[attempt].Logs[index].EndTime = ev.GetEndTime()

	case Type_EXECUTOR_EXIT_CODE:
		t.GetLogs()[attempt].Logs[index].ExitCode = ev.GetExitCode()

	case Type_EXECUTOR_STDOUT:
		t.GetLogs()[attempt].Logs[index].Stdout = ev.GetStdout()

	case Type_EXECUTOR_STDERR:
		t.GetLogs()[attempt].Logs[index].Stderr = ev.GetStderr()
	}

	return nil, nil
}

func (tb TaskBuilder) Close() {}
