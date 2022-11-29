package types

// C4cPageResponse is the json response from the endpoint
// it's shortened, as we dont need all the informations
type C4cPageResponse struct {
	Result struct {
		Data struct {
			Reports struct {
				Edges []struct {
					Node struct {
						Frontmatter struct {
							Slug     string `json:"slug"`
							Findings string `json:"findings"`
						} `json:"frontmatter"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"reports"`
		} `json:"data"`
	} `json:"result"`
}

// C4cReports represents one code4rena report
type C4cReports struct {
	Slug     string `json:"slug"`
	Findings string `json:"findings"`
	Site     string `json:"site"`
}
