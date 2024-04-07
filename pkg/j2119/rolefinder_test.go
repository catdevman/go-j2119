package j2119

// import (
// 	"encoding/json"
// 	"testing"
// )
//
// func TestRoleFinder_AssignAdditionalRoleBasedOnRole(t *testing.T) {
// 	cut := NewRoleFinder()
// 	jsonData := `{"a": 3}`
// 	var jsonObj map[string]interface{}
// 	json.Unmarshal([]byte(jsonData), &jsonObj)
//
// 	cut.AddIsARole("OneRole", "AnotherRole")
// 	roles := []string{"OneRole"}
// 	cut.FindMoreRoles(jsonObj, &roles)
//
// 	if len(roles) != 2 || !contains(roles, "AnotherRole") {
// 		t.Errorf("Expected roles to include 'AnotherRole'")
// 	}
// }
//
// func TestRoleFinder_AssignRoleBasedOnFieldValue(t *testing.T) {
// 	cut := NewRoleFinder()
// 	jsonData := `{"a": 3}`
// 	var jsonObj map[string]interface{}
// 	json.Unmarshal([]byte(jsonData), &jsonObj)
//
// 	cut.AddFieldValueRole("MyRole", "a", "3", "NewRole")
// 	roles := []string{"MyRole"}
// 	cut.FindMoreRoles(jsonObj, &roles)
//
// 	if len(roles) != 2 || !contains(roles, "NewRole") {
// 		t.Errorf("Expected roles to include 'NewRole'")
// 	}
// }
//
// func TestRoleFinder_AssignRoleBasedOnFieldPresence(t *testing.T) {
// 	cut := NewRoleFinder()
// 	jsonData := `{"a": 3}`
// 	var jsonObj map[string]interface{}
// 	json.Unmarshal([]byte(jsonData), &jsonObj)
//
// 	cut.AddFieldPresenceRole("MyRole", "a", "NewRole")
// 	roles := []string{"MyRole"}
// 	cut.FindMoreRoles(jsonObj, &roles)
//
// 	if len(roles) != 2 || !contains(roles, "NewRole") {
// 		t.Errorf("Expected roles to include 'NewRole'")
// 	}
// }
//
// func TestRoleFinder_AddRoleToGrandchildField(t *testing.T) {
// 	cut := NewRoleFinder()
// 	jsonData := `{"a": 3}`
// 	var jsonObj map[string]interface{}
// 	json.Unmarshal([]byte(jsonData), &jsonObj)
//
// 	cut.AddGrandchildRole("MyRole", "a", "NewRole")
// 	roles := []string{"MyRole"}
// 	grandchildRoles := cut.FindGrandchildRoles(roles, "a")
//
// 	if len(grandchildRoles) != 1 || !contains(grandchildRoles, "NewRole") {
// 		t.Errorf("Expected grandchild roles to include 'NewRole'")
// 	}
// }
//
// func TestRoleFinder_AddRoleToChildField(t *testing.T) {
// 	cut := NewRoleFinder()
// 	jsonData := `{"a": {"b": 3}}`
// 	var jsonObj map[string]interface{}
// 	json.Unmarshal([]byte(jsonData), &jsonObj)
//
// 	cut.AddChildRole("MyRole", "a", "NewRole")
// 	roles := []string{"MyRole"}
// 	childRoles := cut.FindChildRoles(roles, "a")
//
// 	if len(childRoles) != 1 || !contains(childRoles, "NewRole") {
// 		t.Errorf("Expected child roles to include 'NewRole'")
// 	}
// }
