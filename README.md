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

## Requirements
-  .NET 9 

Also, this package itself is served to be used as a requirement. 

## Installation / Incorporation in an organisation
The repository is equipped with GitHub actions that allow us to use a GitHub nuget registry.

To add these packages to your organisation wide GitHub NuGet registry:
1. Fork the Repository
2. Add following variables in the repository variables and secrets section for GitHub actions 

   | Variables                 | Values                        |
   |---------------------------|-------------------------------|
   | USERNAME                  | gh username for the packages  |
   | NUGET_SRC_KEY             | By what name nuget source identifies |
   | MAJOR_VERSION             | Incremented with non-reversible (incompatible) lib changes |
   | MINOR_VERSION             | Incremneted with feature adds |

   | Secrets                   | Values                        |
   |---------------------------|-------------------------------|
   | GH_PATH                   | PAT with clearances for Package uploads|

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