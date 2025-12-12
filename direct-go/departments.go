package direct

import (
	"context"
)

// DepartmentTree represents a department tree structure.
type DepartmentTree struct {
	DomainID    interface{}
	Departments []Department
}

// Department represents a department/organizational unit.
type Department struct {
	ID          interface{}
	Name        string
	ParentID    interface{}
	ChildrenIDs []interface{}
	UserCount   int
}

// DepartmentUserCount represents user count statistics for departments.
type DepartmentUserCount struct {
	DepartmentID interface{}
	All          int
	Partial      int
}

// GetDepartmentTree retrieves the department tree for a domain.
func (c *Client) GetDepartmentTree(ctx context.Context, domainID interface{}) (*DepartmentTree, error) {
	params := []interface{}{domainID}
	result, err := c.Call(MethodGetDepartmentTree, params)
	if err != nil {
		return nil, err
	}

	tree := &DepartmentTree{}
	if treeData, ok := result.(map[string]interface{}); ok {
		if v, ok := treeData["domain_id"]; ok {
			tree.DomainID = v
		}
		if departments, ok := treeData["departments"].([]interface{}); ok {
			for _, item := range departments {
				if deptData, ok := item.(map[string]interface{}); ok {
					dept := Department{}
					if v, ok := deptData["id"]; ok {
						dept.ID = v
					}
					if v, ok := deptData["name"].(string); ok {
						dept.Name = v
					}
					if v, ok := deptData["parent_id"]; ok {
						dept.ParentID = v
					}
					if v, ok := deptData["children_ids"].([]interface{}); ok {
						dept.ChildrenIDs = v
					}
					if v, ok := deptData["user_count"].(int); ok {
						dept.UserCount = v
					}
					tree.Departments = append(tree.Departments, dept)
				}
			}
		}
	}

	return tree, nil
}

// GetDepartmentUsers retrieves users in a department.
func (c *Client) GetDepartmentUsers(ctx context.Context, domainID, departmentID interface{}) ([]UserInfo, error) {
	params := []interface{}{domainID, departmentID}
	result, err := c.Call(MethodGetDepartmentUsers, params)
	if err != nil {
		return nil, err
	}

	users := []UserInfo{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if userData, ok := item.(map[string]interface{}); ok {
				user := parseUserInfo(userData)
				users = append(users, user)
			}
		}
	}

	return users, nil
}

// GetDepartmentUserCount retrieves user count statistics for departments.
func (c *Client) GetDepartmentUserCount(ctx context.Context, domainID interface{}) ([]DepartmentUserCount, error) {
	params := []interface{}{domainID}
	result, err := c.Call(MethodGetDepartmentUserCount, params)
	if err != nil {
		return nil, err
	}

	counts := []DepartmentUserCount{}
	if resultData, ok := result.(map[string]interface{}); ok {
		if departments, ok := resultData["departments"].([]interface{}); ok {
			for _, item := range departments {
				if countData, ok := item.(map[string]interface{}); ok {
					count := DepartmentUserCount{}
					if v, ok := countData["department_id"]; ok {
						count.DepartmentID = v
					}
					if v, ok := countData["all"].(int); ok {
						count.All = v
					}
					if v, ok := countData["partial"].(int); ok {
						count.Partial = v
					}
					counts = append(counts, count)
				}
			}
		}
	}

	return counts, nil
}
