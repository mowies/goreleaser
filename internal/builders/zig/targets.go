package zig

import (
	"slices"

	"github.com/goreleaser/goreleaser/v2/pkg/config"
)

type Target struct {
	Target string
	Os     string
	Arch   string
	Abi    string
}

// TemplateFields implements build.Target.
func (t Target) TemplateFields() map[string]string {
	return map[string]string{
		"Os":   t.Os,
		"Arch": t.Arch,
		"Abi":  t.Abi,
	}
}

// String implements fmt.Stringer.
func (t Target) String() string {
	return t.Os + "_" + t.Arch + "_" + t.Abi
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
	case "i386":
		return "386"
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

func targets(b config.Build) []string {
	if len(b.Targets) > 0 {
		return b.Targets
	}

	return []string{
		"x86_64-linux",
		"x86_64-macos",
		"x86_64-windows",
		"aarch64-linux",
		"aarch64-macos",
	}
}
