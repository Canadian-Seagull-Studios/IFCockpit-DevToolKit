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

    // GA panel

    import { onMount } from 'svelte';
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    // Panel needs to receive states and previous objects from main App.svelte
    export let states;
    export let previous;

    // Import all instruments used by the panel
    import Vertspeed from './ifcc-vertspeed/Instrument.svelte';
    import Heading from './ifcc-heading/Instrument.svelte';
    import Airspeed from './ifcc-airspeed/Instrument.svelte';
    import Altimeter from './ifcc-altimeter/Instrument.svelte';
    import Turncoord from './ifcc-turncoord/Instrument.svelte';
    import Horizon from './ifcc-horizon/Instrument.svelte';

    // Holds references to instantiated instruments
    const instruments = {};

    // Function to provide list of states required to App.svelte
    export function getStates() {

        dispatch("logmsg","Panel > getStates");

        let statelist = {};

        // Iterate objects and call getStates in the instrument and added to statelist
        Object.keys(instruments).forEach((instrument) => {
            dispatch("logmsg",instrument);
            // @ts-ignore
            instruments[instrument].getStates().forEach((state) => {
                statelist[state] = "-"
            });
        });

        return statelist;
        
    }

    onMount(function() {
        dispatch("panelloaded",{ state: true });
        dispatch("setstate",{ state: "foo", value: "bar" })
    });

</script>

<div class="grid grid-cols-3 gap-4 m-4">
    <div class="instrument">
        <Heading on:logmsg bind:this={instruments["Heading"]} states={states} previous={previous} />
    </div>
    <div class="instrument">
        <Horizon on:logmsg bind:this={instruments["Horizon"]} states={states} previous={previous} />
    </div>
    <div class="instrument">
        <Vertspeed on:logmsg bind:this={instruments["Vertspeed"]} states={states} previous={previous} />
    </div>
    <div class="instrument">
        <Turncoord on:logmsg bind:this={instruments["Turncoord"]} states={states} previous={previous} />
    </div>
    <div class="instrument">
        <Altimeter on:logmsg bind:this={instruments["Altimeter"]} states={states} previous={previous} />
    </div>
    <div class="instrument">
        <Airspeed on:logmsg bind:this={instruments["Airspeed"]} states={states} previous={previous} />
    </div>
</div>

<style>
    .instrument {
        position: relative;
        height: 32vw;
        width: 32vw;
        overflow: hidden;
    }
</style>