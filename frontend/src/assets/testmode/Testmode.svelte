<script>

    export let states;
    export let previous;

    // Variables for test state values for test mode
    let test_altitude_agl = 0;
    let test_altitude_msl = 0;
    let test_vertical_speed = 0;
    let test_heading_magnetic = 0;
    let test_indicated_airspeed = 0;
    let test_bank = 0;
    let test_pitch = 0;

    // Reset test mode by resetting values of test mode variables and updating previous/states objects
    function resetTestMode() {

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


</script>

<div id="testmode" class="grid grid-cols-9 gap-2 m-0 p-2 bg-neutral">

    <div class="p-0 mx-1 text-sm col-span-2">
      <div>altitude_msl [{test_altitude_msl * 100}]</div>
      <input type="range" min="0" max="480" bind:value={test_altitude_msl} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/altitude_msl"] = states["aircraft/0/altitude_msl"]; states["aircraft/0/altitude_msl"] = test_altitude_msl * 100}} />
    </div>

    <div class="p-0 mx-1 text-sm col-span-2">
      <div>altitude_agl [{test_altitude_agl * 100}]</div>
      <input type="range" min="0" max="480" bind:value={test_altitude_agl} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/altitude_agl"] = states["aircraft/0/altitude_agl"]; states["aircraft/0/altitude_agl"] = test_altitude_agl * 100}} />
    </div>

    <div class="p-0 mx-1 text-sm col-span-2">
      <div>vertical_speed [{Math.round((test_vertical_speed / 10) * 60 * 3.28084)}]</div>
      <input type="range" min="-300" max="300" bind:value={test_vertical_speed} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/vertical_speed"] = states["aircraft/0/vertical_speed"]; states["aircraft/0/vertical_speed"] = test_vertical_speed / 10}} />        
    </div>

    <div class="p-0 mx-1 text-sm col-span-2">
      <div>heading_magnetic [{Math.floor(((test_heading_magnetic / 100) * 180) / Math.PI)}]</div>
      <input type="range" min="0" max="600" bind:value={test_heading_magnetic} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/heading_magnetic"] = states["aircraft/0/heading_magnetic"]; states["aircraft/0/heading_magnetic"] = test_heading_magnetic / 100}} />        
    </div>

    <div class="p-4 mx-1 row-span-2 content-center"><svg on:click={resetTestMode} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--!Font Awesome Free 6.6.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.--><path d="M125.7 160l50.3 0c17.7 0 32 14.3 32 32s-14.3 32-32 32L48 224c-17.7 0-32-14.3-32-32L16 64c0-17.7 14.3-32 32-32s32 14.3 32 32l0 51.2L97.6 97.6c87.5-87.5 229.3-87.5 316.8 0s87.5 229.3 0 316.8s-229.3 87.5-316.8 0c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0c62.5 62.5 163.8 62.5 226.3 0s62.5-163.8 0-226.3s-163.8-62.5-226.3 0L125.7 160z"/></svg></div>

    <div class="p-0 mx-1 text-sm col-span-2">
      <div>indicated_airspeed [{Math.floor(test_indicated_airspeed * 1.94384)}]</div>
      <input type="range" min="0" max="205" bind:value={test_indicated_airspeed} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/indicated_airspeed"] = states["aircraft/0/indicated_airspeed"]; states["aircraft/0/indicated_airspeed"] = test_indicated_airspeed}} />        
    </div>

    <div class="p-0 mx-1 text-sm col-span-2">
      <div>pitch [{Math.round(((test_pitch / 100) * 180) / Math.PI)}]</div>
      <input type="range" min="-200" max="200" bind:value={test_pitch} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/pitch"] = states["aircraft/0/pitch"]; states["aircraft/0/pitch"] = test_pitch / 100}} />        
    </div>

    <div class="p-0 mx-1 text-sm col-span-2">
      <div>bank [{Math.round(((test_bank / 100) * 180) / Math.PI)}]</div>
      <input type="range" min="-100" max="100" bind:value={test_bank} class="range range-warning range-xs" on:input={() => {previous["aircraft/0/bank"] = states["aircraft/0/bank"]; states["aircraft/0/bank"] = test_bank / 100}} />        
    </div>

  </div>

  <style>

    #testmode svg {
        width: 3rem;
    }

    #testmode svg path {
        fill: oklch(var(--a));
    }

    #testmode svg:hover {
        cursor: pointer;
    }

  </style>