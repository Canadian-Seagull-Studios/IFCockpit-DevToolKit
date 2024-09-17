# IFCockpit Development Toolkit

This toolkit supports third-party developers in developing cockpit panels for the IFCockpit application from Canadian Seagull Studios Ltd. This toolkit allows developers to test their panels before submitting them to Canadian Seagull Studios for review and inclusion in IFCockpit.

This is not open source and may only be used by developers who are part of the IFCockpit developer programme under the terms of the programme and may not be shared with anyone outside the terms of that programme or otherwise with the explicit written consent of Canadian Seagull Studios. Contact Canadian Seagull Studios at info@canadianseagulls.com for mor einformation on the developer programme.

Copyright 2024, Canadian Seagull Studios Ltd. All rights reserved.

## Requirements

IFCockpit is a macOS desktop application built with the cross-platform Wails framework for building desktop apps using Web technology. IFCockpit uses Svelte as as its Web framework within Wails.

To use this toolkit to build a cockpit panel requires:

1. A Mac computer running macOS Sonoma or Sequoia

2. Familiarity build Web applications using Svelte along with HTML, CSS and JavaScript

3. XCode command intalled along with its command line tools

4. Go, Wails and NPM installed

## Installing Dependencies: Preparing Your Mac to use the Toolkit

Before downloading the toolkit, several steps are needed to set up your Mac:

1. Install Go

2. Install NPM

3. Install the XCode command-line tools

4. Install Wails

### Installing Go

Download Go for your Mac from the official download page: https://go.dev/dl/.

Once downloaded, follow the installation instructions here: https://go.dev/doc/install.

Once installed, start a new terminal window and verify that your `PATH` includes `~/go/bin` with this command:

```
echo $PATH | grep go/bin
```

### Installing NPM

NPM is the package manager which comes with Node.js. You can download Node (including NPM) if you need it here: https://nodejs.org/en/download/package-manager.

### Installing the XCode Command-Line Tools

First make sure you have the latest version of XCode installed on your Mac. You can find XCode in the macOS App Store.

You can then install the command-line tools with this command:

```
xcode-select --install
```

### Installing Wails

Once you have Go, NPM and the XCode command-line tools installed you can install Wails with this command:

```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

For reference, the complete Wails installation instructions can be found here: https://wails.io/docs/gettingstarted/installation. These instructions include the steps above for installing Go, NPM and the XCode command-line tools.

Once you have Wails installed, you can verify it is installed with the minimum dependencies with this command:

```
wails doctor
```

This should output a verification report which completes with the following message:

```
SUCCESS  Your system is ready for Wails development!
```

## Setup

Once the dependencies are successfully installed, clone the IFCockpit Developers Toolkit Github repository:

```
git clone git@github.com:likeablegeek/ifcockpit-devtoolkit.git
```

Once you have clones the repository, change into the `frontend` subdirectory in the repository and install package dependencies with NPM:

```
npm i
```

## Using the Toolkit

To run the toolkit application to test your panel, use the following command in the root directory of the respository:

```
wails dev
```

This should build the application and launch it.

By default, the toolkit includes the default general aviation cockpit panel from IFCockpit as a reference and to allow you to test your installation of the toolkit is working before beginning your own panel development.

To make sure the toolkit is working, first make sure you enable the IF Connect API by selecting General > Enable Infinite Flight Connect in the Infinite Flight settings. Next, check the IP address of your Infinite Flight device.

Once done, simply enter that IP address in the IFCockpit Developer Toolkit app and click connect. It should indicate a successful connection and the panel instruments should respond to changes in your flights altitude, airspeed, vertical speed and so on.