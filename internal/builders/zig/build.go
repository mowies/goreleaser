package zig

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/caarlos0/log"
	"github.com/goreleaser/goreleaser/v2/internal/artifact"
	"github.com/goreleaser/goreleaser/v2/internal/gio"
	"github.com/goreleaser/goreleaser/v2/internal/tmpl"
	api "github.com/goreleaser/goreleaser/v2/pkg/build"
	"github.com/goreleaser/goreleaser/v2/pkg/config"
	"github.com/goreleaser/goreleaser/v2/pkg/context"
)

// Default builder instance.
//
//nolint:gochecknoglobals
var Default = &Builder{}

//nolint:gochecknoinits
func init() {
	api.Register("zig", Default)
}

// Builder is golang builder.
type Builder struct{}

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

// WithDefaults implements build.Builder.
func (b *Builder) WithDefaults(build config.Build) (config.Build, error) {
	log.Warn("you are using the experimental Zig builder")
	build.Targets = targets(build)

	if build.GoBinary == "" {
		build.GoBinary = "zig"
	}

	if build.Command == "" {
		build.Command = "build"
	}

	if build.Dir == "" {
		build.Dir = "."
	}

	if build.Main != "" {
		return build, errors.New("main is not used for zig")
	}

	if len(build.Ldflags) > 0 {
		return build, errors.New("ldflags is not used for zig")
	}

	if len(slices.Concat(
		build.Goos,
		build.Goarch,
		build.Goamd64,
		build.Go386,
		build.Goarm,
		build.Goarm64,
		build.Gomips,
		build.Goppc64,
		build.Goriscv64,
	)) > 0 {
		return build, errors.New("all go* fields are not used for zig, set targets instead")
	}

	if len(build.Ignore) > 0 {
		return build, errors.New("ignore is not used for zig, set targets instead")
	}

	if build.Buildmode != "" {
		return build, errors.New("buildmode is not used for zig")
	}

	if len(build.Tags) > 0 {
		return build, errors.New("tags is not used for zig")
	}

	if len(build.Asmflags) > 0 {
		return build, errors.New("asmtags is not used for zig")
	}

	for _, t := range build.Targets {
		if !isValid(t) {
			return build, fmt.Errorf("invalid target: %s", t)
		}
	}

	return build, nil
}

// Build implements build.Builder.
func (b *Builder) Build(ctx *context.Context, build config.Build, options api.Options) error {
	prefix := filepath.Dir(options.Path)
	options.Path = filepath.Join(prefix, "bin", options.Name)
	a := &artifact.Artifact{
		Type:      artifact.Binary,
		Path:      options.Path,
		Name:      options.Name,
		Goos:      options.Goos,
		Goarch:    options.Goarch,
		Goamd64:   options.Goamd64,
		Go386:     options.Go386,
		Goarm:     options.Goarm,
		Goarm64:   options.Goarm64,
		Gomips:    options.Gomips,
		Goppc64:   options.Goppc64,
		Goriscv64: options.Goriscv64,
		Target:    options.Target,
		Extra: map[string]interface{}{
			artifact.ExtraBinary: strings.TrimSuffix(filepath.Base(options.Path), options.Ext),
			artifact.ExtraExt:    options.Ext,
			artifact.ExtraID:     build.ID,
		},
	}

	gobin, err := tmpl.New(ctx).WithBuildOptions(options).Apply(build.GoBinary)
	if err != nil {
		return err
	}

	command := []string{
		gobin,
		build.Command,
		"-Dtarget=" + options.Target,
		"-p", prefix,
	}

	env := []string{}
	env = append(env, ctx.Env.Strings()...)
	for _, e := range build.Env {
		ee, err := tmpl.New(ctx).WithEnvS(env).WithArtifact(a).Apply(e)
		if err != nil {
			return err
		}
		log.Debugf("env %q evaluated to %q", e, ee)
		if ee != "" {
			env = append(env, ee)
		}
	}

	// TODO: flags tpl

	/* #nosec */
	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	cmd.Env = env
	cmd.Dir = build.Dir
	log.Debug("running")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, string(out))
	}
	if s := string(out); s != "" {
		log.WithField("cmd", command).Info(s)
	}

	modTimestamp, err := tmpl.New(ctx).WithEnvS(env).WithArtifact(a).Apply(build.ModTimestamp)
	if err != nil {
		return err
	}
	if err := gio.Chtimes(options.Path, modTimestamp); err != nil {
		return err
	}

	ctx.Artifacts.Add(a)
	return nil
}
