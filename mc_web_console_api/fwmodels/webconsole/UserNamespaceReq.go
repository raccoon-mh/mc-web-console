package webconsole

type UserNamespaceReq struct {
	Ns []string `json:"ns"`
	Us []string `json:"us"`
}
