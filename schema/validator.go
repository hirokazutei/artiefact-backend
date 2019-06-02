package schema

import "github.com/lestrrat-go/jsval"

var UserSigninValidator *jsval.JSVal
var UserSignupValidator *jsval.JSVal
var M *jsval.ConstraintMap
var R0 jsval.Constraint
var R1 jsval.Constraint
var R2 jsval.Constraint
var R3 jsval.Constraint

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String().RegexpString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")
	R1 = jsval.String()
	R2 = jsval.String()
	R3 = jsval.String()
	M.SetReference("#/definitions/user/definitions/birthday", R0)
	M.SetReference("#/definitions/user/definitions/email", R1)
	M.SetReference("#/definitions/user/definitions/password", R2)
	M.SetReference("#/definitions/user/definitions/username", R3)
	UserSigninValidator = jsval.New().
		SetName("UserSigninValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("password", "username").
				AdditionalProperties(
					jsval.EmptyConstraint,
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

	UserSignupValidator = jsval.New().
		SetName("UserSignupValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("birthday", "email", "password", "username").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"birthday",
					jsval.Reference(M).RefersTo("#/definitions/user/definitions/birthday"),
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
