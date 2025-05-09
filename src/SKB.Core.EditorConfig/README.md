# SKB.Core.EditorConfig

This provides a centralized `.editorconfig` for all projects within the scope.

## How to use?
Import the Project Reference to `SKB.Core.EditorConfig`, also allow it to use the private
and Compile plus Build time assets, in any project in the solution.
```xml
<PackageReference Include="SKB.Core.EditorConfig">
	<IncludeAssets>build</IncludeAssets>
	<PrivateAssets>all</PrivateAssets>
</PackageReference>
```

The package adds the ``.editorconfig`` to the location where the `.gitignore` is located.

## Testing
Package the Package to test the file location structure
```shell
dotnet pack . -c Release -o out --no-restore
```
