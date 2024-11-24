package zig

import (
	"slices"

	"github.com/goreleaser/goreleaser/v2/internal/tmpl"
)

const keyAbi = "Abi"

// Target is a Zig build target.
type Target struct {
	// The zig formatted target (arch-os-abi).
	Target string
	Os     string
	Arch   string
	Abi    string
}

// Fields implements build.Target.
func (t Target) Fields() map[string]string {
	return map[string]string{
		tmpl.KeyOS:   t.Os,
		tmpl.KeyArch: t.Arch,
		keyAbi:       t.Abi,
	}
}

// String implements fmt.Stringer.
func (t Target) String() string {
	return t.Target
}

func convertToGoos(s string) string {
	switch s {
	case "macos":
		return "darwin"
	default:
		return s
	}
}

func convertToGoarch(s string) string {
	switch s {
	case "aarch64":
		return "arm64"
	case "x86_64":
		return "amd64"
	default:
		return s
	}
}

// generated using ./testdata/targets.sh
// not sure if this is correct, but for now, leaving out targets that fail to
// compile a simple C program.
func isValid(target string) bool {
	return slices.Contains([]string{
		"aarch64-linux",
		"aarch64-linux-gnu",
		"aarch64-linux-musl",
		"aarch64-macos",
		"aarch64-macos-none",
		"aarch64-windows",
		"aarch64-windows-gnu",
		"aarch64_be-linux",
		"aarch64_be-linux-gnu",
		"aarch64_be-linux-musl",
		"arm-linux",
		"arm-linux-gnueabi",
		"arm-linux-gnueabihf",
		"arm-linux-musleabi",
		"arm-linux-musleabihf",
		"mips-linux",
		"mips-linux-gnueabi",
		"mips-linux-gnueabihf",
		"mips-linux-musl",
		"mips64-linux",
		"mips64-linux-gnuabi64",
		"mips64-linux-gnuabin32",
		"mips64-linux-musl",
		"mips64el-linux",
		"mips64el-linux-gnuabi64",
		"mips64el-linux-gnuabin32",
		"mips64el-linux-musl",
		"mipsel-linux",
		"mipsel-linux-gnueabi",
		"mipsel-linux-gnueabihf",
		"mipsel-linux-musl",
		"powerpc-linux",
		"powerpc-linux-musl",
		"powerpc64-linux",
		"powerpc64-linux-gnu",
		"powerpc64-linux-musl",
		"powerpc64le-linux",
		"powerpc64le-linux-gnu",
		"powerpc64le-linux-musl",
		"riscv64-linux",
		"riscv64-linux-musl",
		"thumb-linux",
		"thumb-linux-musleabi",
		"thumb-linux-musleabihf",
		"wasm32-wasi",
		"wasm32-wasi-musl",
		"x86-linux",
		"x86-linux-gnu",
		"x86-linux-musl",
		"x86-windows",
		"x86-windows-gnu",
		"x86_64-linux",
		"x86_64-linux-gnu",
		"x86_64-linux-gnux32",
		"x86_64-linux-musl",
		"x86_64-macos",
		"x86_64-macos-none",
		"x86_64-windows",
		"x86_64-windows-gnu",
	}, target)
}

func defaultTargets() []string {
	return []string{
		"x86_64-linux",
		"x86_64-macos",
		"x86_64-windows",
		"aarch64-linux",
		"aarch64-macos",
	}
}
