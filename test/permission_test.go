package test

import (
	"fmt"
	"github.com/Catlordx/CampusTrade/internal/db/mysql"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/permission"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/role"
	"github.com/Catlordx/CampusTrade/internal/db/mysql/user"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPermissionsByRole(t *testing.T) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}
	err = db.AutoMigrate(&mysql.RolePermission{})

	userPermissions := user.RolePermission(db, role.User)
	require.Len(t, userPermissions, 3, "the number of roles of user should be 3")
	for _, rolePermission := range userPermissions {
		fmt.Printf("role of user has permission of %v\n", rolePermission)
	}

	adminPermissions := user.RolePermission(db, role.Admin)
	require.Len(t, adminPermissions, 4, "the number of roles of admin should be 3")
	for _, rolePermission := range adminPermissions {
		fmt.Printf("role of admin has permission of %v\n", rolePermission)
	}

	memberPermissions := user.RolePermission(db, "member")
	require.Len(t, memberPermissions, 0, "the number of roles of member should be 0")
}

func TestHasPermission(t *testing.T) {
	conf := mysql.DbConfig{}
	db, err := mysql.Connect(&conf)
	if err != nil {
		t.Fatalf("Failed to load DB config from viper: %v", err)
	}
	err = db.AutoMigrate(&mysql.RolePermission{})

	result := user.HasPermission(db, role.User, permission.InquireInfo)
	require.True(t, result, "role of user should have the permission")

	result = user.HasPermission(db, role.User, permission.ModifyAnyoneInfo)
	require.False(t, result, "role of user should not have the permission")

	result = user.HasPermission(db, role.Admin, permission.OperateFavorite)
	require.False(t, result, "role of admin should not have the permission")

	result = user.HasPermission(db, role.User, "test_permission")
	require.False(t, result, "role of user should not have the permission")

	result = user.HasPermission(db, "", permission.InquireInfo)
	require.False(t, result, "empty role should not have the permission")

	result = user.HasPermission(db, role.User, "")
	require.False(t, result, "empty role should not have empty permission")
}
