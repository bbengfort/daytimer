# Daytimer

**Console program for interaction with the [Google Calendar API](https://developers.google.com/google-apps/calendar/v3/reference/)**

Daytimer provides a series of utilities for querying multiple Google Calendars from the command line. Most notably it can generate agendas and send them via email; a task which I have cron'd so I get a daily update about my daily schedule.

This tool is set up for personal use right now, and as such it doesn't necessarily adhere to code practices that make it generally available to others. If you'd like to use Daytimer, feel free, but note that you'll have to jump through some hoops to get things configured and set up!

## Getting Started

First download/clone the code and compile it using `go get` (this will automatically install the tool into your `$GOPATH`):

    $ go get github.com/bbengfort/daytimer/...

Next, you'll have to use the Google Developer Console to turn on the Calendar API and to create credentials for your version of the Daytimer application. You can use the [Calendar API wizard](https://console.developers.google.com/start/api?id=calendar) to quickly set up an app and an OAuth consent screen. Create OAuth client ID credentials and download them to a file called `client_secret.json`. For more on this, see Step 1 of the [Google Calendar Quickstart](https://developers.google.com/google-apps/calendar/quickstart/go).

The daytimer program is configured in a hidden directory in your home directory. Create this folder and move the `client_secret.json` file there.

    $ mkdir ~/.daytimer
    $ mv client_secret.json ~/.daytimer/

Next, authorize the daytimer application as follows:

    $ daytimer auth

This will open a web browser, prompt you to login with your Google account, then provide an authentication code that you must copy and paste into the command line. Your token will be saved and only needs to be provided once, however if you're having trouble accessing the service, you can reauthorize with the `auth` command.

### Usage

Now you can view your primary calendar's agenda:

    $ daytimer agenda

However, one calendar isn't probably enough given that you have work calendars, birthday calendars, school calendars, etc. To list all available calendars in your account:

    $ daytimer calendars --list

You can then activate calendars as follows:

    $ daytimer calendars -a calid1,calid2

Where `calid` is a calendar id, which looks like an email address (and usually is an email address) specified in parentheses in the list output. To show all currently activate calendars:

    $ daytimer calendars

Your primary calendar will be marked with a ☆! If you prefer, you can also add each calendar id on its own line in `~/.daytimer/calendars.txt` or by running `daytimer calendars -e` (more on this in the configuration section). Now the agenda command will list all events for all active calendars.

If you want to see the next 10 upcoming events for a specific calendar (your primary calendar by default) you do the following:

    $ daytimer upcoming -c calid

This will list the next `-n N` (default 10) events on that calendar, no matter how far in the future they will occur.

### Configuration

There are a variety of ways to edit configuration, but all config files are stored in `~/.daytimer/` (note that this means you must run daytimer as a user with a home directory). As a helper, some commands have a `-e` flag to directly edit their configurations. For example:

    $ daytimer config -e

Will open the configuration YAML file for editing. Note that this is fraught with peril, if you make a mistake or save bad YAML, then things won't work the next time and you'll have to edit the config manually!

By default the editor that's chosen is searched for in the path or fetched via the `$EDITOR` environment variable. Additionally you can specify an editor in the config file itself.

You can also edit the calendars list using:

    $ daytimer calendars -e

Write each calendar id (the associated email address) on its own line. Lines that start with `#` are ignored.

### Email

If you would like to email your agenda to someone, you'll first need to setup SMTP. If you're using Gmail or Google Apps, make sure that you set your email account [to allow less secure apps](https://support.google.com/accounts/answer/6010255?hl=en).

Edit the configuration file with `daytimer config -e` or edit `~/daytimer/config.json` with your email credentials:

```yaml
# Add the SMTP configuration to send email agendas
email:
  host: "smtp.gmail.com",
  port: 587,
  user: "me@gmail.com",
  password: "supersecretsquirrel"
  use_tls: true
```

This will (hopefully) be all that's needed to send yourself an HTML email with the agenda printed out:

    $ daytimer agenda -e jane@example.com

As with anything related to email, this is not the nicest HTML in the world (it uses 90s style HTML tables) but hopefully it serves its purpose!

## Development

Just a few notes for development, feel free to ignore these unless you're planning on contributing, or you're working in your own fork!

### Templates

Templates for the agenda HTML and email text format are stored in the `templates` directory, but are compiled with the binary using [`go-bindata`](https://github.com/jteeuwen/go-bindata). To recompile the templates into the `bindata.go` file use the following:

    $ go-bindata -pkg daytimer templates/...

Note that you can also use the `-debug` flag to recompile the templates into a form that are loaded from disk so that you can work on them. Just remember to recompile without the debug flag before pushing to GitHub!

### Tests

Currently there are no tests. This will change. When it changes, we'll probably be using [ginkgo](http://onsi.github.io/ginkgo/) and [gomega](http://onsi.github.io/gomega/).

### Releases

Releases are managed using [GoReleaser](https://goreleaser.com/). To create a release ensure the following have been performed:

1. Tests pass
2. Template assets have been compiled
3. Version has been bumped
4. Godep save has been run

Commit the latest version and push to GitHub, then create a tag:

    $ go tag -a vx.x -m "tag name"
    $ git push origin vx.x

Where `vx.x` is the semantic version number. Then run the release:

    $ gorelease

Note that the environment variable `GITHUB_TOKEN` is required. You can then go to the [release page of the repository](https://github.com/bbengfort/daytimer/releases) and update the release info for the tag if needed.
