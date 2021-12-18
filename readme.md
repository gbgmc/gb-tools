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

1. Download the latest release matching your installed version of Ground Branch
from the [releases](https://github.com/JakBaranowski/gb-tools/releases) section,
2. Place the executable in the Ground Branch root directory, by default 
`C:\Program Files (x86)\Steam\steamapps\common\Ground Branch`.

## Usage

### Packaging game modes

If you're working on game modes and want to deliver them in a manner that makes it easy to add them to the game best way is providing a simple zip archive that retains the directory hierarchy of the game. Doing this by hand is repetetive and error prone. You can create such archives automatically using GB Tools `pack` command the command. Here's how:

1. Make sure to install the tool ([instructions](#installation))
2. Create a manifest file at the Ground Branch root directory, by default 
`C:\Program Files (x86)\Steam\steamapps\common\Ground Branch`.
3. Run the `gbt.exe pack "<manifestFilePath>"` command.
4. The packaged game mode should be waiting for you at the Ground Branch root folder.

For more information on manifests see [Manifests](#manifests).

### Installing/Uninstalling the game modes (EXPERIMENTAL)

GB Tools installation and uninstallation commands are created with the sole purpose of making it easier to keep your game modes repository separate from the game installation folder. This is quite helpful since reinstalling the game does not require you to clone and setup the repository again. Moreover the `install` command follows the same logic as the `pack` command, so if you've messed up the game mode manifest then you'll know it first.

GB Tools `install` command moves the files matching game mode manifest rules from the repository to the game installation directory specified in GB tools config. `uninstall` command removes the files matching game mode manifest rules from game installation directory specified in GB tools config.

To install a game mode use the following command

`gbt.exe install "<manifestFilePath>" [gameInstallationPathName]`

where `<manifestFilePath>` is the relative path from the current directory to the game mode manifest file, and `[gameInstallationPathName] is the name of the installation path from the GB Tools config file.

To uninstall a game mode use the following command

`gbt.exe uninstall "<manifestFilePath>" [gameInstallationPathName]`

where `<manifestFilePath>` is the relative path from the current directory to the game mode manifest file, and `[gameInstallationPathName] is the name of the installation path from the GB Tools config file.

### Using non default Ground Branch installation directory

If you've installed Ground Branch in any directory other than `C:/Program Files (x86)/Steam/steamapps/common/Ground Branch` you will need to change the default game path in the GB Tools config to the actual absolute path to Ground Branch installation directory. To do so

1. First, make sure to save GB Tools config with the following command `gbt.exe config`.
2. Open GB Tools config, available under `%AppData%\gbt\gbt.conf`, with any text editor .
3. Change the path parameter of the `default` entry.
4. Save the changes and you're good to go.

For example if you installed Ground Branch to path `D:/Games/Ground Branch` your GB Tools config file should look like this:

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
* **dependencies** is an array of paths to other manifests coontaining files required by this game mode.
* **files** is an array of paths to files used in the game mode. Files paths can use glob patterns.

### Config

GB Tools config can be used to change the installation path used by the `install` and `uninstall` commands.

To create GB Tools config use the `gbt.exe config`. The config will be created at the following path `%AppData%\gbt\gbt.conf`. You can open the config with any text editor.

If no config is present GB Tools will use the default config.

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

#### Game paths

GB Tools can support multiple game paths. The default GB Tools config has two game paths:

* `default` game path `C:/Program Files (x86)/Steam/steamapps/common/Ground Branch`, and
* `CTE` game path `C:/Program Files (x86)/Steam/steamapps/common/GROUND BRANCH CTE`.

You can change the default game path in the GB Tools config file, or you can add additional game paths.

Game paths stored in GB Tools config are only used with the `install` and `uninstall` commands. For more details go to [Installing game modes](#installing-game-modes) or [Uninstalling game modes](#uninstalling-game-modes).

## Contributing

Everyone is more than welcome to fork the project and post pull requests. I'll try to review and merge as soon as possible.

The project uses GoLang 1.17.1.

## Kudos

* @Bob-AT for contribiutions to this project.

## License

This project uses an [MIT license](license.md).
