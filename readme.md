# Ground Branch Tools

## Description

This is a set of tools that hopefully will help with developing game modes for
Ground Branch. Currently, the tool helps resolve two issues:

* The need to mirror the contents of the AI loadout file in order for the Ground
Branch mission editor to see them.
* Invalid paths to AI classes in the mission files.

## Table of contents:

1. [Installation](#installation)
2. [Usage](#usage)
3. [Contributing](#contributing)
4. [Credits](#credits)
5. [License](#license)

## Installation

To install the tools:

1. Download the latest release from the [releases](/releases) section,
2. Unpack the contentes of the downloaded archive to Ground Branch root 
directory, by default `C:\Program Files (x86)\Steam\steamapps\common\Ground Branch`.

## Usage

### Mirroring AI loadouts

Currently, due to a bug, Ground Branch mission editor will use invalid paths when
selecting AI loadouts (it will use folder `AdGuys` instead of `BadGuys`). As it 
turns out this can be resolved by creating a folder with one extra letter in front
of the folder name that will mirror the contents of the original AI loadout folders.
So if you would like to add your own loadout for AI to use, you would have to place
the file in both the original and mirrored directory. This tool can take care of
the repetetive process. 

1. Make sure to install the tool ([instructions](#installation))
2. Create your loadout file,
2. Place the loadout file in the original AI loadout folder, by default 
`C:\Program Files (x86)\Steam\steamapps\common\Ground Branch\GroundBranch\Content\GroundBranch\AI\Loadouts\BadGuys`
3. Open Command Line in Ground Branch installation folder,
4. Run `gbt mirror` command in the command line prompt.

Note: this will also remove files in the `_BadGuys` folder if they don't have a 
counterpart in the original `BadGuys` folder.

Alternativly, if you use Visual Studio Code to edit the AI loadouts you use an
extension that will run a command whenever you save a file with a mathing path,
e.g.: [Run on save](https://marketplace.visualstudio.com/items?itemName=emeraldwalk.RunOnSave)
and configure it to run the `gbt mirror` command whenever a file is saved to the 
original AI loadout folder. 

### Fixing invalid class paths in mission files

Due to another bug after selecting an AI class in the Ground Branch mission editor
the game will search for the class files under an invalid path. After some digging
it turns out that this can be fixed by editing the .mis file. Doing this manually
takes a lot of time, so you can automate the process using GB Tools:

1. Make sure to install the tool ([instructions](#installation))
2. Create your mission files and set up spawn points, select the classes you want
to use (they will not work at this point).
3. Open Command Line in Ground Branch installation folder,
4. Run `gbt mission <path to mission file>` command in the command line prompt. 
E.g.: `gbt mission GroundBranch\Content\GroundBranch\Mission\Arena\example.mis`.

Note: this is kind of a hacky way to resolve the issue. It did work fine with multiple
mission files I created, but it may not work in some cases. 

The tool will create a backup of the original mission file with the `.mis_backup` 
extension.

## Contributing

Everyone is more than welcome to fork the project and post pull requests. I'll try
to review and merge as soon as possible.

The project uses GoLang 1.16.4.

The versioning will follow the versioning of Ground Branch, but will add minor 
version number for possible tool updates within one Ground Branch release cycle.
E.g.: first release of the tool for Ground Branch version `1032.3` will have version
 `1032.3.1`.

## Kudos

* **BlackFoot** Studios for creating Ground Branch.
* **tjl** for creating this awesome 
[Guide](https://steamcommunity.com/sharedfiles/filedetails/?id=2461956424).
* **AV** for creating this [Video Tutorial](https://www.youtube.com/playlist?list=PLle5osICJhZJwHxGOb1iBXoyu_uk9yXMY).

## License

This project uses an [MIT license](license.md).
