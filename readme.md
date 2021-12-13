# Ground Branch Tools

## Description

This is a set of tools that hopefully will help with developing game modes for Ground Branch.

## Table of contents:

1. [Installation](#installation)
2. [Usage](#usage)
3. [Contributing](#contributing)
4. [Credits](#credits)
5. [License](#license)

## Installation

To install the tools:

1. Download the latest release matching your installed version of Ground Branch
from the [releases](https://github.com/JakBaranowski/gb-tools/releases) section,
2. Place the executable in the Ground Branch root directory, by default 
`C:\Program Files (x86)\Steam\steamapps\common\Ground Branch`.

## Usage

### Mirroring AI loadouts

Currently, due to a bug, Ground Branch mission editor will use invalid paths when selecting AI loadouts (it will use folder `AdGuys` instead of `BadGuys`). As it turns out this can be resolved by creating a folder with one extra letter in front of the folder name that will mirror the contents of the original AI loadout folders. So if you would like to add your own loadout for AI to use, you would have to place the file in both the original and mirrored directory. This tool can take care of the repetetive process. 

1. Make sure to install the tool ([instructions](#installation))
2. Create your loadout file,
2. Place the loadout file in the original AI loadout folder, by default 
`C:\Program Files (x86)\Steam\steamapps\common\Ground Branch\GroundBranch\Content\GroundBranch\AI\Loadouts\BadGuys`
3. Open Command Line in Ground Branch installation folder,
4. Run `gbt.exe loadout` command in the command line prompt.

Note: this will also remove files in the `_BadGuys` folder if they don't have a counterpart in the original `BadGuys` folder.

Alternativly, if you use Visual Studio Code to edit the AI loadouts you can use an extension that will run a command whenever you save a file with a mathing path, e.g.: [Save and Run by wk-j](https://marketplace.visualstudio.com/items?itemName=wk-j.save-and-run) and configure it to run the `gbt.exe loadout` command whenever a file is saved to the original AI loadout folder.

### Packaging game modes

If you're working on game modes and want to deliver them in a manner that will make it easy to "install" them. Best way is providing a simple zip archive that can be unpacked in the game root folder by the "end user". You can create  such archives using the `gbt.exe gamemode pack <manifestFilePath>` command. Here's how:

1. Make sure to install the tool ([instructions](#installation))
2. Create a manifest file at the Ground Branch root directory, by default 
`C:\Program Files (x86)\Steam\steamapps\common\Ground Branch`.
3. Run the `gbt.exe pack "<manifestFilePath>"` command.
4. The packaged game mode should be waiting for you at the Ground Branch root folder.

For more information on manifests see [Manifests](#manifests).

### Config

Ground Branch tools config can be used to set up

* which AI loadouts directories Ground Branch Tools should mirror when the `loadout` command is run, and
* change the installation path used by the `install` and `uninstall` commands.

To create Ground Branch Tools config use the `gbt.exe config`. The config will be created at the following path `%AppData%\gbt\gbt.conf`. You can open the config with any text editor.

If no config is present Ground Branch Tools will use the default config.

#### Game path

If you've installed Ground Branch in any directory different than `C:/Program Files (x86)/Steam/steamapps/common/Ground Branch` you will need to change the value of the `GamePath` parameter in the Ground Branch Tools config with the actual path to Ground Branch installation directory.

#### Loadouts

To mirror additional directories containing AI loadout files, add entries for each extra directory to the `Loadouts` array. Each entry in `Loadouts` array has to have `Name`, `SourceRelativePath`, and `DestinationRelativePath`.

For example to mirror loadout files from `GroundBranch/Content/GroundBranch/AI/Loadouts/Example` directory in `GroundBranch/Content/GroundBranch/AI/Loadouts/_Example` directory add the following entry to the `Loadouts` array:

```json
{
  "Name": "Example",
  "SourceRelativePath": "GroundBranch/Content/GroundBranch/AI/Loadouts/Example",
  "DestinationRelativePath": "GroundBranch/Content/GroundBranch/AI/Loadouts/_Example"
}
```

In the end the config would look like this:

```json
{
  "GamePath": "C:/Program Files (x86)/Steam/steamapps/common/Ground Branch",
  "Loadouts": [
    {
      "Name": "BadGuys",
      "SourceRelativePath": "GroundBranch/Content/GroundBranch/AI/Loadouts/BadGuys",
      "DestinationRelativePath": "GroundBranch/Content/GroundBranch/AI/Loadouts/_BadGuys"
    },
    {
      "Name": "Example",
      "SourceRelativePath": "GroundBranch/Content/GroundBranch/AI/Loadouts/Example",
      "DestinationRelativePath": "GroundBranch/Content/GroundBranch/AI/Loadouts/_Example"
    }
  ]
}
```

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

## Contributing

Everyone is more than welcome to fork the project and post pull requests. I'll try to review and merge as soon as possible.

The project uses GoLang 1.17.1.

## License

This project uses an [MIT license](license.md).
