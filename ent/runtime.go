// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"trec/ent/audittrail"
	"trec/ent/schema"
	"trec/ent/team"
	"trec/ent/teammanager"
	"trec/ent/user"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	audittrailFields := schema.AuditTrail{}.Fields()
	_ = audittrailFields
	// audittrailDescNote is the schema descriptor for note field.
	audittrailDescNote := audittrailFields[5].Descriptor()
	// audittrail.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	audittrail.NoteValidator = audittrailDescNote.Validators[0].(func(string) error)
	// audittrailDescCreatedAt is the schema descriptor for created_at field.
	audittrailDescCreatedAt := audittrailFields[7].Descriptor()
	// audittrail.DefaultCreatedAt holds the default value on creation for the created_at field.
	audittrail.DefaultCreatedAt = audittrailDescCreatedAt.Default.(func() time.Time)
	// audittrailDescID is the schema descriptor for id field.
	audittrailDescID := audittrailFields[0].Descriptor()
	// audittrail.DefaultID holds the default value on creation for the id field.
	audittrail.DefaultID = audittrailDescID.Default.(func() uuid.UUID)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescName is the schema descriptor for name field.
	teamDescName := teamFields[1].Descriptor()
	// team.NameValidator is a validator for the "name" field. It is called by the builders before save.
	team.NameValidator = func() func(string) error {
		validators := teamDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// teamDescCreatedAt is the schema descriptor for created_at field.
	teamDescCreatedAt := teamFields[2].Descriptor()
	// team.DefaultCreatedAt holds the default value on creation for the created_at field.
	team.DefaultCreatedAt = teamDescCreatedAt.Default.(func() time.Time)
	// teamDescUpdatedAt is the schema descriptor for updated_at field.
	teamDescUpdatedAt := teamFields[3].Descriptor()
	// team.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	team.DefaultUpdatedAt = teamDescUpdatedAt.Default.(func() time.Time)
	// teamDescID is the schema descriptor for id field.
	teamDescID := teamFields[0].Descriptor()
	// team.DefaultID holds the default value on creation for the id field.
	team.DefaultID = teamDescID.Default.(func() uuid.UUID)
	teammanagerFields := schema.TeamManager{}.Fields()
	_ = teammanagerFields
	// teammanagerDescCreatedAt is the schema descriptor for created_at field.
	teammanagerDescCreatedAt := teammanagerFields[3].Descriptor()
	// teammanager.DefaultCreatedAt holds the default value on creation for the created_at field.
	teammanager.DefaultCreatedAt = teammanagerDescCreatedAt.Default.(func() time.Time)
	// teammanagerDescUpdatedAt is the schema descriptor for updated_at field.
	teammanagerDescUpdatedAt := teammanagerFields[4].Descriptor()
	// teammanager.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	teammanager.DefaultUpdatedAt = teammanagerDescUpdatedAt.Default.(func() time.Time)
	// teammanagerDescID is the schema descriptor for id field.
	teammanagerDescID := teammanagerFields[0].Descriptor()
	// teammanager.DefaultID holds the default value on creation for the id field.
	teammanager.DefaultID = teammanagerDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescWorkEmail is the schema descriptor for work_email field.
	userDescWorkEmail := userFields[2].Descriptor()
	// user.WorkEmailValidator is a validator for the "work_email" field. It is called by the builders before save.
	user.WorkEmailValidator = userDescWorkEmail.Validators[0].(func(string) error)
	// userDescOid is the schema descriptor for oid field.
	userDescOid := userFields[3].Descriptor()
	// user.OidValidator is a validator for the "oid" field. It is called by the builders before save.
	user.OidValidator = userDescOid.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
