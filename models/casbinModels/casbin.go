package casbinModels

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func NewCasbinService(db *gorm.DB) (*CasbinService, error) {
	// 创建适配器
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	// 创建模型
	//m, err := model.NewModelFromFile("models/model.conf")
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj


[policy_definition]
p = sub, obj

[role_definition]
g = _,_

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj`)
	if err != nil {
		return nil, err
	}

	// 创建执行器
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	return &CasbinService{Enforcer: e, Adapter: a}, nil
}

// (RoleName, Url, Method) 对应于 `CasbinRule` 表中的 (v0, v1)
type RolePolicy struct {
	RoleName string `gorm:"column:v0"`
	MenuId   string `gorm:"column:v1"`
}

// 获取所有角色组
func (c *CasbinService) GetRoles() ([]string, error) {
	return c.Enforcer.GetAllRoles()
}

// 获取所有角色组权限
func (c *CasbinService) GetRolePolicy() (roles []RolePolicy, err error) {
	err = c.Adapter.GetDb().Model(&gormadapter.CasbinRule{}).Where("ptype = 'p'").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return
}

// 创建角色组权限, 已有的会忽略
func (c *CasbinService) CreateRolePolicy(r RolePolicy) error {
	// 不直接操作数据库，利用enforcer简化操作
	err := c.Enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	_, err = c.Enforcer.AddPolicy(r.RoleName, r.MenuId)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// 修改角色组权限
func (c *CasbinService) UpdateRolePolicy(old, new RolePolicy) error {
	_, err := c.Enforcer.UpdatePolicy([]string{old.RoleName, old.MenuId},
		[]string{new.RoleName, new.MenuId})
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// 删除角色组权限
func (c *CasbinService) DeleteRolePolicy(r RolePolicy) error {
	_, err := c.Enforcer.RemovePolicy(r.RoleName, r.MenuId)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

type User struct {
	UserName  string
	RoleNames []string
}

// GetUsers 获取所有用户以及关联的角色
func (c *CasbinService) GetUsers() (users []User) {
	p, err := c.Enforcer.GetGroupingPolicy()
	if err != nil {
		fmt.Println("GetUsers() c.enforcer.GetGroupingPolicy()")
		return
	}
	usernameUser := make(map[string]*User, 0)
	for _, _p := range p {
		username, usergroup := _p[0], _p[1]
		if v, ok := usernameUser[username]; ok {
			usernameUser[username].RoleNames = append(v.RoleNames, usergroup)
		} else {
			usernameUser[username] = &User{UserName: username, RoleNames: []string{usergroup}}
		}
	}
	for _, v := range usernameUser {
		users = append(users, *v)
	}
	return
}

// UpdateUserRole 角色组中添加用户, 没有组默认创建
func (c *CasbinService) UpdateUserRole(username, rolename string) error {
	_, err := c.Enforcer.AddGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// DeleteUserRole 角色组中删除用户
func (c *CasbinService) DeleteUserRole(username, rolename string) error {
	_, err := c.Enforcer.RemoveGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// 验证用户权限
func (c *CasbinService) CanAccess(username, url, method string) (ok bool, err error) {
	return c.Enforcer.Enforce(username, url, method)
}
