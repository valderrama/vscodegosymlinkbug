# README.md

VS Code's go extension will claim there are build errors with golang projects that are symlinked under the $GOPATH. However, running `go build` will successfully build the project without any errors.

## Dependencies

Requires:
 - git
 - VS Code command line enabled

```
$> go version
go version go1.10.1 darwin/amd64
$> code --version
1.22.0
6b4d53cdab8bcae1eaaa4934d93c077319b573db
x64
$> code --list-extensions --show-versions
PeterJausovec.vscode-docker@0.0.26
alefragnani.Bookmarks@0.19.0
eamodio.gitlens@8.2.4
lukehoban.Go@0.6.78
ms-python.python@2018.3.1
ms-vscode.cpptools@0.16.1
redhat.java@0.23.0
streetsidesoftware.code-spell-checker@1.6.9
```

## Repro Steps

Run the following commands clone the repo, symlink it, and open a directory which symlinks to the package in VS Code:
```
cd ~
mkdir valderrama
cd valderrama
git clone git@github.com:valderrama/vscodegosymlinkbug.git

cd $GOPATH/src
mkdir valderrama
cd valderrama
ln -s ~/valderrama/vscodegosymlinkbug vscodegosymlinkbug

code -n $GOPATH/src/valderrama/
```

Then Open vscodegosymlinkbug/main.go. You will need to wait for ~30s for GO code extension to build project then you will see the following error appear:
![Bug Screen Shot](https://raw.githubusercontent.com/valderrama/vscodegosymlinkbug/master/bug_screen_shot.png)

If you run the following commands to build the package following the symlink it will be successful:
```
cd $GOPATH/src/valderrama/vscodegosymlinkbug
go build
```

Ex.
```
$> cd $GOPATH/src/valderrama/vscodegosymlinkbug
$> go build
$> 
```

If you run the following commands to build the package from the source directory it will not be successful with the same error given by VS Code:
```
cd ~/valderrama/vscodegosymlinkbug
go build
```

Ex.
```
$> cd ~/valderrama/vscodegosymlinkbug
$> go build
# _/Users/alex/valderrama/vscodegosymlinkbug
./main.go:19:20: cannot use connection (type *"github.com/hyperledger/fabric/vendor/google.golang.org/grpc".ClientConn) as type *"google.golang.org/grpc".ClientConn in assignment
$>
```

## Expected Behavior

When VS Code opens a project directory in the $GOPATH with symlinks to directories outside the $GOPATH it should build code as if it were accessing the directories from the project directory not as if it were in the symlink target directory.