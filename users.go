package main

import "fmt"

type User struct {
	ID          string
	Name        string
	DisplayName string
	Deleted     bool
}

func fetchSlackUsers() ([]User, error) {
	slackUsers, err := slackClient.GetUsers()

	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve users from Slack. %v", err)
	}

	users := []User{}

	for _, slackUser := range slackUsers {
		users = append(users, User{
			ID:          slackUser.ID,
			Name:        slackUser.Name,
			DisplayName: slackUser.Profile.DisplayName,
			Deleted:     slackUser.Deleted,
		})
	}

	return users, nil
}
