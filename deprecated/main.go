package deprecated

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var (
	redisClient *redis.Client
)

const (
	sessionName   string = "auth_token"
	userFormName  string = "username"
	passFormName  string = "password"
	tokenFormName string = "token"
	usersRedisSet string = "users"
)

func main() {
	loadEnv()
	runRedis()
	setupUrls()
	startServer()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func runRedis() {
	redisPort := os.Getenv("REDIS_PORT")
	redisUrl := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

}

func setupUrls() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/validate", validateTokenHandler)
}

func startServer() {
	authPort := os.Getenv("AUTH_PORT")
	log.Fatal(http.ListenAndServe(":"+authPort, nil))
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(userFormName)
	password := r.FormValue(passFormName)

	exists, err := redisClient.HExists(usersRedisSet, username).Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if exists {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hashedPassword := hashPassword(password)
	err = redisClient.HSet(usersRedisSet, username, hashedPassword).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := createSessionToken(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  sessionName,
		Value: token,
		Path:  "/",
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sessionName)
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = redisClient.Del(cookie.Value).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue(userFormName)
	password := r.FormValue(passFormName)

	if !isValidUser(username, password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := createSessionToken(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  sessionName,
		Value: token,
		Path:  "/",
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

func validateTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue(tokenFormName)
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	username, err := getUsernameBySessionToken(token)
	if err == redis.Nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Redis error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	resp := map[string]string{"username": username}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
