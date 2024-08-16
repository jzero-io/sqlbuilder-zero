package system_dict_key

import (
	"context"
	"github.com/jzero-io/jzero-contrib/condition"
	"testing"
)

func TestBulkUpdate(t *testing.T) {
	model := NewTSystemDictKeyModel(nil)
	err := model.UpdateFieldsByCondition(context.Background(), map[string]any{
		"field1": "value1",
		"field2": "value2",
	}, condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    []int{1, 2, 3},
	})
	if err != nil {
		t.Fatal(err)
	}
}
