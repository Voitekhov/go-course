package sprint_2_answers

func LoginCheck(users map[string]bool, login string) bool {
	_, exists := users[login]
	return exists
}
