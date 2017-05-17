package lib

const (
	ROLE_ADMIN = 1
	ROLE_NORMAL = 2
	ROLE_VIEW = 4
	ROLE_MODERATE = 8
//vmintam added analytic log 
	// QUIET = ERROR | CRITICAL  //setting for errors only
	// NORMAL = INFO | WARN | ERROR | CRITICAL // default setting - all besides debug
	ROLE_SUPERADMIN = 255
	ROLE_NOTHING = 0

)