package gordle

import "testing"

func TestFeedbackString(t *testing.T) {

	testCases := map[string]struct {
		fb   feedback
		want string
	}{
		"three correct": {
			fb:   feedback{correctPosition, correctPosition, correctPosition},
			want: "🟩🟩🟩",
		},
		"one of each": {
			fb:   feedback{correctPosition, wrongPosition, absentCharacter},
			want: "🟩🟨⬜️",
		},
		"one for each (different order)": {
			fb:   feedback{wrongPosition, absentCharacter, correctPosition},
			want: "🟨⬜️🟩",
		},
		"unknown": {
			fb:   feedback{4},
			want: "🟥",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.fb.String()
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
