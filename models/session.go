package models

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

var globalSessions *session.Manager

/*
func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("redis", sessionConfig)
	go globalSessions.GC()

	//	v := this.GetSession("asta")
	//	if v == nil {
	//		this.SetSession("asta", int(1))
	//		this.Data["num"] = 0
	//	} else {
	//		this.SetSession("asta", v.(int)+1)
	//		this.Data["num"] = v.(int)
	//	}
	//	this.TplName = "index.tpl"

}
*/
func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	usr := sess.Get("usr")
	fmt.Println("session usr: ", usr)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		sess.Set("usr", r.Form["usr"])
	}
}
