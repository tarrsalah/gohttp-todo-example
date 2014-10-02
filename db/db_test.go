package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPost(t *testing.T) {
	tasks := []struct {
		task  *Task
		print string
	}{
		{NewTask("todo 1"), "[id=0, todo=\"todo 1\", done=false]"},
		{NewTask("todo 2"), "[id=0, todo=\"todo 2\", done=false]"},
		{NewTask("todo 3"), "[id=0, todo=\"todo 3\", done=false]"},
		{NewTask("todo 4"), "[id=0, todo=\"todo 4\", done=false]"},
	}
	for _, k := range tasks {
		assert.Equal(t, k.task.Done, false)
		assert.Equal(t, k.task.String(), k.print)
	}
}

func TestDB(t *testing.T) {
	BootStrap()
	defer Map.Dbx.DB.Close()
	t.Skip()
}
