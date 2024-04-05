// Code generated by ent, DO NOT EDIT.

package ent

import (
	"trec/ent/jobtitle"
	"trec/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	jobtitleFields := schema.JobTitle{}.Fields()
	_ = jobtitleFields
	// jobtitleDescCode is the schema descriptor for code field.
	jobtitleDescCode := jobtitleFields[1].Descriptor()
	// jobtitle.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	jobtitle.CodeValidator = func() func(string) error {
		validators := jobtitleDescCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(code string) error {
			for _, fn := range fns {
				if err := fn(code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// jobtitleDescName is the schema descriptor for name field.
	jobtitleDescName := jobtitleFields[2].Descriptor()
	// jobtitle.NameValidator is a validator for the "name" field. It is called by the builders before save.
	jobtitle.NameValidator = func() func(string) error {
		validators := jobtitleDescName.Validators
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
	// jobtitleDescDescription is the schema descriptor for description field.
	jobtitleDescDescription := jobtitleFields[3].Descriptor()
	// jobtitle.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	jobtitle.DescriptionValidator = jobtitleDescDescription.Validators[0].(func(string) error)
	// jobtitleDescSpecification is the schema descriptor for specification field.
	jobtitleDescSpecification := jobtitleFields[4].Descriptor()
	// jobtitle.SpecificationValidator is a validator for the "specification" field. It is called by the builders before save.
	jobtitle.SpecificationValidator = jobtitleDescSpecification.Validators[0].(func(string) error)
	// jobtitleDescCreatedAt is the schema descriptor for created_at field.
	jobtitleDescCreatedAt := jobtitleFields[5].Descriptor()
	// jobtitle.DefaultCreatedAt holds the default value on creation for the created_at field.
	jobtitle.DefaultCreatedAt = jobtitleDescCreatedAt.Default.(func() time.Time)
	// jobtitleDescID is the schema descriptor for id field.
	jobtitleDescID := jobtitleFields[0].Descriptor()
	// jobtitle.DefaultID holds the default value on creation for the id field.
	jobtitle.DefaultID = jobtitleDescID.Default.(func() uuid.UUID)
}
