package universalbinary

type unitarget struct{}

func (unitarget) String() string { return "darwin_all" }

func (unitarget) TemplateFields() map[string]string {
	return map[string]string{
		"Os":   "darwin",
		"Arch": "all",
	}
}
