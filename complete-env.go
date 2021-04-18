package gonsole

import (
	"os"
	"sort"
	"strings"

	"github.com/maxlandon/readline"
)

// CompleteEnvironmentVariables - Returns all environment variables as suggestions.
func (c *CommandCompleter) CompleteEnvironmentVariables() (completions []*readline.CompletionGroup) {

	grp := &readline.CompletionGroup{
		Name:         "console OS environment",
		MaxLength:    5, // Should be plenty enough
		DisplayType:  readline.TabDisplayGrid,
		TrimSlash:    true, // Some variables can be paths
		Descriptions: map[string]string{},
	}

	var clientEnv = map[string]string{}
	env := os.Environ()

	for _, kv := range env {
		key := strings.Split(kv, "=")[0]
		value := strings.Split(kv, "=")[1]
		clientEnv[key] = value
	}

	keys := make([]string, 0, len(clientEnv))
	for k := range clientEnv {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		grp.Suggestions = append(grp.Suggestions, key)
		value := clientEnv[key]
		grp.Descriptions[key] = value
	}

	// Add some special ones
	grp.Aliases = map[string]string{
		"~": "HOME",
	}

	completions = append(completions, grp)

	return
}
