package model

type ApiErrorCode string

// 40x
const TokenExpired ApiErrorCode = "TOKEN_EXPIRED"
const InvalidToken ApiErrorCode = "INVALID_TOKEN"

// 50x
const ServerError ApiErrorCode = "SERVER_ERROR"
