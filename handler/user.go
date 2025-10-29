package handler

import (
	dblayer "filestore-server/db"
	"filestore-server/util"
	"fmt"
	"net/http"
)

const (
	pwd_salt = "*#890"
)

// SignupHandler : 处理用户注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignupHandler called") // <- 确认函数是否触发
	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/view/signup.html")
		return
	}

	//if err := r.ParseForm(); err != nil {
	//	w.Write([]byte("parse form error"))
	//	return
	//}

	//username := r.Form.Get("username")
	//password := r.Form.Get("password")
	//err := r.ParseMultipartForm(10 << 20) // 最大 10MB
	//if err != nil {
	//	fmt.Println("ParseMultipartForm error:", err)
	//	w.Write([]byte("fail"))
	//	return
	//}
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		fmt.Println("ParseMultipartForm error:", err)
		w.Write([]byte("fail"))
		return
	}
	usernameArr := r.MultipartForm.Value["username"]
	passwordArr := r.MultipartForm.Value["password"]
	if len(usernameArr) == 0 || len(passwordArr) == 0 {
		fmt.Println("username or password empty")
		w.Write([]byte("fail"))
		return
	}
	//username := r.FormValue("username")
	//password := r.FormValue("password")
	username := usernameArr[0]
	password := passwordArr[0]
	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("invalid parameter"))
		return
	}

	fmt.Println("username:", username)
	fmt.Println("password:", password)
	//fmt.Println("Received username:", r.FormValue("username"))
	//fmt.Println("Received password:", r.FormValue("password"))

	enc_passwd := util.Sha1([]byte(password + pwd_salt))
	fmt.Println("enc_passwd:", enc_passwd)

	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		fmt.Println("UserSignup success")
		w.Write([]byte("success"))
	} else {
		fmt.Println("UserSignup failed!")
		w.Write([]byte("fail"))
	}
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/view/signin.html")
		return
	}
	// 后续可以添加登录 POST 逻辑
}
