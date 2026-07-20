package main

import "fmt"

func main() {
	competitions := [][]string{
		{"AlgoMasters", "FrontPage Freebirds"},
		{"Runtime Terror", "Static Startup"},
		{"WeC#", "Hypertext Assassins"},
		{"AlgoMasters", "WeC#"},
		{"Static Startup", "Hypertext Assassins"},
		{"Runtime Terror", "FrontPage Freebirds"},
		{"AlgoMasters", "Runtime Terror"},
		{"Hypertext Assassins", "FrontPage Freebirds"},
		{"Static Startup", "WeC#"},
		{"AlgoMasters", "Static Startup"},
		{"FrontPage Freebirds", "WeC#"},
		{"Hypertext Assassins", "Runtime Terror"},
		{"AlgoMasters", "Hypertext Assassins"},
		{"WeC#", "Runtime Terror"},
		{"FrontPage Freebirds", "Static Startup"}}

	results := []int{1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0}
	fmt.Println(TournamentWinner(competitions, results))
}

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	m := make(map[string]int)

	for i, arr := range competitions {
		if results[i] == HOME_TEAM_WON {
			m[arr[0]] += 3
		} else {
			m[arr[1]] += 3
		}
	}

	var done string
	win := 0

	for key, val := range m {
		if val > win {
			done = key
		}
	}

	return done
}
