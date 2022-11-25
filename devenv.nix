{ pkgs, ... }:

{
  packages = [ pkgs.git ];

  languages.nix.enable = true;
  languages.go.enable = true;
  languages.javascript.enable = true;

  processes = {
    server.exec = "live-server build";
  };
}
