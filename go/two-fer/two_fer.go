// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

/*
ShareWith returns "One for X, one for me.",
where X is the given name
or if no name is given, the X stands for "you"
*/
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
