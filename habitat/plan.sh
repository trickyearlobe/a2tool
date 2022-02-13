# This file is the heart of your application's habitat.
# See full docs at https://www.habitat.sh/docs/reference/plan-syntax/

pkg_name=a2tool
pkg_origin=trickyearlobe
pkg_version="0.1.0"
pkg_maintainer="Richard Nixon <richard.nixon@btinternet.com>"
pkg_license=("Apache-2.0")

# The scaffolding base for this plan.
pkg_scaffolding="core/scaffolding-go"

#scaffolding_go_build_deps=(
#  # "github.com/trickyearlobe/a2tool"
#  "github.com/spf13/cobra"
# 	"github.com/olekukonko/tablewriter"
#)

pkg_deps=(core/glibc)
pkg_build_deps=(core/make core/gcc)

pkg_bin_dirs=(bin)

do_begin() {
  do_default_begin
}

do_download() {
  do_default_download
}

do_verify() {
  do_default_verify
}

do_clean() {
  do_default_clean
}

do_unpack() {
  do_default_unpack
}

do_prepare() {
  do_default_prepare
}

do_build() {
  do_default_build
}

do_check() {
  return 0
}

do_install() {
  export GOFLAGS="-ldflags=-X=main.Version=$pkg_release"
  do_default_install
}

do_strip() {
  do_default_strip
}

do_end() {
  do_default_end
}
