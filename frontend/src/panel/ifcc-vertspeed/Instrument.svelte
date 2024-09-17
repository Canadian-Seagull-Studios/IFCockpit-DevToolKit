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

    // GA Vertical Speed Instrument
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    // Instrument needs to receive states and previous objects from main App.svelte
    export let states;
    export let previous;

    import dialSVG from "./dial.svg";
    import pointerSVG from "./pointer.svg";

    // Return list of required states
    export function getStates() {
        return [
            "aircraft/0/vertical_speed"
        ]
    }

    dispatch("logmsg","Apply vertical_speed to dial");

    // Define maximum and minimum vertical speeds
    const maxVertSpeed = 2000;
    const minVertSpeed = -2000;

    // Calculate vertical speed
    $: vSpeed = states["aircraft/0/vertical_speed"] * 60 * 3.28084;

    // Apply limits
    $: if ( vSpeed > maxVertSpeed) { vSpeed = maxVertSpeed; }
    $: if ( vSpeed < minVertSpeed) { vSpeed = minVertSpeed; }

    // Derive degrees for dial
    $: angle = vSpeed * (172 / maxVertSpeed);

    $: transform = `rotate(${angle}deg)`;

</script>

<div>
    <img src={dialSVG} />
    <img src={pointerSVG} style:transform={transform} />  
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