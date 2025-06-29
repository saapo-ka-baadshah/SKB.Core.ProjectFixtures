# SKB.Core.EditorConfig

This provides a centralized `.editorconfig` for all projects within the scope.

## How to use?

Current support for the `.editorconfig` loaders is for the following:

1. `.NET`
2. `go`

### `.NET`

Import the Project Reference to `SKB.Core.EditorConfig`, also allow it to use the private
and Compile plus Build time assets, in any project in the solution.

```xml
<PackageReference Include="SKB.Core.EditorConfig">
	<IncludeAssets>build</IncludeAssets>
	<PrivateAssets>all</PrivateAssets>
</PackageReference>
```

The package adds the `.editorconfig` to the location where the `.gitignore` is located.

#### Testing

Package the Package to test the file location structure

```shell
dotnet pack . -c Release -o out --no-restore
```

### `Go`

Import the project reference to the `main` package in your go project while using the Go-directive `//go:generate` for generating the set `.editorconfig`.

#### 1. Install the package

```shell
go get github.com/saapo-ka-baadshah/SKB.Core.ProjectFixtures/src/SKB.Core.EditorConfig
```

#### 2. Add the `go:generate` directive call to your project, mainly at the top of your main package

```go
// <Project Description>

//go:generate go run github.com/saapo-ka-baadshah/SKB.Core.ProjectFixtures/src/SKB.Core.EditorConfig

package main

.
.
.
```

#### 3. Add this `go generate` step in your project targets in the `Makefile`

```
go-generate:
	go generate

go-build:
	go build

all: go-generate go-build
	echo "GO: Building all..."
```
