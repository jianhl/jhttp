//go:build appengine

package jhttp

func init() {
	defaultPlatform = PlatformGoogleAppEngine
}
