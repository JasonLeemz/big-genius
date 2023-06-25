package errors

const (
	// 未指定错误
	ErrNoUnknown = -1

	// 成功
	ErrNoSuccess = 0

	// 参数错误
	ErrNoInvalidInput = 1001
	ErrNoMissingField = 1002
	ErrNoOutOfRange   = 1003

	// 认证和权限错误
	ErrNoUnauthorized = 2001
	ErrNoForbidden    = 2002
	ErrNoExpiredToken = 2003

	// 资源错误
	ErrNoNotFound               = 3001
	ErrNoDuplicateEntry         = 3002
	ErrNoInsufficientPermission = 3003

	// 内部错误
	ErrNoInternal      = 4001
	ErrNoDatabaseError = 4002
	ErrNoNetworkError  = 4003

	// 请求错误
	ErrNoBadRequest           = 5001
	ErrNoMethodNotAllowed     = 5002
	ErrNoUnsupportedMediaType = 5003

	// 文件和上传错误
	ErrNoFileTooLarge      = 6001
	ErrNoInvalidFileFormat = 6002
	ErrNoFileNotFound      = 6003

	// 连接和通信错误
	ErrNoConnectionRefused  = 7001
	ErrNoConnectionTimeout  = 7002
	ErrNoNetworkUnreachable = 7003

	// 第三方服务错误
	ErrNoExternalServiceUnavailable = 8001
	ErrNoExternalServiceTimeout     = 8002
	ErrNoExternalServiceError       = 8003

	// 安全和加密错误
	ErrNoEncryptionFailed = 9001
	ErrNoDecryptionFailed = 9002
	ErrNoInvalidSignature = 9003
)

var errorMessages = map[int]string{
	// 未指定错误
	ErrNoUnknown: "Unknown Error",

	// 成功
	ErrNoSuccess: "Success",

	// 参数错误
	ErrNoInvalidInput: "Invalid input",
	ErrNoMissingField: "Missing field",
	ErrNoOutOfRange:   "Out of range",

	// 认证和权限错误
	ErrNoUnauthorized: "Unauthorized",
	ErrNoForbidden:    "Forbidden",
	ErrNoExpiredToken: "Expired token",

	// 资源错误
	ErrNoNotFound:               "Not found",
	ErrNoDuplicateEntry:         "Duplicate entry",
	ErrNoInsufficientPermission: "Insufficient permission",

	// 内部错误
	ErrNoInternal:      "Internal error",
	ErrNoDatabaseError: "Database error",
	ErrNoNetworkError:  "Network error",

	// 请求错误
	ErrNoBadRequest:           "Bad request",
	ErrNoMethodNotAllowed:     "Method not allowed",
	ErrNoUnsupportedMediaType: "Unsupported media type",

	// 文件和上传错误
	ErrNoFileTooLarge:      "File too large",
	ErrNoInvalidFileFormat: "Invalid file format",
	ErrNoFileNotFound:      "File not found",

	// 连接和通信错误
	ErrNoConnectionRefused:  "Connection refused",
	ErrNoConnectionTimeout:  "Connection timeout",
	ErrNoNetworkUnreachable: "Network unreachable",

	// 第三方服务错误
	ErrNoExternalServiceUnavailable: "External service unavailable",
	ErrNoExternalServiceTimeout:     "External service timeout",
	ErrNoExternalServiceError:       "External service error",

	// 安全和加密错误
	ErrNoEncryptionFailed: "Encryption failed",
	ErrNoDecryptionFailed: "Decryption failed",
	ErrNoInvalidSignature: "Invalid signature",
}
