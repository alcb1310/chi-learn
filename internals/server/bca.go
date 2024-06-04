package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"chi-learn/externals/views/bca"
	"chi-learn/internals/database"
)

var ja = jwtauth.New("HS256", []byte(os.Getenv("SECRET")), nil)

type BCAService struct {
	Service

	Router chi.Router
}

func (s *BCAService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("bca")
		if err != nil {
			slog.Error("AuthMiddleware", "error", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		if cookie.Value == "" {
			slog.Error("AuthMiddleware", "cookie", cookie)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		token, err := jwtauth.VerifyToken(ja, cookie.Value)
		if err != nil {
			slog.Error("AuthMiddleware", "error", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		u, ok := token.(jwt.Token)

		if !ok {
			slog.Error("AuthMiddleware", "error", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		claims, err := u.AsMap(context.Background())
		if err != nil {
			slog.Error("AuthMiddleware", "error", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		user := database.User{}
		user.ID = uuid.MustParse(claims["user_id"].(string))
		user.Email = claims["email"].(string)
		user.Name = claims["name"].(string)
		user.CompanyID = uuid.MustParse(claims["company_id"].(string))

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.Clone(ctx)

		next.ServeHTTP(w, r)
	})
}

func (s *BCAService) BCAHome(w http.ResponseWriter, r *http.Request) error {
	token := readContext(r, "user")
	if token == nil {
		slog.Error("BCAHome", "error", "no token")
		http.Redirect(w, r, "/", http.StatusFound)
		return nil
	}

	slog.Debug("BCAHome", "token", token.ID)
	return render(w, r, bca.Index())
}
