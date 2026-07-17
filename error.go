package flags

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gildas/go-errors"
)

var (
	InvalidEnumValue = errors.NewSentinel(http.StatusBadRequest, "error.value.invalid", "Flag value \"%s\" is invalid. Expected values are %s")
)

// maxAllowedInError caps how many allowed values an InvalidEnumValue error
// lists. Enum flags backed by an AllowedFunc can resolve to large sets (e.g.
// every workspace member), and dumping all of them buries the real error.
const maxAllowedInError = 15

// formatAllowed renders the allowed-value list for an InvalidEnumValue error,
// truncating oversized sets to the first maxAllowedInError entries followed by
// a count of the remainder.
func formatAllowed(allowed []string) string {
	if len(allowed) <= maxAllowedInError {
		return strings.Join(allowed, ", ")
	}
	return fmt.Sprintf(
		"%s, ... (%d more; use shell completion to list all)",
		strings.Join(allowed[:maxAllowedInError], ", "),
		len(allowed)-maxAllowedInError,
	)
}
