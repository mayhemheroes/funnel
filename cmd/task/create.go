package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ohsu-comp-bio/funnel/tes"
	"golang.org/x/net/context"
)

// Create runs the "task create" CLI command, connecting to the server,
// calling CreateTask, and writing output to the given writer.
// Tasks are loaded from the "files" arg. "files" are file paths to JSON objects.
func Create(server string, files []string, reader io.Reader, writer io.Writer) error {
	cli, err := tes.NewClient(server)
	if err != nil {
		return err
	}

	for _, taskFile := range files {
		f, err := os.Open(taskFile)
		if err != nil {
			return err
		}

		var task tes.Task
		dec := json.NewDecoder(reader)
		err = dec.Decode(&task)

		r, err := cli.CreateTask(context.Background(), &task)
		if err != nil {
			return err
		}
		fmt.Fprintln(writer, r.Id)
		f.Close()
	}
	return nil
}
