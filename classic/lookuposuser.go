package classicfinger

import (
	"github.com/reiver/go-finger"

	go_osuser "os/user"
)

// LooupOSUser looks up a user on the operating-system (OS) via their username,
// and returns their real-name, and the path to their home-directory.
func LookupOSUser(username string) (osuser OSUser, err error) {

	{
		osuser.UserName = username
	}

	var user *go_osuser.User
	{
		user, err = go_osuser.Lookup(username)
		if nil != err {

			// Note that we do not pass on the actual error here.
			// We do this to not leak too much information to the client about what the actual problem is.
			// Just in case this is part of an attempt to "crack" the server.

			switch err.(type) {
			case go_osuser.UnknownUserError:
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
		osuser.RealName = user.Name
	}

	{
		osuser.HomeDirPath = user.HomeDir
		if "" == osuser.HomeDirPath {
			err = finger.ErrServerFailed
			return
		}
	}


	return osuser, nil
}
