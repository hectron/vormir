package main

import (
	"fmt"
	"regexp"
	"strings"
)

type User struct {
	ID          string
	Name        string
	DisplayName string
	Deleted     bool
}

var (
	lastDayRegex = regexp.MustCompile(`(?i)last\s*day`)
)

func fetchSlackUsers() ([]User, error) {
	slackUsers, err := slackClient.GetUsers()

	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve users from Slack. %v", err)
	}

	users := []User{}

	for _, slackUser := range slackUsers {
		displayName := slackUser.Profile.DisplayName

		if len(strings.TrimSpace(displayName)) == 0 {
			displayName = slackUser.Name
		}

		users = append(users, User{
			ID:          slackUser.ID,
			Name:        slackUser.Name,
			DisplayName: displayName,
			Deleted:     slackUser.Deleted,
		})
	}

	return users, nil
}

func (u User) IsQuitting() bool {
	return lastDayRegex.MatchString(u.DisplayName)
}
