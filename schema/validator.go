package schema

import "github.com/lestrrat-go/jsval"

var ArtiefactObjectPostTextValidator *jsval.JSVal
var ArtiefactUserGetUserValidator *jsval.JSVal
var ArtiefactUserSignInValidator *jsval.JSVal
var ArtiefactUserSignUpValidator *jsval.JSVal
var ArtiefactUserUsernameAvailabilityValidator *jsval.JSVal
var M *jsval.ConstraintMap
var R0 jsval.Constraint
var R1 jsval.Constraint
var R2 jsval.Constraint
var R3 jsval.Constraint
var R4 jsval.Constraint
var R5 jsval.Constraint
var R6 jsval.Constraint
var R7 jsval.Constraint
var R8 jsval.Constraint
var R9 jsval.Constraint
var R10 jsval.Constraint

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String().RegexpString("^.{0,140}$")
	R1 = jsval.Number()
	R2 = jsval.Number()
	R3 = jsval.String()
	R4 = jsval.String().RegexpString("^.{1,64}$")
	R5 = jsval.String().Enum("text", "image", "audieo", "video")
	R6 = jsval.String().RegexpString("^[1-2]{1}[0-9]{3}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}$")
	R7 = jsval.Integer()
	R8 = jsval.String()
	R9 = jsval.String().RegexpString("^[0-9a-zA-Z._-]{4,32}$")
	R10 = jsval.String().Format("email")
	M.SetReference("#/definitions/artiefact_object/definitions/hint", R0)
	M.SetReference("#/definitions/artiefact_object/definitions/latitude", R1)
	M.SetReference("#/definitions/artiefact_object/definitions/longitude", R2)
	M.SetReference("#/definitions/artiefact_object/definitions/text", R3)
	M.SetReference("#/definitions/artiefact_object/definitions/title", R4)
	M.SetReference("#/definitions/artiefact_object/definitions/type", R5)
	M.SetReference("#/definitions/artiefact_user/definitions/birthday", R6)
	M.SetReference("#/definitions/artiefact_user/definitions/id", R7)
	M.SetReference("#/definitions/artiefact_user/definitions/password", R8)
	M.SetReference("#/definitions/artiefact_user/definitions/username", R9)
	M.SetReference("#/definitions/registered_email/definitions/email", R10)
	ArtiefactObjectPostTextValidator = jsval.New().
		SetName("ArtiefactObjectPostTextValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("hint", "latitude", "longitude", "title", "type", "user_id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"hint",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_object/definitions/hint"),
				).
				AddProp(
					"latitude",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_object/definitions/latitude"),
				).
				AddProp(
					"longitude",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_object/definitions/longitude"),
				).
				AddProp(
					"text",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_object/definitions/text"),
				).
				AddProp(
					"title",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_object/definitions/title"),
				).
				AddProp(
					"type",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_object/definitions/type"),
				).
				AddProp(
					"user_id",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_user/definitions/id"),
				),
		)

	ArtiefactUserGetUserValidator = jsval.New().
		SetName("ArtiefactUserGetUserValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.EmptyConstraint,
		)

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

	ArtiefactUserUsernameAvailabilityValidator = jsval.New().
		SetName("ArtiefactUserUsernameAvailabilityValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("username").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"username",
					jsval.Reference(M).RefersTo("#/definitions/artiefact_user/definitions/username"),
				),
		)

}
