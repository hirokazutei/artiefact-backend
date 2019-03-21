package schema

import "github.com/lestrrat-go/jsval"

var UserGetUserValidator *jsval.JSVal
var M *jsval.ConstraintMap
var R0 jsval.Constraint
var R1 jsval.Constraint
var R2 jsval.Constraint

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String()
	R1 = jsval.String()
	R2 = jsval.String()
	M.SetReference("#/definitions/user/definitions/email", R0)
	M.SetReference("#/definitions/user/definitions/password", R1)
	M.SetReference("#/definitions/user/definitions/username", R2)
	UserGetUserValidator = jsval.New().
		SetName("UserGetUserValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("email", "password", "username").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"email",
					jsval.Reference(M).RefersTo("#/definitions/user/definitions/email"),
				).
				AddProp(
					"password",
					jsval.Reference(M).RefersTo("#/definitions/user/definitions/password"),
				).
				AddProp(
					"username",
					jsval.Reference(M).RefersTo("#/definitions/user/definitions/username"),
				),
		)

}
