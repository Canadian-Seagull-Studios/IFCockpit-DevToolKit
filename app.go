/*

	This file is part of the IFCockpit Developer Toolkit
	from Canadian Seagull Studios Ltd.

	This is not open source and may only be used by developers
	who are part of the IFCockpit developer programme under the
	terms of the programme and may not be shared with anyone
	outside the terms of that programme or otherwise without the
	explicit written consent of Canadian Seagull Studios Ltd.

	Copyright 2024, Canadian Seagull Studios Ltd.
	All rights reserved.

*/

package main

import (
	"bufio"
	"context"
	"io"
	"math"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"encoding/binary"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Command struct { // Representation of commands from the IF manifest
	code    int
	name    string
	cmdType int
}

type State struct { // Representation of a state value from IF
	Name string `json:"name"`
	Type any    `json:"type"`
	Val  any    `json:"val"`
}

type Device struct { // Representation of an IF device as returned by UDP
	Addresses  []string `json:"addresses"`
	Port       int      `json:"port"`
	DeviceID   string   `json:"deviceID"`
	Version    string   `json:"version"`
	DeviceName string   `json:"deviceName"`
	State      string   `json:"state"`
	Aircraft   string   `json:"aircraft"`
	Livery     string   `json:"livery"`
}

var c net.Conn  // TCP client connect
var err error   // Errors
var r io.Reader // Buffer reader for responses

var manifestCmd = []byte{0xff, 0xff, 0xff, 0xff, 0x00} // Default command to fetch manifest

var manifest string      // String to hold raw manifst
var commandList []string // Raw manifest split by line into array of strings

var pollStates = make(map[string]int) // List of states to poll
var states = make(map[string]any)     // Latest state data for all polled states

var CommandsByName = make(map[string]Command) // Map of all commands organised by command name
var CommandsByCode = make(map[int]Command)    // Map of all commands organised by command code

var device Device // Current active IF device

var delayFactor time.Duration = 2  // 1 = fastest, 3 = slowest
var queueDelay time.Duration = 250 // 1 = fastest, 3 = slowest
var pollDelay time.Duration = 10   // 25 ms delay before refetching a state from IF -- assumes 10 states x 3 times per second
var stickDelay uint32 = 50         // 50 ms between each stick state update to IF -- assumes 20 times per second

var envInfo runtime.EnvironmentInfo

var isIFConnected = false // Is IF connected?

// Queue a state for subsequent polling by adding to pollStates
func queue(state string) {

	pollStates[state] = CommandsByName[state].code

}

// Deqeueue a state for polling by removing from pollStates
func dequeue(state string) {

	delete(pollStates, state)

}

// Fetch a single state from IF
func fetch(a *App, state string) (int, error) {

	// Get integer command/state code from the command/state name
	cmdCode := uint32(CommandsByName[state].code)

	// Convert command code to bytes in little endian per API spec
	cmdBytes := make([]byte, 5)
	binary.LittleEndian.PutUint32(cmdBytes, cmdCode)
	cmdBytes[4] = 0x00

	// Send command to IF
	_, err = c.Write(cmdBytes)
	if err != nil {
		return 0, err
	}

	return 1, nil

}

// Handle the incoming buffer and when a state is received refetch the state
func poll(a *App) error {

	// Iterate reading states from buffer
	for {

		// Get command details based on first byte
		cmd := make([]byte, 4)
		if _, err = io.ReadFull(r, cmd); err != nil {
			runtime.EventsEmit(a.ctx, "errormsg", "Could not read command from poll")
			return err
		}
		cmdCode := int(binary.LittleEndian.Uint32(cmd))    // convert to integer based on little endian
		cmdName := CommandsByCode[cmdCode].name            // get the command name based on the code
		cmdType := uint32(CommandsByCode[cmdCode].cmdType) // get the type of state (i.e. string, boolean, etc)

		// Get length of data returned based on second byte
		len := make([]byte, 4)
		if _, err = io.ReadFull(r, len); err != nil {
			runtime.EventsEmit(a.ctx, "errormsg", "Could not read length from poll")
			return err
		}
		reslen := binary.LittleEndian.Uint32(len)

		// Get reponse data based on the length
		res := make([]byte, reslen)
		if _, err = io.ReadFull(r, res); err != nil {
			runtime.EventsEmit(a.ctx, "errormsg", "Could not read data from poll")
			return err
		}

		if cmdType == 0 { // Boolean
			states[cmdName] = (res[0] == 1)
		}

		if cmdType == 1 { // 32-bit integer
			result := binary.LittleEndian.Uint32(res)
			states[cmdName] = result
		}

		if cmdType == 2 { // 32-bit float
			result := math.Float32frombits(binary.LittleEndian.Uint32(res))
			states[cmdName] = result
		}

		if cmdType == 3 { // 64-bit double
			result := math.Float64frombits(binary.LittleEndian.Uint64(res))
			states[cmdName] = result
		}

		if cmdType == 4 { // String
			states[cmdName] = string(res[4:])
		}

		if cmdType == 5 { // 64-bit integer
			result := binary.LittleEndian.Uint64(res)
			states[cmdName] = result
		}

		// Update front end of the new state data by emitting event
		runtime.EventsEmit(a.ctx, "statechange", State{cmdName, cmdType, states[cmdName]})

		// Pause before fetching the state again
		time.Sleep(pollDelay * delayFactor * time.Millisecond)

		// Fetch the state again
		fetch(a, cmdName)

	}

	return nil

}

// Front-end broker function to initialise IF connection
func (a *App) IFCinit(stateList []string, ip string) error {

	runtime.EventsEmit(a.ctx, "statusmsg", "Connecting to IF")

	// Connect to IF
	c, err = net.Dial("tcp", ip+":10112") // Arman's iPad
	if err != nil {
		runtime.EventsEmit(a.ctx, "errormsg", "Could not connect")
		return err
	}

	runtime.EventsEmit(a.ctx, "statusmsg", "Sending Manifest command")

	// Send Manifest command
	_, err = c.Write(manifestCmd)
	if err != nil {
		runtime.EventsEmit(a.ctx, "errormsg", "Could not send manifest command")
		return err
	}

	// Set up read buffer
	r = bufio.NewReader(c)

	runtime.EventsEmit(a.ctx, "statusmsg", "Receiving manifest command")

	// Fetch command from byte stream
	cmd := make([]byte, 4)
	if _, err = io.ReadFull(r, cmd); err != nil {
		runtime.EventsEmit(a.ctx, "errormsg", "Could not fetch manifest command from buffer")
		return err
	}

	runtime.EventsEmit(a.ctx, "statusmsg", "Receiving manifest length")

	// Fetch manifest length from byte stream
	len := make([]byte, 4)
	if _, err = io.ReadFull(r, len); err != nil {
		runtime.EventsEmit(a.ctx, "errormsg", "Could not fetch manifest length from buffer")
		return err
	}

	// Fetch manifest string length from byte stream
	strlen := make([]byte, 4)
	if _, err = io.ReadFull(r, strlen); err != nil {
		runtime.EventsEmit(a.ctx, "errormsg", "Could not fetch manifest string length from buffer")
		return err
	}
	manlen := binary.LittleEndian.Uint32(strlen)

	runtime.EventsEmit(a.ctx, "statusmsg", "Receiving manifest data")

	// Fetch manifest from byte stream
	man := make([]byte, manlen)
	if _, err = io.ReadFull(r, man); err != nil {
		runtime.EventsEmit(a.ctx, "errormsg", "Could not fetch manifest from buffer")
		return err
	}

	runtime.EventsEmit(a.ctx, "statusmsg", "Processing manifest")

	// Convert manifest to string and split into lines
	manifest := string(man)
	commandList := strings.Split(manifest, "\n")

	// Set up regex to strip leading non-numbers
	r := regexp.MustCompile("^[^0-9]*")

	// Iterate lines of manifests
	for _, command := range commandList {

		// Split line into component parts
		commandParts := strings.Split(command, ",")

		if cap(commandParts) == 3 { // cap here gives us number of elements in array

			var commandCode int
			var commandName string
			var commandType int

			// Get command code
			code := strings.Replace(commandParts[0], "\xe3", "", -1)
			code = strings.Replace(code, "\x00", "", -1)
			codeStripped := r.ReplaceAllString(code, "")

			commandCode, err = strconv.Atoi(codeStripped)
			if err != nil {
				runtime.EventsEmit(a.ctx, "errormsg", "Invalid Command Code")
				return err
			}

			// Get command type
			cmdType := strings.Replace(commandParts[1], "\xe3", "", -1)
			cmdType = strings.Replace(cmdType, "\x00", "", -1)
			commandType, err = strconv.Atoi(cmdType)
			if err != nil {
				runtime.EventsEmit(a.ctx, "errormsg", "Invalid Command Type")
				return err
			}

			// Get command name
			commandName = commandParts[2]

			// Save command data to CommandsByCode and CommandsByName
			CommandsByCode[commandCode] = Command{
				commandCode, commandName, commandType,
			}
			CommandsByName[commandName] = Command{
				commandCode, commandName, commandType,
			}

		}

	}

	runtime.EventsEmit(a.ctx, "statusmsg", "Queue states for fetching")

	// Queue states for fetching
	for _, state := range stateList {
		queue(state)
	}

	runtime.EventsEmit(a.ctx, "statusmsg", "Start polling monitor")

	// Start polling loop
	go func() {
		poll(a)
	}()

	time.Sleep(queueDelay * delayFactor * time.Millisecond)

	isIFConnected = true // We're now connected even if not polling
	runtime.EventsEmit(a.ctx, "ifconnected")

	runtime.EventsEmit(a.ctx, "statusmsg", "Start polling states")

	// Send initial fetch for each state
	for state, _ := range pollStates {
		fetch(a, state)
		time.Sleep(queueDelay * delayFactor * time.Millisecond)
	}

	return nil

}

// Front-end broker function to end IF connection
func (a *App) IFCend() error {

	c.Close()

	runtime.EventsEmit(a.ctx, "ifdisconnected")

	return nil

}

// Get environment info from front end
func (a *App) GetEnvironment() runtime.EnvironmentInfo {
	return envInfo
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	envInfo = runtime.Environment(ctx)
}
