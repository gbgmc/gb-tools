# Ground Branch Tools

## Description

This is a set of tools that aim to automate the repetitive work required when working on
Ground Branch game modes.

## Table of contents:

1. [Installation](#installation)
2. [Usage](#usage)
3. [Contributing](#contributing)
4. [Kudos](#kudos)
5. [License](#license)

## Installation

To install the tools:

1. Download the latest release matching your installed version of Ground Branch from the [releases](https://github.com/JakBaranowski/gb-tools/releases) section,
2. Place the executable in your Ground Branch game modes directory.

## Usage

### Packaging game modes

Best way of delivering your game modes to the users is by providing a simple zip archive that retains the files hierarchy of the game. Preparing a package like this by hand is repetetive and error prone. GB Tools can automate this process for you with the `pack` command. Here's how to use it:

```
gbt.exe pack "<manifestFilePath>"
```

where:

* `"<manifestFilePath>"` is the relative path from the current directory to the game mode manifest file.

For more information on manifests see [Manifests](#manifests).

### Installing/Uninstalling the game modes (EXPERIMENTAL)

GB Tools installation and uninstallation commands are created with the sole purpose of making it easier to keep your game modes repository separate from the game installation folder. This is quite helpful since reinstalling the game does not require you to clone and setup the repository again. Moreover the `install` command follows the same logic as the `pack` command, so if you've messed up the game mode manifest then you'll know it first.

GB Tools `install` command moves the files, matching game mode manifest rules, from the repository to the game installation directory specified in GB tools config. To install a game mode use the following command:

```
gbt.exe install "<manifestFilePath>" [gameInstallationPathName]
```

where:

* `<manifestFilePath>` is the relative path from the current directory to the game mode manifest file, and 
* `[gameInstallationPathName]` is an optional argument specifying the name of the installation path from the GB Tools config file. By default it's `default`.

`uninstall` command removes the files, matching game mode manifest rules, from game installation directory specified in GB tools config. To uninstall a game mode use the following command:

```
gbt.exe uninstall "<manifestFilePath>" [gameInstallationPathName]
```

where:

* `<manifestFilePath>` is the relative path from the current directory to the game mode manifest file, and 
* `[gameInstallationPathName]` is an optional argument specifying the name of the installation path from the GB Tools config file. By default it's `default`.

### Using non default Ground Branch installation directory

If you've installed Ground Branch in any directory other than `C:/Program Files (x86)/Steam/steamapps/common/Ground Branch` you will need to change the default game path in the GB Tools config to the actual absolute path to Ground Branch installation directory.

1. First, make sure to save GB Tools config with the following command `gbt.exe config`.
2. Open GB Tools config, available under `%AppData%\gbt\gbt.conf`, with any text editor.
3. Change the path parameter of the `default` entry.
4. Save the changes.

For example, if you installed Ground Branch to path `D:/Games/Ground Branch` your GB Tools config file should look like this:

```json
{
    "gamePaths": [
        {
            "name": "default",
            "path": "D:/Games/Ground Branch"
        },
        {
            "name": "CTE",
            "path": "C:/Program Files (x86)/Steam/steamapps/common/GROUND BRANCH CTE"
        }
    ]
}
```

Alternatively, you can add an extra path to the config file.

### Manifests

Game mode manifest files can have any extension, but has to follow json formatting. See example below:

```json 
{
    "name" : "BreakOut",
    "version" : "0.3.22",
    "dependencies" : [
        "Common.gbm"
    ],
    "files" : [
        "GroundBranch/Content/GroundBranch/Mission/*/BreakOut.mis"
        "GroundBranch/Content/GroundBranch/Lua/BreakOut.lua",
        "GroundBranch/Content/Localization/GroundBranch/en/BreakOut.csv",
        "GroundBranch/Content/GroundBranch/DefaultLoadouts/BreakOut.kit"
    ]
}
```

* **name** can be any string.
* **version** is not really used atm.
* **dependencies** is an array of paths to other manifests containing files required by the game mode.
* **files** is an array of paths to files used in the game mode. Files paths can use glob patterns.

### Config

To create GB Tools config use the `gbt.exe config`. The config will be created at the following path `%AppData%\gbt\gbt.conf`. You can open the config with any text editor.

GB Tools config can be used to change the installation path used by the `install` and `uninstall` commands. If no config is present GB Tools will use the default config. The default GB Tools config has two game paths:

* `default` game path `C:/Program Files (x86)/Steam/steamapps/common/Ground Branch`, and
* `CTE` game path `C:/Program Files (x86)/Steam/steamapps/common/GROUND BRANCH CTE`.

```json
{
    "gamePaths": [
        {
            "name": "default",
            "path": "C:/Program Files (x86)/Steam/steamapps/common/Ground Branch"
        },
        {
            "name": "CTE",
            "path": "C:/Program Files (x86)/Steam/steamapps/common/GROUND BRANCH CTE"
        }
    ]
}
```

You can change the default game path in the GB Tools config file, or you can add additional game paths.

## Contributing

Everyone is more than welcome to fork the project and post pull requests. I'll try to review and merge as soon as possible.

The project uses GoLang 1.17.5.

## Kudos

* @Bob-AT for contribiutions to this project.

## License

This project uses an [MIT license](license.md).
