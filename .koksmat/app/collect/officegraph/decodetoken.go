package officegraph

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// TokenClaims represents the claims in the JWT token
type TokenClaims struct {
	Aud               string   `json:"aud"`
	Iss               string   `json:"iss"`
	Iat               int      `json:"iat"`
	Nbf               int      `json:"nbf"`
	Exp               int      `json:"exp"`
	Aio               string   `json:"aio"`
	AppDisplayname    string   `json:"app_displayname"`
	Appid             string   `json:"appid"`
	Appidacr          string   `json:"appidacr"`
	Idp               string   `json:"idp"`
	Idtyp             string   `json:"idtyp"`
	Oid               string   `json:"oid"`
	Rh                string   `json:"rh"`
	Roles             []string `json:"roles"`
	Sub               string   `json:"sub"`
	TenantRegionScope string   `json:"tenant_region_scope"`
	Tid               string   `json:"tid"`
	Uti               string   `json:"uti"`
	Ver               string   `json:"ver"`
	Wids              []string `json:"wids"`
	XmsTcdt           int      `json:"xms_tcdt"`
	XmsTdbr           string   `json:"xms_tdbr"`
}

// DecodeClaim decodes a JWT token and returns the claims
func DecodeClaim(accessToken string) (*TokenClaims, error) {
	parts := strings.Split(accessToken, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid token")
	}

	decoded, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode token: %v", err)
	}

	var claims TokenClaims
	err = json.Unmarshal(decoded, &claims)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal token claims: %v", err)
	}

	return &claims, nil
}
