<script>

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

  // Main application Svelte file

  import { onMount } from 'svelte';

  import Controls from './assets/controls/Controls.svelte'; // Common UI controls
  import Testmode from './assets/testmode/Testmode.svelte'; // Testmode controls
  import LogModal from './assets/logmodal/LogModal.svelte'; // Log Modal

  import Panel from './panel/Panel.svelte'; // GA panel

  import * as rt from "../wailsjs/runtime/runtime.js";  // Wails runtime

  import TailwindCss from './TailwindCSS.svelte'; // Baseline tailwind CSS

  import {IFCinit,IFCend,IFCset,GetEnvironment} from '../wailsjs/go/main/App'; // Front end functions exposed by Go backend

  let isDev = false; // Are we in a dev environment?

  let testMode = false; // Is test mode visible?
  let logModal = false; // Is log model visible?

  // Will reference the selected panel component later
  let PanelComponent;

  let ifConnected = false; // Are we connected to IF?
  let ifConnectedMsg = "IF Disconnected";

  let ip = ""; // IP address of IF
  let ipList = []; // List of IP addresses found by UDP in an array

  let errors = "Errors\n"; // Error log

  let status = "Log\n"; // Message log

  let states = {}; // Holds current state values
  let previous = {}; // Holds last state value
  let counts = {}

  let devices = []; // Store list of IF device IP address -- the RAW JSON list from IF

  let envInfo; // Environment info from backend

  // Handle errormsg event from backend -- update error log
  rt.EventsOn('errormsg', (msg) => {
    errors += `${msg}\n`;
  });

  // Handle statusmsg event from backend -- updatge message log
  rt.EventsOn('statusmsg', (msg) => {
    status += `${msg}\n`;
  });

  // Handle statechange event from backend
  rt.EventsOn('statechange', (state)=>{

    log(state);

    // Keep count of number of times state has been returned -- used in debugging
    counts[state.name] = (counts.hasOwnProperty(state.name)) ? (counts[state.name] + 1) : 0;

    // Keep record of last value of the state
    previous[state.name] = states[state.name];

    // Store the state in states object based on state type
    switch (state.type) {
      case 0: // Boolean
        states[state.name] = state.val;
        break;
      case 1: // Integer
        states[state.name] = parseInt(state.val);
        break;
      case 2: // Float
        let floatval = Math.round(parseFloat(state.val) * 100) / 100;
        states[state.name] = floatval;
        break;
      case 3: // Double
        let doubleval = Math.round(parseFloat(state.val) * 100) / 100;
        states[state.name] = doubleval;
        break;
      case 4: // String
        states[state.name] = state.val;
        break;
      case 5: // Int64
        states[state.name] = parseInt(state.val);
        break;
    }

  });

  // Handle devicefound event -- IF has returned a list of IPs by UDP
  rt.EventsOn('devicefound', (device) => {

    log("Device found, checking addresses");

    // Get list of device addresses from JSON
    devices = device.addresses;

    // Iterate list of addresses and get first non-local IP 4 to use as the assumed preferred/default address for the device
    for (let addr of device.addresses) {

      log(addr);

      // IP 4 means three dots separating four parts of the string
      if (addr != "127.0.0.1" && addr.split(".").length == 4) {
        ip = addr;
        log(ip);
      }

      // Add to ipList array
      ipList.push(addr);

    }

  });

  // Handle ifconnected event -- update UI
  rt.EventsOn('ifconnected', () => {
    log("IF Connected");
    ifConnectedMsg = "IF Connected";
    ifConnected = true;
  });

  // Handle ifdisconnected event -- update UI and reset global vars rto defaults
  rt.EventsOn('ifdisconnected', () => {
    log("IF Disconnected");
    ifConnectedMsg = "IF Disconnected";
    ifConnected = false;
    ip = "";
    ipList = [];
    devices = [];
  });

  // Console logging -- depending on whether in dev or prod
  function log(msg) {
    if (isDev) { console.log(msg); }
  }

  // Set a state from the front end by calling IFCset backend function
  function setState(state,val) {
    if (ifConnected) {
      try {
          IFCset(state,val);
      } catch (err) {
          console.error(err);
          return false;
      }
      return true;
    }
    return false;
  }

  // Connect to IF on a specified IP by calling IFCinit back-end function
  function IFconnect() {

    log(states);
    log(`Connecting to ${ip}`);

    try {

      log("Call IFCinit");

      // Call IFCinit with the list of states to monitor and the IP address of the IF device
      IFCinit(Object.keys(states),ip).then( result => {
        log("IFCinit called");
      });

    } catch (err) {

      console.error(err);

    }

  }

  // Disconnect from IF by calling IFCend backend function
  function IFdisconnect() {
    log(`Disconnecting from ${ip}`);
    try {
      log("Call IFdisconnect");
      IFCend().then( result => {
        log("IFCend called");
      });
    } catch (err) {
        console.error(err);
    }
  }

/*  function showPanel() {
    PanelComponent = new Panel({
      target: document.getElementById("panel"),
      props: {
        states: {states},
        previous: {previous}
      }
    });
  }*/

  // Start panel by connecting to a flight
  function startFlight(addr) {

    log(addr);

    // Select IP address to connect
    ip = addr;
//    loader = "none";
//    showPanel();

    // Connect
    IFconnect();

  }

  // Initialise front-end once the App.svelte component loads
  onMount(async function() {

//    showPanel();

    log("onMount");

    envInfo = await GetEnvironment();
    isDev = (envInfo.buildType == "dev");
    log(`isDev: ${isDev}`);

    // Get panel's list of required states
//    states = PanelComponent.getStates();
//    log(states);

    // For weird behaviour of bank on load -- may be able to remove later after debug
    states["aircraft/0/bank"] = 0.1;
    states["aircraft/0/bank"] = 0;

  });

</script>

<TailwindCss/>

<main>

  <!-- Controls bar -->
  <Controls isDev={isDev} on:logmsg={(event) => log(event.detail)} bind:testMode={testMode} bind:logModal={logModal} />

  <!-- Test mode if selected -->
  <div class="fixed bottom-0 left-0 right-0 grid-cols-1 h-24 bg-base-100">
    {#if testMode}
      <Testmode bind:states={states} bind:previous={previous} />
    {:else}
      <div class="p-4">
        <input bind:value={ip} type="text" placeholder="Enter IF Device IP Address" class="input input-bordered w-full max-w-xs" />
        <button class="btn btn-accent" on:click={() => IFconnect() }>Connect</button>
        <button class="btn btn-error" on:click={() => IFdisconnect() }>Disconnect</button>    
        <span class="p-4 text-lg" class:text-accent={ifConnected} class:text-error={!ifConnected}>{ifConnectedMsg}</span>
      </div>
    {/if}
  </div>

  <!-- Logs modal -->
  <LogModal bind:logModal={logModal} errors={errors} status={status} />

  <div class="fixed top-12 bottom-24 left-0 right-0 p-0 m-0">

    <!-- Panel to render depending on selected panel -->
    <Panel bind:this={PanelComponent} states={states} previous={previous} 
      on:panelloaded={() => states = PanelComponent.getStates()} 
      on:setstate={(event) => { setState(event.detail.state,event.detail.value) }} />

  </div>

</main>

<style>

  main {
    background-color: #000;
  }

</style>
