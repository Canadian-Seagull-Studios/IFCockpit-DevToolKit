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

        This instrument uses images distributed under the:

        Creative Commons bution-Share Alike 4.0 International
        (CC BY-SA 4.0) license
        https://creativecommons.org/licenses/by-sa/4.0/deed.en
        
        by רונאלדיניו המלך on WikiMedia Commons:
        
        https://commons.wikimedia.org/wiki/File:BASIC_Flight_instruments_Improved.svg
        
        The images of this instrument can be redistributed under the
        same licenses.        

    */

    // GA Altimeter instrument
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    // Instrument needs to receive states and previous objects from main App.svelte
    export let states;
    export let previous;

    import barometerOverlaySVG from "./barometer_overlay.svg";
    import barometerSVG from "./barometer.svg";
    import centreBackgroundSVG from "./centre_background.svg";
    import crosshatchSVG from "./crosshatch.svg";
    import dialSVG from "./dial.svg";
    import outerRimSVG from "./outer_rim.svg";
    import pointer100SVG from "./pointer_100.svg";
    import pointer1000SVG from "./pointer_1000.svg";
    import pointer10000SVG from "./pointer_10000.svg";

    let crosshatchDeg;
    let pointer10000Transition;
    let pointer1000Transition;
    let pointer100Transition;

    // Return list of required states
    export function getStates() {
        return [
            "aircraft/0/altitude_msl"
        ]
    }

    dispatch("logmsg","Apply altitude_msl to dial");

    // Cross hatch
    const threshold = 10000; // Threshold for crosshatch is 10000
    $: if (states["aircraft/0/altitude_msl"] < threshold) { crosshatchDeg = 0; } else { crosshatchDeg = 90; }
    $: crosshatchTransform = 'rotate('+crosshatchDeg+'deg)';

    // Pointer 10000
    $: pointer10000Deg = (states["aircraft/0/altitude_msl"] / 100000) * 360; // This is the the thousands pointer so max is 100000
    $: if (pointer10000Deg < 2 || pointer10000Deg > 358) { pointer10000Transition = "0s";} else { pointer10000Transition = "0.2s"; }
    $: pointer10000Transform = 'rotate('+pointer10000Deg+'deg)';

    // Pointer 1000
    $: pointer1000Deg = (states["aircraft/0/altitude_msl"] / 10000) * 360; // This is the thousands pointer so max is 10000
    $: if (pointer1000Deg < 2 || pointer1000Deg > 358) { pointer1000Transition = "0s";} else { pointer10000Transition = "0.2s"; }
    $: pointer1000Transform = 'rotate('+pointer1000Deg+'deg)';

    // Pointer 100
    $: pointer100Deg = (states["aircraft/0/altitude_msl"] / 1000) * 360; // This is the hundreds pointer so max is 1000
    $: if (pointer100Deg < 2 || pointer100Deg > 358) { pointer100Transition = "0s";} else { pointer10000Transition = "0.2s"; }
    $: pointer100Transform = 'rotate('+pointer100Deg+'deg)';

</script>

<div>
    <img src={barometerSVG} />
    <img src={barometerOverlaySVG} />
    <img src={crosshatchSVG} style:transform={crosshatchTransform} />
    <img src={dialSVG} />
    <img src={outerRimSVG} />
    <img src={centreBackgroundSVG} />
    <img src={pointer10000SVG} style:transition-duraction={pointer10000Transition} style:transform={pointer10000Transform} />
    <img src={pointer1000SVG} style:transition-duraction={pointer1000Transition} style:transform={pointer1000Transform} />
    <img src={pointer100SVG} style:transition-duraction={pointer100Transition} style:transform={pointer100Transform} />
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