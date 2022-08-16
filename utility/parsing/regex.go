package parsing

//
//
// This file contains all regex patterns utilized
//
//

const (
	FUNC              = "f +([a-z]+) +{"                       // Extracts all functions
	DOMAIN_FUNC_VALUE = "([a-z]+)::([a-z]+)\\(([a-z0-9]+)?\\)" // Extracts the domain, function being called and if a value was sent with it
	COMMENT           = "@.+@"                                 // Identifies a comment
	COMMENT_WRONG_RHS = ".+@"                                  // Identifies wrongly formatted comments
	COMMENT_WRONG_LHS = "@.+"                                  // Identifies wrongly formatted comments
)
