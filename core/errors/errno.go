package errors

const (
	// 参数错误
	ErrNoInvalidInput  = 1001
	ErrNoMissingField  = 1002
	ErrNoOutOfRange    = 1003
	ErrNoUnknownParams = 1000

	// 认证和权限错误
	ErrNoUnauthorized          = 2001
	ErrNoForbidden             = 2002
	ErrNoExpiredToken          = 2003
	ErrNoUnknownAuthentication = 2000

	// 资源错误
	ErrNoNotFound               = 3001
	ErrNoDuplicateEntry         = 3002
	ErrNoInsufficientPermission = 3003
	ErrNoUnknownResource        = 3000

	// 内部错误
	ErrNoInternal        = 4001
	ErrNoDatabaseError   = 4002
	ErrNoNetworkError    = 4003
	ErrNoUnknownInternal = 4000

	// 请求错误
	ErrNoBadRequest           = 5001
	ErrNoMethodNotAllowed     = 5002
	ErrNoUnsupportedMediaType = 5003
	ErrNoUnknownRequest       = 5000

	// 文件和上传错误
	ErrNoFileTooLarge      = 6001
	ErrNoInvalidFileFormat = 6002
	ErrNoFileNotFound      = 6003
	ErrNoUnknownFile       = 6000

	// 连接和通信错误
	ErrNoConnectionRefused  = 7001
	ErrNoConnectionTimeout  = 7002
	ErrNoNetworkUnreachable = 7003
	ErrNoUnknownConnection  = 7000

	// 第三方服务错误
	ErrNoExternalServiceUnavailable = 8001
	ErrNoExternalServiceTimeout     = 8002
	ErrNoExternalServiceError       = 8003
	ErrNoUnknownExternalService     = 8000

	// 安全和加密错误
	ErrNoEncryptionFailed = 9001
	ErrNoDecryptionFailed = 9002
	ErrNoInvalidSignature = 9003
	ErrNoUnknownSecurity  = 9000
)

var errorMessages = map[int]string{
	// 参数错误
	ErrNoInvalidInput:  "Invalid input",
	ErrNoMissingField:  "Missing field",
	ErrNoOutOfRange:    "Out of range",
	ErrNoUnknownParams: "Unknown parameter error",

	// 认证和权限错误
	ErrNoUnauthorized:          "Unauthorized",
	ErrNoForbidden:             "Forbidden",
	ErrNoExpiredToken:          "Expired token",
	ErrNoUnknownAuthentication: "Unknown authentication error",

	// 资源错误
	ErrNoNotFound:               "Not found",
	ErrNoDuplicateEntry:         "Duplicate entry",
	ErrNoInsufficientPermission: "Insufficient permission",
	ErrNoUnknownResource:        "Unknown resource error",

	// 内部错误
	ErrNoInternal:        "Internal error",
	ErrNoDatabaseError:   "Database error",
	ErrNoNetworkError:    "Network error",
	ErrNoUnknownInternal: "Unknown internal error",

	// 请求错误
	ErrNoBadRequest:           "Bad request",
	ErrNoMethodNotAllowed:     "Method not allowed",
	ErrNoUnsupportedMediaType: "Unsupported media type",
	ErrNoUnknownRequest:       "Unknown request error",

	// 文件和上传错误
	ErrNoFileTooLarge:      "File too large",
	ErrNoInvalidFileFormat: "Invalid file format",
	ErrNoFileNotFound:      "File not found",
	ErrNoUnknownFile:       "Unknown file error",

	// 连接和通信错误
	ErrNoConnectionRefused:  "Connection refused",
	ErrNoConnectionTimeout:  "Connection timeout",
	ErrNoNetworkUnreachable: "Network unreachable",
	ErrNoUnknownConnection:  "Unknown connection error",

	// 第三方服务错误
	ErrNoExternalServiceUnavailable: "External service unavailable",
	ErrNoExternalServiceTimeout:     "External service timeout",
	ErrNoExternalServiceError:       "External service error",
	ErrNoUnknownExternalService:     "Unknown external service error",

	// 安全和加密错误
	ErrNoEncryptionFailed: "Encryption failed",
	ErrNoDecryptionFailed: "Decryption failed",
	ErrNoInvalidSignature: "Invalid signature",
	ErrNoUnknownSecurity:  "Unknown security error",
}
