package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"n2x.dev/x-api-go/grpc/resources/iam/auth"
	"n2x.dev/x-cli/pkg/vars"
	"n2x.dev/x-lib/pkg/errors"
	"n2x.dev/x-lib/pkg/utils"
)

type Credentials struct {
	AccountID    string `json:"accountID,omitempty"`
	LocationID   string `json:"locationID,omitempty"`
	FederationID string `json:"federationID,omitempty"`
	Key          string `json:"key,omitempty"`
	ExpiresAt    int64  `json:"expiresAt,omitempty"`
	UserID       string `json:"userID,omitempty"`
}

// GetNoAuthKey gets void authKey
func GetNoAuthKey() *auth.AuthKey {
	return &auth.AuthKey{
		Key: "no-auth",
	}
}

// GetAuthKey gets the authorization bearer string key
func GetAuthKey() (*auth.AuthKey, error) {
	cred, err := getCredentials()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	return &auth.AuthKey{Key: cred.Key}, nil
}

func GetAuthExpiresAt() (*time.Time, error) {
	cred, err := getCredentials()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	tm := time.UnixMilli(cred.ExpiresAt)

	return &tm, nil
}

func GetUserID() (string, error) {
	cred, err := getCredentials()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	if len(cred.UserID) == 0 {
		return "", fmt.Errorf("missing userID")
	}

	return cred.UserID, nil
}

func GetAccountID() (string, error) {
	if len(vars.AccountID) > 0 {
		return vars.AccountID, nil
	}

	cred, err := getCredentials()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	if len(cred.AccountID) == 0 {
		return "", fmt.Errorf("missing accountID")
	}

	return cred.AccountID, nil
}

/*
func getLocationID() (string, error) {
	if len(vars.LocationID) > 0 {
		return vars.LocationID, nil
	}

	cred, err := getCredentials()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	if len(cred.LocationID) == 0 {
		return "", fmt.Errorf("missing locationID")
	}

	return cred.LocationID, nil
}
*/

func getFederationID() (string, error) {
	if len(vars.FederationID) > 0 {
		return vars.FederationID, nil
	}

	cred, err := getCredentials()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getCredentials()", errors.Trace())
	}

	if len(cred.FederationID) == 0 {
		return "", fmt.Errorf("missing federationID")
	}

	return cred.FederationID, nil
}

func GetControllerEndpoint() (string, error) {
	federationID, err := getFederationID()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getFederationID()", errors.Trace())
	}

	apiserver := viper.GetString("apiserver.endpoint")
	if len(apiserver) == 0 {
		return "", fmt.Errorf("invalid apiserver endpoint")
	}

	return fmt.Sprintf("%s.%s", federationID, apiserver), nil
}

func GetAPIKeyFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function os.UserHomeDir()", errors.Trace())
	}

	n2xDir := filepath.Join(homeDir, ".n2x")

	if err := os.MkdirAll(n2xDir, 0700); err != nil {
		return "", errors.Wrapf(err, "[%v] function os.MkdirAll()", errors.Trace())
	}

	return filepath.Join(n2xDir, "apikey"), nil
}

func getCredentials() (*Credentials, error) {
	apiKeyFile, err := GetAPIKeyFile()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function getAPIKeyFile()", errors.Trace())
	}

	jsonBlob, err := utils.ReadJsonFile(apiKeyFile)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function utils.ReadJsonFile()", errors.Trace())
	}

	var cred Credentials

	if err := json.Unmarshal(jsonBlob, &cred); err != nil {
		return nil, errors.Wrapf(err, "[%v] function json.Unmarshal()", errors.Trace())
	}

	return &cred, nil
}
