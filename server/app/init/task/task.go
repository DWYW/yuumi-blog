package task

type Task struct {
	Host           string
	Port           uint64
	Administrators []AdministratorData
}

func (t Task) Run() {
	t.CreateAdministrators(t.Administrators)
}
