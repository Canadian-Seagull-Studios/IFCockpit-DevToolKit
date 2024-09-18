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

        This instrument uses images distributed under the:

        Creative Commons bution-Share Alike 4.0 International
        (CC BY-SA 4.0) license
        https://creativecommons.org/licenses/by-sa/4.0/deed.en
        
        by רונאלדיניו המלך on WikiMedia Commons:
        
        https://commons.wikimedia.org/wiki/File:BASIC_Flight_instruments_Improved.svg
        
        The images of this instrument can be redistributed under the
        same licenses. 
                
    */

    // GA Turn Coordinator instrument
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    // Instrument needs to receive states and previous objects from main App.svelte
    export let states;
    export let previous;

    import backgroundSVG from "./background.svg";
    import slipSVG from "./slip.svg";
    import foregroundSVG from "./foreground.svg";
    import bankSVG from "./bank.svg";

    let crosshatchDeg;
    let pointer10000Transition;
    let pointer1000Transition;
    let pointer100Transition;

    // Return list of required states
    export function getStates() {
        return [
            "aircraft/0/bank",
            "aircraft/0/2_minutes_turn_ratio"
        ]
    }

    dispatch("logmsg","Apply bank and 2_minutes_turn_ratio to instrument");

    // Bank
    const bankLimit = 20; // max rotation in degrees
    $: bankDeg = Math.round((states["aircraft/0/bank"] * 180) / Math.PI);
    $: if (bankDeg > bankLimit) { bankDeg = bankLimit; }
    $: if (bankDeg < -bankLimit) { bankDeg = -bankLimit; }
    $: bankTransform = 'rotate('+bankDeg+'deg)';

    // Slip
    const xrange = 85; // Width of centre to edge is 85px
    const yrange = 13; // Y diff between centre and edge is 10px
    const max = 5; // Max slip value when indicator hits edge
    $: turn = -states["aircraft/0/2_minutes_turn_ratio"];
    $: if (turn > max) { turn = max; }
    $: if (turn < -max) { turn = -max; }
    $: transX = turn * (xrange/max);
    $: transY = -Math.abs(turn) * (yrange/max);
    $: slipTransform = 'translateX('+transX+'px) translateY('+transY+'px)';
</script>

<div>
    <img src={backgroundSVG} />
    <img src={slipSVG} style:transform={slipTransform} />
    <img src={foregroundSVG} />
    <img src={bankSVG} style:transform={bankTransform} />
</div>

<style>
    img {
        position: absolute;
        top: 0;
        left: 0;
        transition-property: transform;
        transition-duration: 0.2s;
        transition-timing-function: linear;
        transition-delay: 0s;
        width: 100%;
        height: auto;
    }
</style>