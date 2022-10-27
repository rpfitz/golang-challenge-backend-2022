package config

import "os"

// Authentication
var SECRET = os.Getenv("SECRET")

// Database
var DB_DRIVER = os.Getenv("DB_DRIVER")
var DB_USER = os.Getenv("DB_USER")
var DB_PASS = os.Getenv("DB_PASS")
var DB_PROTOCOL = os.Getenv("DB_PROTOCOL")
var DB_HOST = os.Getenv("DB_HOST")
var DB_PORT = os.Getenv("DB_PORT")
var DB_NAME = os.Getenv("DB_NAME")
var DB_USERS_TABLE = os.Getenv("DB_USERS_TABLE")
var DB_URL = os.Getenv("DB_URL")

// Google Auth
var (
	GOOGLE_CLIENT_ID     = os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT_URL  = os.Getenv("GOOGLE_REDIRECT_URL")
)

// Server
var (
	SERVER_PROTOCOL = os.Getenv("SERVER_PROTOCOL")
	SERVER_HOST     = os.Getenv("SERVER_HOST")
	PORT            = os.Getenv("PORT")
)
