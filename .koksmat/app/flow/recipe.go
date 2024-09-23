package flow

type FlowDefinition struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stage []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Raci struct {
			Informed    string `json:"informed"`
			Consulted   string `json:"consulted"`
			Responsible string `json:"responsible"`
		} `json:"raci,omitempty"`
		Actions []struct {
			ID          string `json:"id"`
			UI          string `json:"ui"`
			Icon        string `json:"icon"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Interactive bool   `json:"interactive"`
			Transaction string `json:"transaction,omitempty"`
			Sql         string `json:"sql,omitempty"`
		} `json:"actions"`
		Triggers struct {
			Actions []struct {
				ID          string `json:"id"`
				Sql         string `json:"sql"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"actions"`
		} `json:"triggers,omitempty"`
		Decisions []struct {
			ID          string `json:"id"`
			Icon        string `json:"icon"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"decisions,omitempty"`
	} `json:"stage"`
	Actors struct {
		App struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"app"`
		Shop struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"shop"`
		User struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"user"`
		System struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"system"`
		Auditlog struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"auditlog"`
		Restaurant struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"restaurant"`
		DeliveryPerson struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"delivery_person"`
	} `json:"actors"`
	Purpose     string `json:"purpose"`
	Version     int    `json:"version"`
	Description string `json:"description"`
}
type RecipeV1 struct {
	// ID is the unique identifier of the recipe
	Version    int            `json:"version"`
	Definition FlowDefinition `json:"definition"`
}
