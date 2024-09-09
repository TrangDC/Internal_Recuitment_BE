package dto

import (
	"sort"
	"trec/ent"
	"trec/models"

	"github.com/samber/lo"
)

type EntityPermissionDto interface {
	ProcessAuditTrail(oldRecords, newRecords []*ent.EntityPermission, parentRecordAudit *models.AuditTrailData)
}

type entityPermissionDtoImpl struct{}

func NewEntityPermissionDto() EntityPermissionDto {
	return &entityPermissionDtoImpl{}
}

func (d *entityPermissionDtoImpl) ProcessAuditTrail(oldRecords, newRecords []*ent.EntityPermission, parentRecordAudit *models.AuditTrailData) {
	moduleAudit := models.AuditTrailData{
		Module:    EntityPermissionI18n,
		SubModule: []interface{}{},
	}
	sort.SliceStable(oldRecords, func(i, j int) bool {
		return oldRecords[i].Edges.PermissionEdges.OrderID < oldRecords[j].Edges.PermissionEdges.OrderID
	})
	sort.SliceStable(newRecords, func(i, j int) bool {
		return newRecords[i].Edges.PermissionEdges.OrderID < newRecords[j].Edges.PermissionEdges.OrderID
	})
	permissionGroupsAudit := map[string]*models.AuditTrailData{}
	for _, record := range oldRecords {
		groupPermissionTitle := record.Edges.PermissionEdges.Edges.GroupPermissionEdge.Title
		if _, ok := permissionGroupsAudit[groupPermissionTitle]; !ok {
			permissionGroupsAudit[groupPermissionTitle] = &models.AuditTrailData{
				Module: groupPermissionTitle,
				Create: []interface{}{},
				Update: []interface{}{},
				Delete: []interface{}{},
			}
		}
		newRecord, exists := lo.Find(newRecords, func(item *ent.EntityPermission) bool {
			return item.PermissionID == record.PermissionID
		})
		fieldName := record.Edges.PermissionEdges.Title
		if exists {
			if record.ForAll != newRecord.ForAll || record.ForTeam != newRecord.ForTeam || record.ForOwner != newRecord.ForOwner {
				// update audit
				recordAudit := models.AuditTrailUpdate{
					Field: fieldName,
					Value: models.ValueChange{
						OldValue: d.permissionI18n(record),
						NewValue: d.permissionI18n(newRecord),
					},
				}
				permissionGroupsAudit[groupPermissionTitle].Update = append(permissionGroupsAudit[groupPermissionTitle].Update, recordAudit)
			}
			continue
		}
		// delete audit
		recordAudit := models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: d.permissionI18n(record),
		}
		permissionGroupsAudit[groupPermissionTitle].Delete = append(permissionGroupsAudit[groupPermissionTitle].Delete, recordAudit)
	}
	for _, record := range newRecords {
		groupPermissionTitle := record.Edges.PermissionEdges.Edges.GroupPermissionEdge.Title
		if _, ok := permissionGroupsAudit[groupPermissionTitle]; !ok {
			permissionGroupsAudit[groupPermissionTitle] = &models.AuditTrailData{
				Module: groupPermissionTitle,
				Create: []interface{}{},
				Update: []interface{}{},
				Delete: []interface{}{},
			}
		}
		if !lo.ContainsBy(oldRecords, func(item *ent.EntityPermission) bool {
			return item.PermissionID == record.PermissionID
		}) {
			// create audit
			fieldName := record.Edges.PermissionEdges.Title
			recordAudit := models.AuditTrailCreateDelete{
				Field: fieldName,
				Value: d.permissionI18n(record),
			}
			permissionGroupsAudit[groupPermissionTitle].Create = append(permissionGroupsAudit[groupPermissionTitle].Create, recordAudit)
		}
	}
	// permission group audit
	moduleAudit.SubModule = lo.MapToSlice(
		lo.PickBy(
			permissionGroupsAudit,
			func(_ string, value *models.AuditTrailData) bool {
				return len(value.Create) != 0 || len(value.Update) != 0 || len(value.Delete) != 0
			},
		),
		func(_ string, value *models.AuditTrailData) interface{} { return value },
	)
	parentRecordAudit.SubModule = append(parentRecordAudit.SubModule, moduleAudit)
}

func (d *entityPermissionDtoImpl) permissionI18n(record *ent.EntityPermission) string {
	switch {
	case record.ForAll:
		return "model.entity_permissions.for_all"
	case record.ForTeam:
		return "model.entity_permissions.for_team"
	case record.ForOwner:
		return "model.entity_permissions.for_owner"
	}
	return ""
}
