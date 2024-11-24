package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/Surya-7890/book_store/gateway/config"
	"github.com/Surya-7890/book_store/gateway/gen"
	"github.com/Surya-7890/book_store/gateway/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type Middleware struct {
	Key   string
	Kafka *config.KafkaWriters
}

func storeInRedis[T utils.AuthResponse](id int32, key string, user T, role string, Kafka *config.KafkaWriters) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token_string, err := token.SignedString([]byte(key))
	if err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte(err.Error()),
		})
		return "", err
	}

	user_bytes, err := json.Marshal(user)
	if err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte(err.Error()),
		})
		return "", err
	}

	id_string := fmt.Sprintf("%v", id)
	var db *redis.Client

	if role == utils.ROLE_ADMIN {
		db = Redis.AdminDB
	} else if role == utils.ROLE_USER {
		db = Redis.UserDB
	}

	if err := db.Set(ctx, id_string, user_bytes, 0).Err(); err != nil {
		Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte(err.Error()),
		})
		return "", err
	}

	return token_string, nil
}

// Custom response interceptor to capture login response
func (m *Middleware) responseInterceptor(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	switch r := resp.(type) {
	case *gen.AdminLoginResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, utils.ROLE_ADMIN, m.Kafka)
		if err != nil {
			return err
		}
		w.Header().Set("Authorization", "Bearer "+token_string)
		break
	case *gen.AdminCreateResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, utils.ROLE_ADMIN, m.Kafka)
		if err != nil {
			return err
		}
		w.Header().Set("Authorization", "Bearer "+token_string)
		break
	case *gen.UserLoginResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, utils.ROLE_USER, m.Kafka)
		if err != nil {
			return err
		}
		w.Header().Set("Authorization", "Bearer "+token_string)
		break
	case *gen.UserSignupResponse:
		user := r.User
		token_string, err := storeInRedis(user.Id, m.Key, user, utils.ROLE_USER, m.Kafka)
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
	if slices.Contains(utils.AuthRoutes[:], utils.AUTH_ROUTES(path)) {
		return metadata.MD{}
	}
	authHeader := req.Header.Get("Authorization")

	if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer ") {
		m.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte("missing header"),
		})
		return metadata.Pairs(utils.AUTH_ERROR, "missing header")
	}

	token_string := strings.TrimPrefix(authHeader, "Bearer ")

	if len(token_string) == 0 {
		m.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte("missing token"),
		})
		return metadata.Pairs(utils.AUTH_ERROR, "missing token")
	}

	token, err := jwt.Parse(token_string, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			m.Kafka.Error.WriteMessages(ctx, kafka.Message{
				Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
				Value: []byte("error while parsing token, unexpected signing method"),
			})
			return nil, fmt.Errorf("error while parsing token, unexpected signing method")
		}
		return []byte(m.Key), nil
	})

	if err != nil {
		m.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte("error while parsing token " + err.Error()),
		})
		return metadata.Pairs(utils.AUTH_ERROR, "error while parsing token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		m.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte("invalid token claims"),
		})
		return metadata.Pairs(utils.AUTH_ERROR, "invalid token claims")
	}

	id, ok := claims["id"]
	if !ok {
		m.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
			Value: []byte("invalid key id for jwt claims"),
		})
		return metadata.Pairs(utils.AUTH_ERROR, "invalid jwt claims")
	}

	id_string := fmt.Sprintf("%v", id)
	map_ := make(map[string]string)

	fmt.Println("id string", id_string)

	if strings.Contains(path, utils.AdminRoutes) && req.Method != http.MethodGet {
		role, ok := claims["role"]
		if !ok {
			m.Kafka.Error.WriteMessages(ctx, kafka.Message{
				Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
				Value: []byte("unauthorized"),
			})
			return metadata.Pairs(utils.AUTH_ERROR, "unauthorized")
		}
		roleStr, ok := role.(string)
		if !ok {
			m.Kafka.Error.WriteMessages(ctx, kafka.Message{
				Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
				Value: []byte("invalid role format"),
			})
			return metadata.Pairs(utils.AUTH_ERROR, "invalid role format")
		}
		if roleStr != utils.ROLE_ADMIN {
			m.Kafka.Error.WriteMessages(ctx, kafka.Message{
				Key:   []byte(utils.JWT_AUTHORIZATION_ERROR),
				Value: []byte("unauthorized, requires admin privilege"),
			})
			return metadata.Pairs(utils.AUTH_ERROR, "unauthorized")
		}
		user, err := Redis.AdminDB.Get(ctx, id_string).Result()
		if err == redis.Nil {
			m.Kafka.Error.WriteMessages(ctx, kafka.Message{
				Key:   []byte(utils.REDIS_ERROR),
				Value: []byte(err.Error()),
			})
			return metadata.Pairs(utils.AUTH_ERROR, err.Error())
		}
		if err == nil {
			fmt.Println("admin db:", user)
			map_["user"] = user
			map_["role"] = string(utils.ROLE_ADMIN)
			return metadata.New(map_)
		} else if err != redis.Nil {
			m.Kafka.Error.WriteMessages(ctx, kafka.Message{
				Key:   []byte(utils.REDIS_ERROR),
				Value: []byte(err.Error()),
			})
			return metadata.Pairs(utils.AUTH_ERROR, err.Error())
		}
	}

	user, err := Redis.UserDB.Get(ctx, id_string).Result()
	if err == redis.Nil {
		m.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.REDIS_ERROR),
			Value: []byte(err.Error()),
		})
		return metadata.Pairs(utils.AUTH_ERROR, err.Error())
	} else if err == nil {
		map_["user"] = user
		map_["role"] = string(utils.ROLE_USER)
		return metadata.New(map_)
	} else if err != redis.Nil {
		m.Kafka.Error.WriteMessages(ctx, kafka.Message{
			Key:   []byte(utils.REDIS_ERROR),
			Value: []byte(err.Error()),
		})
		return metadata.Pairs(utils.AUTH_ERROR, err.Error())
	}

	return metadata.Pairs(utils.AUTH_ERROR, err.Error())
}
