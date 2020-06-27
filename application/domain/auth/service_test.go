package auth_test

import (
	"fmt"
	"github.com/casbin/casbin"
	redisAdapter "github.com/casbin/redis-adapter"
	adapter "github.com/memwey/casbin-sqlx-adapter"
	"testing"
)

func TestCasbinGetPolicy(t *testing.T) {
	m := casbin.NewModel()
	m.AddDef("r", "r", "sub, dom, obj, act")
	m.AddDef("p", "p", "sub, dom, obj, act")
	m.AddDef("g", "g", "_, _, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act")

	a := adapter.NewAdapter("postgres", "dbname=c9 user=root password=root host=localhost sslmode=disable")
	e := casbin.NewEnforcer(m, a)
	err := e.LoadPolicy()
	t.Log("Error", err)
	err = e.SavePolicy()
	t.Log("Error", err)

	r := e.GetAllRoles()
	t.Log(r)

	r = e.GetAllObjects()
	t.Log(r)

	r = e.GetAllSubjects()
	t.Log(r)

	rr := e.GetPolicy()
	t.Log(rr)

	p := GetPermission(e, "user:5", "merchant:1")
	t.Log(p)
	p = GetPermission(e, "user:5", "merchant:2")
	t.Log(p)
	p = GetPermission(e, "user:10", "merchant:2.branch:1")
	t.Log(p)
}

func TestCasbinOnRedis(t *testing.T) {
	m := casbin.NewModel()
	m.AddDef("r", "r", "sub, dom, obj, act")
	m.AddDef("p", "p", "sub, dom, obj, act")
	m.AddDef("g", "g", "_, _, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act")

	a := redisAdapter.NewAdapter("tcp", "127.0.0.1:6379")
	e := casbin.NewEnforcer(m, a)
	err := e.LoadPolicy()
	t.Log("Error", err)

	e.AddPolicy("package:basic", "merchant.credit", "allow")
	e.AddPolicy("package:basic", "merchant.user.invite", "allow")
	e.AddPolicy("package:pro", "merchant.branch.create", "allow")
	e.AddPolicy("role:owner", "merchant.user.read", "allow")
	e.AddPolicy("role:cashier", "merchant.bill.read", "allow")

	e.AddGroupingPolicy("user:5", "package:basic", "merchant:2")
	e.AddGroupingPolicy("user:5", "role:owner", "merchant:2")
	e.AddGroupingPolicy("user:5", "package:pro", "merchant:1")
	e.AddGroupingPolicy("user:10", "role:cashier", "merchant:2.branch:1")

	err = e.SavePolicy()
	t.Log("Error", err)

	r := e.GetAllRoles()
	t.Log(r)

	r = e.GetAllObjects()
	t.Log(r)

	r = e.GetAllSubjects()
	t.Log(r)

	rr := e.GetPolicy()
	t.Log(rr)

	p := GetPermission(e, "user:5", "merchant:1")
	t.Log(p)
	p = GetPermission(e, "user:5", "merchant:2")
	t.Log(p)
	p = GetPermission(e, "user:10", "merchant:2.branch:1")
	t.Log(p)
}

func GetPermission(e *casbin.Enforcer, user, domain string) string {
	pd := e.GetRolesForUserInDomain(user, domain)
	var roles, permissions []string
	for _, v := range pd {
		p := e.GetFilteredPolicy(0, v)
		roles = append(roles, v)
		for _, r := range p {
			permissions = append(permissions, r[1])
		}
	}
	return fmt.Sprintf("\nUser %s\nDomain %s\nRole %v\nPermission %v", user, domain, roles, permissions)
}
