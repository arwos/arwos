package enums

const (
	ConnStatusAllow    = "allow"
	ConnStatusDisallow = "disallow"
)

var connStatus = map[string]struct{}{
	ConnStatusAllow:    {},
	ConnStatusDisallow: {},
}

func IsValidConnStatus(v string) bool {
	_, ok := connStatus[v]
	return ok
}
