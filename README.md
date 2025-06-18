# SKB.Core.ProjectFixtures

## Theory
ProjectFixtures are part of a toolchain which is commonly referred within the organisations as 
`Core.Tools`. Core tools allow us to have a more centralized approach over organisation specific
software maintenance.

For more detailed insight in the idea, you can refer to the **[Theory](https://github.com/saapo-ka-baadshah/Theory.Core/blob/main/Organisational/SoftwareMaintainance/CoreTools/IMPORTANCE.md)**

## General Overview

Here we specify all the necessary Fixtures for the Projects within following scopes:
1. Shared `LICENSE` zones: 
    - Here we expect the Projects to follow a common software license
    - Common ``CODE_OF_CONDUCT`` (If Necessary)
2. Shared `.editorconfig`.
   These are the scopes for the shared coding style guidlines

## How to Use:
Include these files in the head/ face end of your solution/ project.
Add references as followed:
```xml
<!--For SKB.Core.LicenseConfig-->
<PackageReference Include="SKB.Core.LicenseConfig">
	<IncludeAssets>build</IncludeAssets>
	<PrivateAssets>all</PrivateAssets>
</PackageReference>
        
<!--For SKB.Core.EditorConfig-->
<PackageReference Include="SKB.Core.EditorConfig">
   <IncludeAssets>build</IncludeAssets>
   <PrivateAssets>all</PrivateAssets>
</PackageReference>
```