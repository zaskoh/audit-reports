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

// SherlockReports represents one sherlock report
type SherlockReports struct {
	Slug     string `json:"slug"`
	Findings string `json:"findings"`
	Site     string `json:"site"`
}

type SherlockListResult struct {
	EndsAt                    int         `json:"ends_at"`
	ID                        int         `json:"id"`
	JudgingEndsAt             int         `json:"judging_ends_at"`
	JudgingPrizePool          interface{} `json:"judging_prize_pool"`
	LeadSeniorAuditorFixedPay int         `json:"lead_senior_auditor_fixed_pay"`
	LeadSeniorAuditorHandle   string      `json:"lead_senior_auditor_handle"`
	LogoURL                   string      `json:"logo_url"`
	Private                   bool        `json:"private"`
	PrizePool                 int         `json:"prize_pool"`
	ShortDescription          string      `json:"short_description"`
	Sponsor                   string      `json:"sponsor"`
	StartsAt                  int         `json:"starts_at"`
	Status                    string      `json:"status"`
	TemplateRepoName          string      `json:"template_repo_name"`
	Title                     string      `json:"title"`
}
