package main

import (
	"reflect"
	"testing"
)

func TestUsersIsQuitting(t *testing.T) {
	testCases := []struct {
		Description string
		User        User
		Want        bool
	}{
		{
			Description: "A user without last day in the DisplayName should not be quitting",
			User: User{
				DisplayName: "Troy",
			},
			Want: false,
		},
		{
			Description: "A user with last day at the end of DisplayName should be quitting",
			User: User{
				DisplayName: "Troy (Last Day 10/10)",
			},
			Want: true,
		},
		{
			Description: "A user with last day in the middle of DisplayName should be quitting",
			User: User{
				DisplayName: "Troy (Last Day 10/10) Barnes",
			},
			Want: true,
		},
		{
			Description: "A user with last day lowercased in the middle of DisplayName should be quitting",
			User: User{
				DisplayName: "Troy (last day 10/10) Barnes",
			},
			Want: true,
		},
		{
			Description: "A user with quitting date, then last day in the end of DisplayName should be quitting",
			User: User{
				DisplayName: "Troy (10/10 last day)",
			},
			Want: true,
		},
		{
			Description: "A user with quitting date, then last day in the beginning of DisplayName should be quitting",
			User: User{
				DisplayName: "(10/10 last day) Troy",
			},
			Want: true,
		},
		{
			Description: "A user with a similar name should not be quitting",
			User: User{
				DisplayName: "Lasting Dayum",
			},
			Want: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			got := test.User.IsQuitting()

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}

func TestUsersQuitDate(t *testing.T) {
	testCases := []struct {
		Description string
		User        User
		Want        string
	}{
		{
			Description: "A user without last day in the DisplayName",
			User: User{
				DisplayName: "Troy",
			},
			Want: "",
		},
		{
			Description: "A user with last day at the end of DisplayName",
			User: User{
				DisplayName: "Troy (Last Day 10/10)",
			},
			Want: "10/10",
		},
		{
			Description: "A user with last day in the middle of DisplayName",
			User: User{
				DisplayName: "Troy (Last Day 10/10) Barnes",
			},
			Want: "10/10",
		},
		{
			Description: "A user with last day lowercased in the middle of DisplayName",
			User: User{
				DisplayName: "Troy (last day 10/10) Barnes",
			},
			Want: "10/10",
		},
		{
			Description: "A user with quitting date, then last day in the end of DisplayName",
			User: User{
				DisplayName: "Troy (10/10 last day)",
			},
			Want: "10/10",
		},
		{
			Description: "A user with quitting date, then last day in the beginning of DisplayName",
			User: User{
				DisplayName: "(10/10/2021 last day) Troy",
			},
			Want: "10/10",
		},
		{
			Description: "A user with a similar name",
			User: User{
				DisplayName: "Lasting Dayum",
			},
			Want: "",
		},
	}

	for _, test := range testCases {
		t.Run(test.Description, func(t *testing.T) {
			got := test.User.QuitDate()

			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
