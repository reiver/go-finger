package classicfinger

import (
	"github.com/reiver/go-finger"

	osuser "os/user"
)

// LooupOSUser looks up a user on the operating-system (OS) via their username,
// and returns their real-name, and the path to their home-directory.
func LookupOSUser(username string) (OSUser, error) {

	var user *osuser.User
	{
		user, err = osuser.Lookup(username)
		if nil != err {

			// Note that we do not pass on the actual error here.
			// We do this to not leak too much information to the client about what the actual problem is.
			// Just in case this is part of an attempt to "crack" the server.

			switch err.(type) {
			case osuser.UnknownUserError:
				err = finger.ErrServerFailed
				return
			default:
				err = finger.ErrServerErred
				return
			}
		}
		if nil == user {
			err = finger.ErrServerErred
			return
		}
	}

	{
		homedirpath = user.HomeDir
		if "" == homedirpath {
			err = finger.ErrServerFailed
			return
		}
	}

	{
		realname = user.Name
	}

	return realname, homedirpath, nil
}
