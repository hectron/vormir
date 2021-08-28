package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type User struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	DisplayName string `db:"display_name"`
	Deleted     bool   `db:"deleted"`
}

var (
	lastDayDisplayNameRegex = regexp.MustCompile(`(?i)\(.*last\s*day.*\)`)
	lastDayRegex            = regexp.MustCompile(`\d{2}\/\d{2}`)
)

func updateOrCreateUsers() error {
	slackUsers, err := fetchSlackUsers()

	if err != nil {
		return err
	}

	usersInDb, err := fetchUsersFromDb()

	if err != nil {
		return err
	}

	lookup := make(map[string]User)

	for _, user := range usersInDb {
		lookup[user.ID] = user
	}

	fmt.Printf("Found %v users in Slack\n", len(slackUsers))

	for _, slackUser := range slackUsers {
		if _, ok := lookup[slackUser.ID]; ok {
			err := updateUser(&slackUser)

			if err != nil {
				return errors.Wrapf(err, "updating user %#v", slackUser)
			}
		} else {
			err := createUser(&slackUser)

			if err != nil {
				return errors.Wrapf(err, "creating user %#v", slackUser)
			}
		}
	}

	return nil
}

func createUser(user *User) error {
	_, err := db.NamedExec(`
		INSERT INTO users
			(id, name, display_name, deleted)
		VALUES
			(:id, :name, :display_name, :deleted)
	`, user)

	return errors.Wrapf(err, "inserting user %#v", user)
}

func updateUser(user *User) error {
	_, err := db.NamedExec(`
		UPDATE users
		SET
			name = :name,
			display_name = :display_name,
			deleted = :deleted
		WHERE
			id = :id
	`, user)

	return errors.Wrapf(err, "updating user %#v", user)
}

func fetchUsersFromDb() ([]User, error) {
	users := []User{}

	err := db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, errors.Wrap(err, "Selecting all users")
	}

	return users, nil
}

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
	users := []User{}

	err := db.Select(&users, "SELECT * FROM users WHERE display_name ilike '%last day%' AND NOT deleted")

	if err != nil {
		return []User{}
	}

	return users
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
