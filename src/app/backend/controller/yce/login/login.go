package login

import (
	"log"
	"fmt"
	"strconv"
	"github.com/kataras/iris"
	"app/backend/common/util/encrypt"
	myuser "app/backend/model/mysql/user"
	mysession "app/backend/common/util/session"
	myerror "app/backend/common/error"
)

type LoginController struct {
	*iris.Context
}

// Check username && password
func (lc *LoginController) check(name, password string) (*myuser.User, *myerror.YceError) {

	encryptPass := encrypt.NewEncryption(password).String()

	// YceError pointer
	var ye *myerror.YceError

	// New a User
	user := new(myuser.User)

	err := user.QueryUserByNameAndPassword(name, encryptPass)

	if err != nil {
		log.Printf("Can not find the user: username=%s, err=%s\n", name, err)
		ye = myerror.NewYceError(1001, err.Error())
		return nil, ye
	}

	return user, nil

}

// Store Session and Set cookie
func (lc *LoginController) session(user *myuser.User) *myerror.YceError {

	// Store (id,orgId) in SessionStore
	id := strconv.Itoa(int(user.Id))
	orgId := strconv.Itoa(int(user.OrgId))

	session := mysession.NewSession(id, user.Name, orgId)

	ss := mysession.SessionStoreInstance()

	err := ss.Set(session)

	fmt.Printf("Session: sessionId=%s, userId=%s, userName=%s, orgId=%s\n", session.SessionId, session.UserId, session.UserName, session.OrgId)

	if err != nil {
		log.Fatal("Set session error: sessionId=%s, err=%s\n", session.SessionId, err)
		ye := myerror.NewYceError(2001, err.Error())
		json, _ := ye.EncodeJson()
		lc.Write(json)
		return ye
	}

	// Set cookie
	lc.SetCookieKV("sessionId", session.SessionId)

	return nil

}

// POST /api/v1/users/{email}/login
func (lc LoginController) Post() {

	email := lc.Param("email")
	password := string(lc.FormValue("password"))

	user, ye := lc.check(email, password)
	if ye != nil {
		json, _ := ye.EncodeJson()
		lc.Write(json)
		return
	}

	ye = lc.session(user)
	if ye != nil {
		json, _ := ye.EncodeJson()
		lc.Write(json)
		return
	}

	// Auth pass
	ye = myerror.NewYceError(0, "OK")
	json, _ := ye.EncodeJson()
	lc.Write(json)
}