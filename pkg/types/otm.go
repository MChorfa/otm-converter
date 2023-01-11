package types

type Otm struct {
	OtmVersion      string           `json:"otmVersion"`
	Project         Project          `json:"project"`
	Representations []Representation `json:"representations"`
	Assets          []Asset          `json:"assets"`
	Components      []Component      `json:"components"`
	Dataflows       []Dataflow       `json:"dataflows"`
	TrustZones      []TrustZone      `json:"trustZones"`
	Threats         []Threat         `json:"threats"`
	Mitigations     []Asset          `json:"mitigations"`
}

type Asset struct {
	Name          string      `json:"name"`
	ID            string      `json:"id"`
	Description   string      `json:"description"`
	Risk          *AssetRisk  `json:"risk,omitempty"`
	Attributes    interface{} `json:"attributes"`
	RiskReduction *int64      `json:"riskReduction,omitempty"`
}

type AssetRisk struct {
	Confidentiality int64  `json:"confidentiality"`
	Integrity       int64  `json:"integrity"`
	Availability    int64  `json:"availability"`
	RiskComments    string `json:"riskComments"`
}

type Component struct {
	Name            string                    `json:"name"`
	ID              string                    `json:"id"`
	Description     string                    `json:"description"`
	Parent          Parent                    `json:"parent"`
	Type            string                    `json:"type"`
	Tags            []string                  `json:"tags,omitempty"`
	Representations []ComponentRepresentation `json:"representations"`
	Assets          *Assets                   `json:"assets,omitempty"`
	Threats         []ComponentThreat         `json:"threats,omitempty"`
	Attributes      interface{}               `json:"attributes"`
}

type Assets struct {
	Processed []string `json:"processed"`
	Stored    []string `json:"stored"`
}

type Parent struct {
	TrustZone string `json:"trustZone"`
}

type ComponentRepresentation struct {
	Representation string    `json:"representation"`
	ID             string    `json:"id"`
	Position       *Position `json:"position,omitempty"`
	Size           *Size     `json:"size,omitempty"`
	Package        *string   `json:"package,omitempty"`
	File           *string   `json:"file,omitempty"`
	Line           *int64    `json:"line,omitempty"`
	CodeSnippet    *string   `json:"codeSnippet,omitempty"`
}

type Position struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Size struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

type ComponentThreat struct {
	Threat      string       `json:"threat"`
	State       string       `json:"state"`
	Mitigations []Mitigation `json:"mitigations"`
}

type Mitigation struct {
	Mitigation string `json:"mitigation"`
	State      string `json:"state"`
}

type Dataflow struct {
	Name            string            `json:"name"`
	ID              string            `json:"id"`
	Bidirectional   bool              `json:"bidirectional"`
	Source          string            `json:"source"`
	Destination     string            `json:"destination"`
	Tags            []string          `json:"tags"`
	Assets          []string          `json:"assets"`
	Representations interface{}       `json:"representations"`
	Threats         []ComponentThreat `json:"threats"`
	Attributes      interface{}       `json:"attributes"`
}

type Project struct {
	Name         string     `json:"name"`
	ID           string     `json:"id"`
	Description  string     `json:"description"`
	Owner        string     `json:"owner"`
	OwnerContact string     `json:"ownerContact"`
	Attributes   Attributes `json:"attributes"`
}

type Attributes struct {
	CMDBID string `json:"cmdbId"`
}

type Representation struct {
	Name       string      `json:"name"`
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	Size       *Size       `json:"size,omitempty"`
	Attributes interface{} `json:"attributes"`
	Repository *Repository `json:"repository,omitempty"`
}

type Repository struct {
	URL string `json:"url"`
}

type Threat struct {
	Name        string      `json:"name"`
	ID          string      `json:"id"`
	Description string      `json:"description"`
	Categories  []string    `json:"categories"`
	Cwes        []string    `json:"cwes"`
	Risk        ThreatRisk  `json:"risk"`
	Attributes  interface{} `json:"attributes"`
	Tags        []string    `json:"tags"`
}

type ThreatRisk struct {
	Likelihood        int64  `json:"likelihood"`
	LikelihoodComment string `json:"likelihoodComment"`
	Impact            int64  `json:"impact"`
	ImpactComment     string `json:"impactComment"`
}

type TrustZone struct {
	Name            string                    `json:"name"`
	ID              string                    `json:"id"`
	Type            string                    `json:"type"`
	Description     string                    `json:"description"`
	Risk            TrustZoneRisk             `json:"risk"`
	Representations []TrustZoneRepresentation `json:"representations"`
	Attributes      interface{}               `json:"attributes"`
	Parent          *Parent                   `json:"parent,omitempty"`
}

type TrustZoneRepresentation struct {
	Representation string   `json:"representation"`
	ID             string   `json:"id"`
	Position       Position `json:"position"`
	Size           Size     `json:"size"`
}

type TrustZoneRisk struct {
	TrustRating int64 `json:"trustRating"`
}
