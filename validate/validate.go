package validate

import "net/url"

// URLAddress validates if an address is a valid url
func URLAddress(address string) bool {
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return false
	}
	return true
}
