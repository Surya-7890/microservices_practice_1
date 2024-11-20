package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/Surya-7890/book_store/gateway/gen"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func storeInRedis[T AuthResponse](id int32, key string, user T, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token_string, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	user_bytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	id_string := fmt.Sprintf("%v", id)
	var db *redis.Client

	if role == ROLE_ADMIN {
		db = Redis.AdminDB
	} else if role == ROLE_USER {
		db = Redis.UserDB
	}

	if err := db.Set(ctx, id_string, user_bytes, 0).Err(); err != nil {
		return "", err
	}

	return token_string, nil
}

// Custom response interceptor to capture login response
func (m *Middleware) responseInterceptor(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	switch r := resp.(type) {
	case *gen.AdminLoginResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, ROLE_ADMIN)
		if err != nil {
			return err
		}
		w.Header().Set("Authorization", "Bearer "+token_string)
		break
	case *gen.AdminCreateResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, ROLE_ADMIN)
		if err != nil {
			return err
		}
		w.Header().Set("Authorization", "Bearer "+token_string)
		break
	case *gen.UserLoginResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, ROLE_USER)
		if err != nil {
			return err
		}
		w.Header().Set("Authorization", "Bearer "+token_string)
		break
	case *gen.UserSignupResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, ROLE_USER)
		if err != nil {
			return err
		}
		w.Header().Set("Authorization", "Bearer "+token_string)
		break
	}
	return nil
}

func (m *Middleware) requestInterceptor(ctx context.Context, req *http.Request) metadata.MD {
	path := req.URL.Path
	if slices.Contains(AuthRoutes[:], AUTH_ROUTES(path)) {
		return metadata.MD{}
	}
	authHeader := req.Header.Get("Authorization")

	if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer ") {
		return metadata.Pairs(AUTH_ERROR, "missing header")
	}

	token_string := strings.TrimPrefix(authHeader, "Bearer ")

	if len(token_string) == 0 {
		return metadata.Pairs(AUTH_ERROR, "missing token")
	}

	token, err := jwt.Parse(token_string, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error while parsing token, unexpected signing method")
		}
		return []byte(m.Key), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return metadata.Pairs(AUTH_ERROR, "invalid token")
	}

	fmt.Println(claims)

	id, ok := claims["id"]
	if !ok {
		return metadata.Pairs(AUTH_ERROR, "invalid jwt claims")
	}
	id_string := fmt.Sprintf("%v", id)
	map_ := make(map[string]string)

	fmt.Println("id string", id_string)

	if strings.Contains(path, AdminRoutes) && req.Method != http.MethodGet {
		role, ok := claims["role"]
		if !ok {
			return metadata.Pairs(AUTH_ERROR, "unauthorized")
		}
		roleStr, ok := role.(string)
		if !ok {
			return metadata.Pairs(AUTH_ERROR, "invalid role format")
		}
		if roleStr != ROLE_ADMIN {
			return metadata.Pairs(AUTH_ERROR, "unauthorized")
		}
		user, err := Redis.AdminDB.Get(ctx, id_string).Result()
		if err == redis.Nil {
			fmt.Println("redis nil ngommala")
		}
		if err == nil {
			fmt.Println("admin db:", user)
			map_["user"] = user
			map_["role"] = string(ROLE_ADMIN)
			return metadata.New(map_)
		} else if err != redis.Nil {
			log.Printf("Error fetching from AdminDB: %v", err)
			return metadata.Pairs(AUTH_ERROR, err.Error())
		}
	}

	user, err := Redis.UserDB.Get(ctx, id_string).Result()
	if err == redis.Nil {
		log.Printf("Error fetching from UserDB: %v", err)
	} else if err == nil {
		fmt.Println("user db:", user)
		map_["user"] = user
		map_["role"] = string(ROLE_USER)
		return metadata.New(map_)
	}

	return metadata.Pairs(AUTH_ERROR, err.Error())
}
