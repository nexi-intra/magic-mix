package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

// GetMicrosoftPublicKeys fetches the public keys from Microsoft
func GetMicrosoftPublicKeys() (jwk.Set, error) {
	resp, err := http.Get("https://login.microsoftonline.com/common/discovery/v2.0/keys")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch keys: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch keys: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	jwks, err := jwk.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWKS: %v", err)
	}

	return jwks, nil
}

// DecodeAndValidateMicrosoftJWT decodes and validates a JWT token from Microsoft
func DecodeAndValidateMicrosoftJWT(tokenString string) (map[string]interface{}, error) {
	jwks, err := GetMicrosoftPublicKeys()
	if err != nil {
		return nil, fmt.Errorf("failed to get Microsoft public keys: %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid in token header")
		}

		key, _ := jwks.LookupKeyID(kid)
		if key == nil {
			return nil, fmt.Errorf("unable to find key with kid: %v", kid)
		}

		var raw interface{}
		if err := key.Raw(&raw); err != nil {
			return nil, fmt.Errorf("unable to parse key: %v", err)
		}
		return raw, nil
	})

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse token: %v", err)
	// }

	if claims, ok := token.Claims.(jwt.MapClaims); ok { //&& token.Valid {
		// Check for expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token is expired")
			}
		}

		// Validate the issuer and audience if necessary
		// if iss, ok := claims["iss"].(string); ok {
		// 	if iss != "https://login.microsoftonline.com/{tenantid}/v2.0" {
		// 		return nil, fmt.Errorf("unexpected issuer: %v", iss)
		// 	}
		// }

		// if aud, ok := claims["aud"].(string); ok {
		// 	if aud != "your-audience-id" {
		// 		return nil, fmt.Errorf("unexpected audience: %v", aud)
		// 	}
		// }

		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

// func main() {
// 	// Example token (Replace with a real token for testing)
// 	tokenString := "your.jwt.token.here"

// 	decoded, err := DecodeAndValidateMicrosoftJWT(tokenString)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Decoded token:", decoded)
// 	}
// }
