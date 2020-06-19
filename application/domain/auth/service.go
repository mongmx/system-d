package auth

// Service is the auth service
type Service interface {
	
}

// NewService creates new auth service
func NewService(authRepo Repository, a *adapter.Adapter) (Service, error) {
	m := casbin.NewModel()
	m.AddDef("r", "r", "sub, dom, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub, r.dom) && r.obj == p.obj && r.act == p.act")
	e := casbin.NewEnforcer(m, a)
	err := e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	s := service{authRepo, e}
	return &s, nil
}

