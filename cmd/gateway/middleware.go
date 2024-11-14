package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Surya-7890/book_store/gateway/gen"
	"google.golang.org/protobuf/proto"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Println(path)
		if path == "/v1/admin/create" || path == "/v1/admin/login" {
			err := storeUserInRedis("AdminDB", r)
			if err != nil {
				http.Error(w, "Failed to store admin data in Redis", http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)
			return
		}

		if path == "/v1/user/login" || path == "/v1/user/create" {
			err := storeUserInRedis("UserDB", r)
			if err != nil {
				http.Error(w, "Failed to store user data in Redis", http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("authorizaiton")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing header", http.StatusUnauthorized)
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := getUser(token)
		if err != nil {
			http.Error(w, "unable to process your request", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
	})
}

func getUser(token string) (string, error) {
	user, err := Redis.UserDB.Get(ctx, token).Result()
	if err != nil {
		return "", err
	}

	if user != "" {
		return user, nil
	}

	user, err = Redis.AdminDB.Get(ctx, token).Result()
	if err != nil {
		return "", err
	}

	if user != "" {
		return user, nil
	}

	return "", nil
}

// Helper function to store user data in Redis
func storeUserInRedis(dbType string, r *http.Request) error {
	token, tokenOk := r.Context().Value("token").(string)
	username, usernameOk := r.Context().Value("username").(string)
	if !tokenOk || !usernameOk {
		return fmt.Errorf("missing token or username in request context")
	}

	var err error
	if dbType == "AdminDB" {
		err = Redis.AdminDB.Set(ctx, token, username, 0).Err()
	} else if dbType == "UserDB" {
		err = Redis.UserDB.Set(ctx, token, username, 0).Err()
	}
	return err
}

// Custom response interceptor to capture login response
func responseInterceptor(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	if loginResp, ok := resp.(*gen.AdminLoginResponse); ok {
		// Extract token from login response
		username := loginResp.GetStatus()
		fmt.Println("tokens", ctx.Value("user_id"), ctx.Value("token"))
		fmt.Println("from res: ", username)
		loginResp.Status = ""

		// Store token and username in Redis
		// if err := Redis.AdminDB.Set(ctx, token, username, 0).Err(); err != nil {
		// 	return err
		// }
	}

	return nil
}
