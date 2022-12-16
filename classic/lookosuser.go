package classicfinger

import (
	osuser "os/user"
)

// LooupOSUser looks up a user on the operating-system (OS) via their username,
// and returns their real-name, and the path to their home-directory.
func LookupOSUser(username string) (realname string, homedirpath string, err error) {

	var user *osuser.User
	{
		user, err = osuser.Lookup(username)
		if nil != err {

			// Note that we do not pass on the actual error here.
			// We do this to not leak too much information to the client about what the actual problem is.
			// Just in case this is part of an attempt to "crack" the server.

			switch err.(type) {
			case osuser.UnknownUserError:
				err = errFailed
				return
			default:
				err = errErred
				return
			}
		}
		if nil == user {
			err = errErred
			return
		}
	}

	{
		homedirpath = user.HomeDir
		if "" == homedirpath {
			err = errFailed
			return
		}
	}

	{
		realname = user.Name
	}

	return realname, homedirpath, nil
}
