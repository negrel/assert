{
  inputs = { flake-utils.url = "github:numtide/flake-utils"; };

  outputs = { nixpkgs, flake-utils, ... }:
    let
      outputsWithoutSystem = { };
      outputsWithSystem = flake-utils.lib.eachDefaultSystem (system:
        let
          pkgs = import nixpkgs { inherit system; };
          #lib = pkgs.lib;
        in {
          devShells = {
            default = pkgs.mkShell {
              buildInputs = with pkgs; [ go gopls gotools shellcheck ];
            };
          };
        });
    in outputsWithSystem // outputsWithoutSystem;
}
