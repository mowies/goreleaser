package golang

import "github.com/goreleaser/goreleaser/v2/internal/tmpl"

// Target is a Go build target.
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

// Fields implements build.Target.
func (t Target) Fields() map[string]string {
	return map[string]string{
		tmpl.KeyOS:      t.Goos,
		tmpl.KeyArch:    t.Goarch,
		tmpl.KeyAmd64:   t.Goamd64,
		tmpl.Key386:     t.Go386,
		tmpl.KeyArm:     t.Goarm,
		tmpl.KeyArm64:   t.Goarm64,
		tmpl.KeyMips:    t.Gomips,
		tmpl.KeyPpc64:   t.Goppc64,
		tmpl.KeyRiscv64: t.Goriscv64,
	}
}

// String implements fmt.Stringer.
func (t Target) String() string {
	return t.Target
}
