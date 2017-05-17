package lib

// Response when return to client
type Response struct {
	Error       ErrorCode   `json:"error"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

// ErrorCode type for Response from http
type ErrorCode int

const (
	ResponseOK            ErrorCode = 0
	ResponseJSONParseFail ErrorCode = 1
	ResponseTokenFail     ErrorCode = 2

	ResponseInputValidFrom   ErrorCode = 100
	ResponseInputValidUntil  ErrorCode = 101
	ResponseInputValidAmount ErrorCode = 102

	ResponseWrongUser 		 ErrorCode = 10
	ResponseUserFailPermission 		 ErrorCode = 11

	ResponseDatabaseErrorCoupon     ErrorCode = 201
	ResponseDatabaseNotFoundCoupon  ErrorCode = 202
	ResponseDatabaseIDNotNullCoupon ErrorCode = 203

	ResponseCodeFail                 ErrorCode = 301
	ResponseCodeInvalid              ErrorCode = 302
	ResponseCodeUseLimit             ErrorCode = 303
	ResponseCodeUseInvalidUntil      ErrorCode = 304
	ResponseCodeUseInvalidCategories ErrorCode = 305
	ResponseCodeUseInvalidProducts   ErrorCode = 306
	ResponseCodeUseInvalidUsername   ErrorCode = 307
)

func (e ErrorCode) String() string {
	switch e {
	case ResponseOK:
		return "successful"
	case ResponseJSONParseFail:
		return "Không thể parse json request"
	case ResponseTokenFail:
		return "wrong token"
	case ResponseWrongUser:
		return "wrong email or password"
	case ResponseUserFailPermission:
		return "Ban khong co quyen truy cap vao API nay"
	case ResponseInputValidFrom:
		return "valid_from input sai format hoặc nhỏ hơn ngày hiện tại"
	case ResponseInputValidUntil:
		return "valid_until input sai format hoặc nhỏ hơn ngày hiện tại"
	case ResponseInputValidAmount:
		return "discount phải lớn hơn 0"
	case ResponseDatabaseErrorCoupon:
		return "Loi ket noi database"
	case ResponseDatabaseIDNotNullCoupon:
		return "Id cua coupon khong duoc null hoac khong ton tai"
	case ResponseDatabaseNotFoundCoupon:
		return "không tìm thấy bản ghi yêu cầu"
	case ResponseCodeFail:
		return "Code không được empty"
	case ResponseCodeInvalid:
		return "Code bạn nhập vào không đúng"
	case ResponseCodeUseLimit:
		return "Code bạn dùng đã hết số lần sử dụng"
	case ResponseCodeUseInvalidUntil:
		return "Code bạn dùng đã hết thoi gian su dung"
	case ResponseCodeUseInvalidCategories:
		return "Code khong duoc dung trong loai san pham nay"
	case ResponseCodeUseInvalidProducts:
		return "Code khong duoc dung trong san pham nay"
	case ResponseCodeUseInvalidUsername:
		return "Code khong duoc dung cho ban, co the ban can dang nhap"
	default:
		return "successful"
	}
}
