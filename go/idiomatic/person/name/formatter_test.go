package name_test

import (
	"testing"

	"github.com/boundedinfinity/canonical_model/idiomatic/person/name"
	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	tcs := []struct {
		name      string
		input     name.Name
		formatter *name.Formatter
		expected  string
	}{
		{
			name: "Given1 Family1 -> Western",
			input: name.Name{
				Given:  []string{"Given1"},
				Family: []string{"Family1"},
			},
			formatter: name.WesternFormatter(),
			expected:  "Given1 Family1",
		},
		{
			name: "Given1 Middle1 Family1 -> Western",
			input: name.Name{
				Given:  []string{"Given1"},
				Middle: []string{"Middle1"},
				Family: []string{"Family1"},
			},
			formatter: name.WesternFormatter(),
			expected:  "Given1 Middle1 Family1",
		},
		{
			name: "Given1 Family1 -> Eastern",
			input: name.Name{
				Given:  []string{"Given1"},
				Family: []string{"Family1"},
			},
			formatter: name.EasternFormatter(),
			expected:  "Family1 Given1",
		},
		{
			name: "Given1 Middle1 Family1 -> Eastern",
			input: name.Name{
				Given:  []string{"Given1"},
				Middle: []string{"Middle1"},
				Family: []string{"Family1"},
			},
			formatter: name.EasternFormatter(),
			expected:  "Family1 Given1 Middle1",
		},
		{
			name: "Given1 Family1 -> Email1",
			input: name.Name{
				Given:  []string{"Given1"},
				Family: []string{"Family1"},
			},
			formatter: name.Email1("example.com", 10),
			expected:  "given1.family1@example.com",
		},
		{
			name: "Given1 Middle1 Family1 Family2 -> Email1",
			input: name.Name{
				Given:  []string{"Given1"},
				Middle: []string{"Middle1"},
				Family: []string{"Family1", "Family2"},
			},
			formatter: name.Email1("example.com", 10),
			expected:  "given1.family2@example.com",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.formatter.Format(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
