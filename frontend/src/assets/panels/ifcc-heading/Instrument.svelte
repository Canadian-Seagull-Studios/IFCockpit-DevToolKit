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

    // GA Heading instrument
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    // Instrument needs to receive states and previous objects from main App.svelte
    export let states;
    export let previous;

    import backgroundSVG from "./background.svg";
    import dialSVG from "./dial.svg";
    import foregroundSVG from "./foreground.svg";

    let transition;

    // Return list of required states
    export function getStates() {
        return [
            "aircraft/0/heading_magnetic"
        ]
    }

    dispatch("logmsg","Apply heading_magnetic to dial");

    // Get the heading in radians from IF and convert to degrees
    $: headingDeg = (states["aircraft/0/heading_magnetic"] * 180) / Math.PI;

    // Inverse it because of how our dial graphics work (so a heading of 1 degree requires rotating the dial to 359 degrees)
    $: finalHeading = 360 - headingDeg

    dispatch("logmsg",`finalHeading: ${finalHeading}`);

    // Remove the tansition when we are close to North or the visual behaviour is weird when we cross north
    $: if (finalHeading < 2 || finalHeading > 358) { transition = "0s";} else { transition = "0.2s"; }

    // Transform
    $: transform = 'rotate('+finalHeading+'deg)';

</script>

<div>
    <img src={backgroundSVG} />
    <img src={dialSVG} style:transition-duration={transition} style:transform={transform} />
    <img src={foregroundSVG} />
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