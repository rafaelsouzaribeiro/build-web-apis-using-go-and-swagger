package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/rafaelsouzaribeiro/9-API/internal/dto"
	"github.com/rafaelsouzaribeiro/9-API/internal/entity"
	"github.com/rafaelsouzaribeiro/9-API/internal/infra/database"
)

type UserHandlers struct {
	UserDB database.UserInterface
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandler(DB database.UserInterface) *UserHandlers {
	return &UserHandlers{
		UserDB: DB,
	}
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dtouser dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&dtouser)

	if err != nil {
		//log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	resp, errs := entity.NewUser(dtouser.Name, dtouser.Email, dtouser.Password)

	if errs != nil {
		//log.Println(errs)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(resp)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	message := []byte("Criado com sucesso!\n")
	w.Write(message)
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (u *UserHandlers) GetJwt(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(jwtauth.JWTAuth)
	jwtExpiresin := r.Context().Value("jwtExpiresin").(int)

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	resp, err := u.UserDB.FindByEmail(user.Email)

	if err != nil {
		http.Error(w, "Not Found", http.StatusFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !resp.ValidatePassword(user.Password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": resp.Id.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresin)).Unix(),
	})

	acessToken := dto.GetJWTOutput{AcessToken: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)

}
