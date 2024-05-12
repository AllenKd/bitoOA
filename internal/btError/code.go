package btError

const (
	codeUnAuthorized Code = iota + 1024
	codeInvalidBody
	codeUserNotFound
	codeUnknown
)

const (
	typeAuth Type = iota + 32
	typeInvalidData
	typeDb
	typeOther
)

var (
	UnAuthorized = &Error{
		Code:    codeUnAuthorized,
		Message: "token unauthorized",
		Type:    typeAuth,
	}

	InvalidBody = &Error{
		Code:    codeInvalidBody,
		Message: "invalid body",
		Type:    typeInvalidData,
	}

	UserNotFound = &Error{
		Code:    codeUserNotFound,
		Message: "user not found",
		Type:    typeInvalidData,
	}

	Unknown = &Error{
		Code:    codeUnknown,
		Message: "unknown",
		Type:    typeOther,
	}
)
