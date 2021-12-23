package controller

import (
	"LinkedTo-go/model"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"time"
)

func getLoginStatus(c echo.Context) error {
	if username := c.Get("username").(string); username != "" {
		return c.String(http.StatusOK, username)
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}

func register(c echo.Context) error {
	var username, password, email, tel string
	if err := echo.FormFieldBinder(c).
		MustString("username", &username).
		MustString("password", &password).
		MustString("email", &email).
		MustString("tel", &tel).
		BindError(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if len(password) <= 6 || len(username) <= 6 {
		return c.NoContent(http.StatusBadRequest)
	}

	r, _ := regexp.Compile("^\\w+([\\.-]?\\w+)*@\\w+([\\.-]?\\w+)*(\\.\\w{2,3})+$")
	if !r.MatchString(email) {
		return c.NoContent(http.StatusBadRequest)
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	if model.DB.Create(&model.User{
		Username:  username,
		Password:  passwordHashed,
		Email:     email,
		SessionId: uuid.NullUUID{}.UUID, // FIXME: maybe a bug
	}).Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func login(c echo.Context) error {
	var username, password string
	if err := echo.FormFieldBinder(c).
		MustString("username", &username).
		MustString("password", &password).
		BindError(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	user := model.User{Username: username}
	if model.DB.First(&user).Error == gorm.ErrRecordNotFound {
		return c.NoContent(http.StatusNotFound)
	}

	if bcrypt.CompareHashAndPassword(user.Password, []byte(password)) != nil {
		return c.NoContent(http.StatusForbidden)
	}

	sessionID, err := uuid.NewV4()
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	user.SessionId = sessionID
	model.DB.Save(&user)

	cookie := new(http.Cookie)
	cookie.Name = viper.GetString("token.name")
	cookie.Value = sessionID.String()
	cookie.Expires = time.Now().Add(time.Duration(viper.GetInt("token.max_age")) * time.Second)
	cookie.Domain = viper.GetString("token.domain")
	cookie.Secure = viper.GetBool("token.secure")
	cookie.HttpOnly = viper.GetBool("token.http_only")
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}
