package user

type User struct {
	ID        string // Unique identifier for the user
	Username  string // Username of the user, used for login and identification
	Email     string // Email address of the user, used for notifications and login
	Password  string // Hashed password for the user, used for authentication
	IsActive  bool   // Indicates if the user account is active or not
	IsAdmin   bool   // Indicates if the user has admin privileges
	TenantID  string // ID of the tenant this user belongs to, used for multi-tenancy support
	CreatedAt string // Timestamp when the user was created
	UpdatedAt string // Timestamp when the user was last updated
}
