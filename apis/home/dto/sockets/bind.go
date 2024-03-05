package dtosockets

// BindParams 绑定用户ID
type BindParams struct {
	UUID string `json:"uuid" validate:"required"` //	标识
}
