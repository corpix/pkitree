with import <nixpkgs> {};
stdenv.mkDerivation {
  name = "pkitree-shell";
  buildInputs = [
    go
    jq
  ];
  shellHook = ''
    unset GOPATH
  '';
}
