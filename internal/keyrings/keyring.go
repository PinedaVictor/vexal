package keyrings

import (
	"fmt"

	zKeyring "github.com/zalando/go-keyring"
)

const service = "vexal"

// SetSecret — stores a secret in the OS keychain under vexal/<context>/<field>
func SetSecret(contextName, field, value string) error {
	return zKeyring.Set(service, fmt.Sprintf("%s/%s", contextName, field), value)
}

// GetSecret — retrieves a secret from the keychain by context and field name
func GetSecret(contextName, field string) (string, error) {
	return zKeyring.Get(service, fmt.Sprintf("%s/%s", contextName, field))
}

// DeleteSecret — removes a secret from the keychain (used when deleting a context)
func DeleteSecret(contextName, field string) error {
	return zKeyring.Delete(service, fmt.Sprintf("%s/%s", contextName, field))
}
