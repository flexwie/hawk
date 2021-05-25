package exec

import (
	"time"

	"github.com/google/uuid"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"hawk.wie.gg/models"
)

func CreateEntry(name string, in string, out string) (models.Entry, error) {
	var start string
	var end string
	var entry models.Entry

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	//check if in is empty
	if in == "" {
		start = time.Now().Format(time.RFC3339)
	} else {
		r, err := w.Parse(in, time.Now())
		if err != nil {
			return models.Entry{}, err
		}
		if r == nil {
			start = in
		} else {
			start = r.Time.Local().Format(time.RFC3339)
		}
	}
	entry.Start = start

	//check if out is provided
	if out != "" {
		r, err := w.Parse(out, time.Now())
		if err != nil {
			return models.Entry{}, err
		}
		if r == nil {
			end = out
		} else {
			end = r.Time.Local().Format(time.RFC3339)
		}
	}
	entry.End = end

	entry.Name = name
	entry.Id = uuid.NewString()

	return entry, nil
}
