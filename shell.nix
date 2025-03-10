# npin: https://github.com/NixOS/nixpkgs/commit/4d9e4457f8e83120c9fdf6f1707ed0bc603e5ac9 - pined to latest successful build on hydra
{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/2d9e4457f8e83120c9fdf6f1707ed0bc603e5ac9.tar.gz") { } }:

pkgs.mkShell {
  packages = with  pkgs;[
    go_1_24
    gotools
    golangci-lint
  ];
}
