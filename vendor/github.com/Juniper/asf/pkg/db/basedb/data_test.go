package basedb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseData_RowCount(t *testing.T) {
	tests := []struct {
		name      string
		d         DatabaseData
		wantCount int
	}{{
		name: "nil",
	}, {
		name: "empty tables",
		d:    DatabaseData{"a": {}, "b": {}},
	}, {
		name:      "some rows in tables",
		d:         DatabaseData{"a": {{}, {}, {}}, "b": {{}, {}}},
		wantCount: 5,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := tt.d.RowCount(); gotCount != tt.wantCount {
				t.Errorf("DatabaseData.RowsCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestDatabaseData_ForEachRow(t *testing.T) {
	type schemaIDRow struct {
		schemaID string
		row      RowData
	}

	tests := []struct {
		name       string
		d          DatabaseData
		calledWith []schemaIDRow
		wantErr    bool
	}{{
		name: "nil",
	}, {
		name: "some database data",
		d:    DatabaseData{"a": {{"id": "x"}}, "b": {{"id": "y"}, {"id": "z"}}},
		calledWith: []schemaIDRow{
			{"a", RowData{"id": "x"}},
			{"b", RowData{"id": "y"}},
			{"b", RowData{"id": "z"}},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []schemaIDRow
			if err := tt.d.ForEachRow(func(schemaID string, row RowData) error {
				result = append(result, schemaIDRow{schemaID, row})
				return nil
			}); (err != nil) != tt.wantErr {
				t.Errorf("DatabaseData.ForEachRow() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.calledWith, result)
		})
	}
}
