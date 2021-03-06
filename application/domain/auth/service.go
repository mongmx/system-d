package auth

import (
	"os"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	redisadapter "github.com/casbin/redis-adapter/v2"
	"github.com/dgrijalva/jwt-go"
)

// Service is the auth service
type Service interface {
	GetAllRole()
	GetAllResource()
	GetAllDomain()
	Authorize(user, domain, resource string) (bool, error)
	CheckUser(username, password string) (*User, error)
	CreateToken(user *User) (string, error)
}

type service struct {
	repo     Repository
	enforcer *casbin.Enforcer
}

// NewService creates new auth service
func NewService(authRepo Repository, ra *redisadapter.Adapter) (Service, error) {
	a := redisadapter.NewAdapter("tcp", "127.0.0.1:6379")
	m := model.NewModel()
	m.AddDef("r", "r", "sub, dom, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub, r.dom) && r.obj == p.obj && r.act == p.act")
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}

	e.AddPolicy("role:user", "resource:profile", "allow")
	e.AddGroupingPolicy("user:test", "role:user", "domain:system")
	e.SavePolicy()

	s := service{authRepo, e}
	return &s, nil
}

func (s *service) GetAllRole() {
	s.enforcer.GetAllRoles()
}

func (s *service) GetAllResource() {
	// s.enforcer.Get
}

func (s *service) GetAllDomain() {

}

func (s *service) Authorize(user, domain, resource string) (bool, error) {
	// ps := s.enforcer.GetPermissionsForUserInDomain("user"+user, "domain"+domain)
	// for _, p := range ps {

	// }
	return true, nil
}

func (s *service) CheckUser(username, Password string) (*User, error) {
	return &User{
		ID: "zzzzzz",
	}, nil
}

func (s *service) CreateToken(user *User) (string, error) {
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["session_id"] = user.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}
