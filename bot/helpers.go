// Different helper functions

package bot

// Function "contains" check whether slice a includes string x
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
