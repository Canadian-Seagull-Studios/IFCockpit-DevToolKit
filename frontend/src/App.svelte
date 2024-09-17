<script>

  /*

    This file is part of the IFCockpit Developer Toolkit
    from Canadian Seagull Studios Ltd.

    This is not open source and may only be used by developers
    who are part of the IFCockpit developer programme under the
    terms of the programme and may not be shared with anyone
    outside the terms of that programme or otherwise with the
    explicit written consent of Canadian Seagull Studios Ltd.

    Copyright 2024, Canadian Seagull Studios Ltd.
    All rights reserved.

  */

  // Main application Svelte file

  import { onMount } from 'svelte';

  import Controls from './assets/controls/Controls.svelte'; // Common UI controls

  import Panel from './panel/Panel.svelte'; // GA panel

  import * as rt from "../wailsjs/runtime/runtime.js";  // Wails runtime

  import TailwindCss from './TailwindCSS.svelte'; // Baseline tailwind CSS

  import {IFCinit,IFCend,GetEnvironment} from '../wailsjs/go/main/App'; // Front end functions exposed by Go backend

  let isDev = false; // Are we in a dev environment?

  let testMode = false; // Is test mode visible?
  let logModal = false; // Is log model visible?

  // Variables for test state values for test mode
  let test_altitude_agl = 0;
  let test_altitude_msl = 0;
  let test_vertical_speed = 0;
  let test_heading_magnetic = 0;
  let test_indicated_airspeed = 0;
  let test_bank = 0;
  let test_pitch = 0;

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

  // Reset test mode by resetting values of test mode variables and updating previous/states objects
  function resetTestMode() {

    log(resetTestMode);

    test_altitude_agl = 0;
    previous["aircraft/0/altitude_agl"] = states["aircraft/0/altitude_agl"]; states["aircraft/0/altitude_agl"] = test_altitude_agl * 100;

    test_altitude_msl = 0;
    previous["aircraft/0/altitude_msl"] = states["aircraft/0/altitude_msl"]; states["aircraft/0/altitude_msl"] = test_altitude_msl * 100;

    test_bank = 0;
    previous["aircraft/0/bank"] = states["aircraft/0/bank"]; states["aircraft/0/bank"] = test_bank / 100;

    test_heading_magnetic = 0;
    previous["aircraft/0/heading_magnetic"] = states["aircraft/0/heading_magnetic"]; states["aircraft/0/heading_magnetic"] = test_heading_magnetic / 100;

    test_indicated_airspeed = 0;
    previous["aircraft/0/indicated_airspeed"] = states["aircraft/0/indicated_airspeed"]; states["aircraft/0/indicated_airspeed"] = test_indicated_airspeed;

    test_pitch = 0;
    previous["aircraft/0/pitch"] = states["aircraft/0/pitch"]; states["aircraft/0/pitch"] = test_pitch / 100;

    test_vertical_speed = 0;
    previous["aircraft/0/vertical_speed"] = states["aircraft/0/vertical_speed"]; states["aircraft/0/vertical_speed"] = test_vertical_speed / 10;

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

  <div class="grid-cols-1">

    <!-- Navbar -->
    <div class="navbar">
      <!-- Controls bar -->
      <Controls isDev={isDev} on:logmsg={(event) => log(event.detail)} bind:testMode={testMode} bind:logModal={logModal} />
    </div>

    <!-- Panel to render depending on selected panel -->
    <Panel bind:this={PanelComponent} states={states} previous={previous} on:panelloaded={() => states = PanelComponent.getStates()} />

  </div>

  <!-- Test mode if selected -->
  <!-- May externalise test mode later in a Svelte component -->
  {#if testMode}
    <div id="testmode" class="grid grid-cols-7 gap-2 m-2">

      <div class="p-0 text-sm">
        <div>altitude_msl [{test_altitude_msl * 100}]</div>
        <input type="range" min="0" max="480" bind:value={test_altitude_msl} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/altitude_msl"] = states["aircraft/0/altitude_msl"]; states["aircraft/0/altitude_msl"] = test_altitude_msl * 100}} />
      </div>

      <div class="p-0 text-sm">
        <div>altitude_agl [{test_altitude_agl * 100}]</div>
        <input type="range" min="0" max="480" bind:value={test_altitude_agl} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/altitude_agl"] = states["aircraft/0/altitude_agl"]; states["aircraft/0/altitude_agl"] = test_altitude_agl * 100}} />
      </div>

      <div class="p-0 text-sm">
        <div>vertical_speed [{Math.round((test_vertical_speed / 10) * 60 * 3.28084)}]</div>
        <input type="range" min="-300" max="300" bind:value={test_vertical_speed} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/vertical_speed"] = states["aircraft/0/vertical_speed"]; states["aircraft/0/vertical_speed"] = test_vertical_speed / 10}} />        
      </div>

      <div class="p-0 text-sm">
        <div>heading_magnetic [{Math.floor(((test_heading_magnetic / 100) * 180) / Math.PI)}]</div>
        <input type="range" min="0" max="600" bind:value={test_heading_magnetic} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/heading_magnetic"] = states["aircraft/0/heading_magnetic"]; states["aircraft/0/heading_magnetic"] = test_heading_magnetic / 100}} />        
      </div>

      <div class="p-0 text-sm">
        <div>indicated_airspeed [{Math.floor(test_indicated_airspeed * 1.94384)}]</div>
        <input type="range" min="0" max="205" bind:value={test_indicated_airspeed} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/indicated_airspeed"] = states["aircraft/0/indicated_airspeed"]; states["aircraft/0/indicated_airspeed"] = test_indicated_airspeed}} />        
      </div>

      <div class="p-0 text-sm">
        <div>pitch [{Math.round(((test_pitch / 100) * 180) / Math.PI)}]</div>
        <input type="range" min="-200" max="200" bind:value={test_pitch} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/pitch"] = states["aircraft/0/pitch"]; states["aircraft/0/pitch"] = test_pitch / 100}} />        
      </div>

      <div class="p-0 text-sm">
        <div>bank [{Math.round(((test_bank / 100) * 180) / Math.PI)}]</div>
        <input type="range" min="-100" max="100" bind:value={test_bank} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/bank"] = states["aircraft/0/bank"]; states["aircraft/0/bank"] = test_bank / 100}} />        
      </div>

      <div class="col-span-7"><button class="btn btn-accent" on:click={resetTestMode}>Reset</button></div>

    </div>
  {:else}
    <div class="grid grid-cols-1">
      
      <div>
        <input bind:value={ip} type="text" placeholder="Enter IF Device IP Address" class="input input-bordered w-full max-w-xs" />
        <button class="btn btn-accent" on:click={() => IFconnect() }>Connect</button>
        <button class="btn btn-error" on:click={() => IFdisconnect() }>Disconnect</button>    
      </div>

      <div class="p-4 text-lg" class:text-accent={ifConnected} class:text-error={!ifConnected}>{ifConnectedMsg}</div>

    </div>

  {/if}

  <!-- Logs modal -->
  <dialog id="logModal" class="modal" class:modal-open={logModal}>
    <div class="modal-box h-3/4">
      <div class="grid grid-cols-2 gap-4 m-4 h-4/5 overflow-y-scroll">
        <pre class="text-left text-red-600">{errors}</pre>
        <pre class="text-left">{status}</pre>
      </div>
      <div class="modal-action">
        <button class="btn" on:click={()=>logModal = false}>Close</button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>

</main>

<style>

  main {
    background-color: #000;
  }

</style>
