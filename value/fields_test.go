
/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package value

import (
	"testing"

	"github.com/go-json-experiment/json"
)

func TestFieldListMarshalUnmarshal(t *testing.T) {
	testCases := []struct {
		name    string
		fl      FieldList
		jsonStr string
	}{
		{
			name: "simple case",
			fl: FieldList{
				{Name: "a", Value: NewValueInterface("b")},
				{Name: "c", Value: NewValueInterface(int64(1))},
			},
			jsonStr: `{"a":"b","c":1}`,
		},
		{
			name:    "empty",
			fl:      FieldList{},
			jsonStr: `{}`,
		},
		{
			name: "nested",
			fl: FieldList{
				{Name: "a", Value: NewValueInterface(map[string]interface{}{
					"b": "c",
				})},
			},
			jsonStr: `{"a":{"b":"c"}}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test Marshal
			b, err := json.Marshal(tc.fl)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}
			if string(b) != tc.jsonStr {
				t.Errorf("expected marshaled json to be %v, got %v", tc.jsonStr, string(b))
			}

			// Test Unmarshal
			var fl FieldList
			err = json.Unmarshal([]byte(tc.jsonStr), &fl)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}
			if !fl.Equals(tc.fl) {
				t.Errorf("expected unmarshaled field list to be %#v, got %#v", tc.fl, fl)
			}
		})
	}
}
