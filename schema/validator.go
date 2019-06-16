package schema

import "github.com/lestrrat-go/jsval"

var ArtiefactUserSignInValidator *jsval.JSVal
var ArtiefactUserSignUpValidator *jsval.JSVal
var M *jsval.ConstraintMap
var R0 jsval.Constraint
var R1 jsval.Constraint
var R2 jsval.Constraint
var R3 jsval.Constraint

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String().RegexpString("^[1-2]{1}[0-9]{3}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}")
	R1 = jsval.String()
	R2 = jsval.String().RegexpString("^[0-9a-zA-Z._-]{4,32}")
	R3 = jsval.String().Format("email")
	M.SetReference("#/definitions/artiefact_user/definitions/birthday", R0)
	M.SetReference("#/definitions/artiefact_user/definitions/password", R1)
	M.SetReference("#/definitions/artiefact_user/definitions/username", R2)
	M.SetReference("#/definitions/registered_email/definitions/email", R3)
	ArtiefactUserSignInValidator = jsval.New().
		SetName("ArtiefactUserSignInValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("password", "username").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"password",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_user/definitions/password"),
				).
				AddProp(
					"username",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_user/definitions/username"),
				),
		)

	ArtiefactUserSignUpValidator = jsval.New().
		SetName("ArtiefactUserSignUpValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("birthday", "password", "username").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"birthday",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_user/definitions/birthday"),
				).
				AddProp(
					"email",
					jsval.Reference(M).RefersTo("#/definitions/registered_email/definitions/email"),
				).
				AddProp(
					"password",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_user/definitions/password"),
				).
				AddProp(
					"username",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_user/definitions/username"),
				),
		)

}
