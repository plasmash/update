//go:build windows

package update

func hasWritePermissions(path string) error {
	//@todo handle windows update

	return errNoWritePermission
}
