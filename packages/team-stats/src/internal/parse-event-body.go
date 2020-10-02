package internal

// TeamStatsLambdaEvent is the expected lambda event for the team stats lambda
type TeamStatsLambdaEvent struct {
	TeamAbrs  []string
	TeamNames []string
	PlayerIDs []string
}

// TeamStatsResponse is the expected response for the team stats lambda
type TeamStatsResponse struct {
	Response     TeamStatsViewModel
	ErrorMessage string `json:"errorMessage"`
}

// ParseLambdaEventBody should parse the body with team IDs to retrieve from dynamo
func ParseLambdaEventBody(
	body TeamStatsLambdaEvent,
	listAbrs [30]string,
	listNames [30]string,
) ([]string, string) {
	// Initialize response slice
	response := make([]string, 0)

	// If TeamAbrs is found, loop through and append to response
	if body.TeamAbrs != nil {
		for _, val := range body.TeamAbrs {
			if contains(listAbrs, val) {
				response = append(response, val)
			}
		}
	}

	// If TeamNames is found, loop through and append to response
	if body.TeamNames != nil {
		for _, val := range body.TeamNames {
			if contains(listAbrs, val) {
				response = append(response, val)
			}
		}
	}

	// If PlayerIDs is found, append all player ids to response
	if len(body.PlayerIDs) != 0 {
		for _, playerID := range body.PlayerIDs {
			response = append(response, playerID)
		}
	}

	// If response array is not empty, return response array
	if len(response) != 0 {
		return response, ""
	}

	// Otherwise return an error message
	errorMsg := "Error: no expected parameters found (teamAbrs, teamNames, playerID)"
	return nil, errorMsg
}

func contains(a [30]string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
