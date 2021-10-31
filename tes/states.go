package tes

import "fmt"


// TransitionError describes an invalid state transition.
type TransitionError struct {
	From, To State
}

func (te *TransitionError) Error() string {
	return fmt.Sprintf("invalid state transition from %s to %s",
		te.From, te.To)
}

// ValidateTransition validates a task state transition.
// Returns a TransitionError if the transition is not valid.
func ValidateTransition(from, to State) error {

	if from == to {
		return nil
	}

	if from == State_PAUSED || to == State_PAUSED {
		return fmt.Errorf("paused state is not implemented")
	}

	switch from {
	case State_UNKNOWN:
		// May transition from Unknown to anything
		return nil

	case State_QUEUED:
		// May transition from Queued to anything except Unknown
		if to == State_UNKNOWN {
			return &TransitionError{from, to}
		}
		return nil

	case State_INITIALIZING:

		switch to {
		case State_UNKNOWN, State_QUEUED:
			return &TransitionError{from, to}
		case State_RUNNING, State_EXECUTOR_ERROR, State_SYSTEM_ERROR, State_CANCELED:
			return nil
		}

	case State_RUNNING:

		switch to {
		case State_UNKNOWN, State_QUEUED:
			return &TransitionError{from, to}
		case State_COMPLETE, State_EXECUTOR_ERROR, State_SYSTEM_ERROR, State_CANCELED:
			return nil
		}

	case State_EXECUTOR_ERROR, State_SYSTEM_ERROR, State_CANCELED, State_COMPLETE:
		// May not transition out of terminal state.
		return &TransitionError{from, to}

	default:
		return &TransitionError{from, to}
	}
	// Shouldn't be reaching this point, but just in case.
	return &TransitionError{from, to}
}
