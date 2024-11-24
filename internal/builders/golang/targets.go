package golang

type Target struct {
	Target    string
	Goos      string
	Goarch    string
	Goamd64   string
	Go386     string
	Goarm     string
	Goarm64   string
	Gomips    string
	Goppc64   string
	Goriscv64 string
}

// TemplateFields implements build.Target.
func (t Target) TemplateFields() map[string]string {
	return map[string]string{
		"Os":      t.Goos,
		"Arch":    t.Goarch,
		"Amd64":   t.Goamd64,
		"I386":    t.Go386,
		"Arm":     t.Goarm,
		"Arm64":   t.Goarm64,
		"Mips":    t.Gomips,
		"Ppc64":   t.Goppc64,
		"Riscv64": t.Goriscv64,
	}
}

// String implements fmt.Stringer.
func (t Target) String() string {
	return t.Target
}
