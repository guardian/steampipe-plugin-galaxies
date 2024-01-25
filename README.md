# Steampipe Plugin Galaxies

A [Steampipe](https://steampipe.io/) plugin to query the [Guardian of the Galaxies](https://galaxies.gutools.co.uk/) data.

## Usage
> [!NOTE]
> This plugin is currently in development and is not yet [released](https://steampipe.io/docs/develop/plugin-release-checklist).

To run this plugin locally:
1. Clone the repository
2. Run the [setup script](script/setup). This will build the plugin binary, and symlink it to the Steampipe home directory.
3. Set the environment variable `GALAXIES_BUCKET`
4. Run a query. For example `steampipe query "select * from galaxies.galaxies_people"`

## Contributing
1. Make a code change
2. Build the plugin `go build -o dist/galaxies.plugin`
3. Run a query
