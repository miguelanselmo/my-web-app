package controllers

import (
	"log"
	"net/http"

	repository "github.com/miguelanselmo/my-web-app/repository"
	utils "github.com/miguelanselmo/my-web-app/utils"

	models "github.com/miguelanselmo/my-web-app/models"
)

func signupView(ctrl *ControllerSrvc, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/signup.html")
}

func signup(ctrl *ControllerSrvc, w http.ResponseWriter, r *http.Request) {
	hasValue := r.FormValue("email")
	if hasValue == "" {
		log.Println("View Signup")
		signupView(ctrl, w, r)
	} else {
		log.Println("Ctrl Signup")
		email := r.FormValue("email")
		password := r.FormValue("password")
		userSignup := models.UserModel{
			Email:    email,
			Password: password,
		}
		_, exist := ctrl.repo.GetUser(userSignup)
		if exist {
			log.Println("User already exist")
			w.Write([]byte("User already exists"))
			return
		} else {
			passwordHash, err := utils.GetPasswordHash(password)
			if err != nil {
				log.Fatalln(err)
				w.Write([]byte("Error"))
				return
			}
			userSignup.Password = passwordHash
			ctrl.repo.CreateUser(userSignup)
			log.Println("User created")
			w.Write([]byte("User created"))
			return
		}
	}
}

func signinView(ctrl *ControllerSrvc, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/signin.html")
}

func signin(ctrl *ControllerSrvc, w http.ResponseWriter, r *http.Request) {
	hasValue := r.FormValue("email")
	if hasValue == "" {
		log.Println("View Signin")
		signinView(ctrl, w, r)
	} else {
		log.Println("Ctrl Signin")
		email := r.FormValue("email")
		password := r.FormValue("password")
		userSignin := models.UserModel{
			Email:    email,
			Password: password,
		}
		userRepo, exist := ctrl.repo.GetUser(userSignin)
		if exist {
			passwordHash, err := utils.GetPasswordHash(userSignin.Password)
			if err != nil {
				log.Fatalln(err)
				w.Write([]byte("Error"))
				return
			}
			if userRepo.Password == passwordHash {
				log.Println("Signin success")
				w.Write([]byte("Sigbin success"))
				return
			} else {
				log.Println("Password incorrect")
				w.Write([]byte("Password incorrect"))
				return
			}
		} else {
			log.Println("User created")
			w.Write([]byte("User does not exist"))
		}
	}

}

func signout(ctrl *ControllerSrvc, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout"))
}

type ControllerSrvc struct {
	log  *log.Logger
	repo *repository.RepositorySrvc
}

func New(repo *repository.RepositorySrvc, log *log.Logger) *ControllerSrvc {
	return &ControllerSrvc{repo: repo, log: log}
}

func (ctrl *ControllerSrvc) AuthController(w http.ResponseWriter, r *http.Request) {
	log.Println("URL: ", r.URL.Path)
	switch r.URL.Path {
	case "/signup":
		signup(ctrl, w, r)
	case "/signin":
		signin(ctrl, w, r)
	case "/signout":
		signout(ctrl, w, r)
	default:
		http.NotFound(w, r)
	}
}
