// Code generated by ent, DO NOT EDIT.

package permissiongroup

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the permissiongroup type in the database.
	Label = "permission_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldGroupType holds the string denoting the group_type field in the database.
	FieldGroupType = "group_type"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// EdgeGroupPermissionParent holds the string denoting the group_permission_parent edge name in mutations.
	EdgeGroupPermissionParent = "group_permission_parent"
	// EdgeGroupPermissionChildren holds the string denoting the group_permission_children edge name in mutations.
	EdgeGroupPermissionChildren = "group_permission_children"
	// EdgePermissionEdges holds the string denoting the permission_edges edge name in mutations.
	EdgePermissionEdges = "permission_edges"
	// Table holds the table name of the permissiongroup in the database.
	Table = "permission_groups"
	// GroupPermissionParentTable is the table that holds the group_permission_parent relation/edge.
	GroupPermissionParentTable = "permission_groups"
	// GroupPermissionParentColumn is the table column denoting the group_permission_parent relation/edge.
	GroupPermissionParentColumn = "parent_id"
	// GroupPermissionChildrenTable is the table that holds the group_permission_children relation/edge.
	GroupPermissionChildrenTable = "permission_groups"
	// GroupPermissionChildrenColumn is the table column denoting the group_permission_children relation/edge.
	GroupPermissionChildrenColumn = "parent_id"
	// PermissionEdgesTable is the table that holds the permission_edges relation/edge.
	PermissionEdgesTable = "permissions"
	// PermissionEdgesInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionEdgesInverseTable = "permissions"
	// PermissionEdgesColumn is the table column denoting the permission_edges relation/edge.
	PermissionEdgesColumn = "group_id"
)

// Columns holds all SQL columns for permissiongroup fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTitle,
	FieldParentID,
	FieldGroupType,
	FieldOrderID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// GroupType defines the type for the "group_type" enum field.
type GroupType string

// GroupTypeFunction is the default value of the GroupType enum.
const DefaultGroupType = GroupTypeFunction

// GroupType values.
const (
	GroupTypeFunction GroupType = "function"
	GroupTypeSystem   GroupType = "system"
)

func (gt GroupType) String() string {
	return string(gt)
}

// GroupTypeValidator is a validator for the "group_type" field enum values. It is called by the builders before save.
func GroupTypeValidator(gt GroupType) error {
	switch gt {
	case GroupTypeFunction, GroupTypeSystem:
		return nil
	default:
		return fmt.Errorf("permissiongroup: invalid enum value for group_type field: %q", gt)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e GroupType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *GroupType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = GroupType(str)
	if err := GroupTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid GroupType", str)
	}
	return nil
}