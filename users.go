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
	lastDayDisplayNameRegex = regexp.MustCompile(`(?i)\(.*last\s*day.*\)`)
	lastDayRegex            = regexp.MustCompile(`\d{2}\/\d{2}`)
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

func findUsersThatAreQuitting() []User {
	users, err := fetchSlackUsers()

	if err != nil {
		return []User{}
	}

	matches := []User{}

	for _, user := range users {
		if user.IsQuitting() {
			matches = append(matches, user)
		}
	}

	return matches
}

func (u User) IsQuitting() bool {
	return lastDayDisplayNameRegex.MatchString(u.DisplayName)
}

func (u User) QuitDate() string {
	match := lastDayRegex.FindStringSubmatch(u.DisplayName)

	if len(match) > 0 {
		lastDayString := match[0]

		dateMatch := lastDayRegex.FindStringSubmatch(lastDayString)

		if len(dateMatch) > 0 {
			return dateMatch[0]
		}
	}

	return ""
}
