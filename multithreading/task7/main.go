package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

func main() {
	ctx := context.Background()
	ctx, _ = AddJWTToContext(ctx, 425)
	go func(ctx context.Context) {
		userId, error := ExtractUserIDFromContext(ctx)
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println(userId, error)
	}(ctx)
	time.Sleep(100 * time.Millisecond)

}
func AddJWTToContext(ctx context.Context, userID int) (context.Context, error) {

	var val contextKey = "jwt"
	claims := jwt.MapClaims{
		"userId": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("fake-secret"))
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, val, tokenString), nil
}

func ExtractUserIDFromContext(ctx context.Context) (int, error) {

	var val contextKey = "jwt"
	tokenString, ok := ctx.Value(val).(string)
	if !ok {
		return 0, fmt.Errorf("токен не найден")
	}
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return 0, fmt.Errorf("ошибка при парсинге")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("неверный формат данных в токене")
	}
	userID, ok := claims["userId"].(float64)
	if !ok {
		return 0, fmt.Errorf("userId не найден или не является числом")
	}
	return int(userID), nil
}
