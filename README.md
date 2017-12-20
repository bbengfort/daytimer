# Daytimer

**Console program for interaction with the [Google Calendar API](https://developers.google.com/google-apps/calendar/v3/reference/)**

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

Now you can view your agenda and add and remove calendar events from the command line.

    $ daytimer agenda

Note that this app is for simple use of your primary Google calendar from the command line, so you need to setup a Google account if you don't have one in order to use the app. Finally, this app is not a substitute for any other calendar app, it is meant to be a quick helper on the command line.
