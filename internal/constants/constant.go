package constants

const (
	SuccessMessage = "Success"

	ErrFailedBadParseRequest       = "Failed to parse request"
	ErrEmailorUsernameAlreadyExist = "Email or username already exist"
	ErrServerError                 = "Internal server error"
)

var (
	// ErrUsernameAlreadyExist = errors.New("Username already exist")
	// ErrEmailAlreadyExist    = errors.New("Email already exist")
	ErrUserNotFound    = "User not found"
	ErrInvalidPassword = "Invalid password"
	ErrTokenExpired    = "Token expired"
	ErrTokenInvalid    = "Token invalid"
)
