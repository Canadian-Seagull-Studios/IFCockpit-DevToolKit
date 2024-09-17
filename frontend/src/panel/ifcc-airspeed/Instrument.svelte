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

        Creative Commons Attribution-ShareAlike 3.0 Unported
        (CC BY-SA 3.0) license
        https://creativecommons.org/licenses/by-sa/3.0/deed.en
        
        by Mysid on WikiMedia Commons:
        
        https://commons.wikimedia.org/wiki/File:Airspeed_indicator.svg

        The images have been extended using parts of images distributed
        under the:
        
        Creative Commons bution-Share Alike 4.0 International
        (CC BY-SA 4.0) license
        https://creativecommons.org/licenses/by-sa/4.0/deed.en
        
        by רונאלדיניו המלך on WikiMedia Commons:
        
        https://commons.wikimedia.org/wiki/File:BASIC_Flight_instruments_Improved.svg
        
        The images of this instrument can be redistributed under the
        respective licenses.

    */

    // GA Airspeed instrument
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    // Instrument needs to receive states and previous objects from main App.svelte
    export let states;
    export let previous;

    import backgroundSVG from "./background.svg";
    import dialSVG from "./dial.svg";

    let transform;

    // Return list of required states
    export function getStates() {
        return [
            "aircraft/0/indicated_airspeed"
        ]
    }

    dispatch("logmsg","Apply indicated_airspeed to dial");

    // Steps are used to reflect variable scale on the airspeed gauge
    const steps = [
        {
            knots: 40,
            degrees: 35,
            increment: (55/30)
        },
        {
            knots: 70,
            degrees: 90,
            increment: (120/60)
        },
        {
            knots: 130,
            degrees: 210,
            increment: (54/30)
        },
        {
            knots: 160,
            degrees: 264,
            increment: (48/40)
        },
        {
            knots: 200,
            degrees: 312,
            increment: (48/40)
        },
    ];

    // Before 40 knots we cover 35 degrees on the dial
    const baseIncrement = (35/40);

    // We limit this airspeed gauge to 210 knots
    const maxSpeed = 210;

    // Calculate the placement of the dial at 210 knots (10 knots after 200 knots step)
    const maxDegrees = 312 + (10 * (48/40));

    // If provides airspeed as metres per second -- let's convert to knots
    $: knots = states["aircraft/0/indicated_airspeed"] * 1.94384;

    // Assume the knots are below the base increment and calculate the dial degrees
    $: deg = knots * baseIncrement;

    // Enforce the maximum speed for the gauge
    $: if (knots > maxSpeed) { deg = maxDegrees; }

    // Scale adjustments per the steps in the dial layout
    $: if (knots > steps[0].knots) { deg = steps[0].degrees + ((knots - steps[0].knots) * steps[0].increment); }
    $: if (knots == steps[0].knots) { deg = steps[0].knots; }

    $: if (knots > steps[1].knots) { deg = steps[1].degrees + ((knots - steps[1].knots) * steps[1].increment); }
    $: if (knots == steps[1].knots) { deg = steps[1].knots; }

    $: if (knots > steps[2].knots) { deg = steps[2].degrees + ((knots - steps[2].knots) * steps[2].increment); }
    $: if (knots == steps[2].knots) { deg = steps[2].knots; }

    $: if (knots > steps[3].knots) { deg = steps[3].degrees + ((knots - steps[3].knots) * steps[3].increment); }
    $: if (knots == steps[3].knots) { deg = steps[3].knots; }

    $: if (knots > steps[4].knots) { deg = steps[4].degrees + ((knots - steps[4].knots) * steps[4].increment); }
    $: if (knots == steps[4].knots) { deg = steps[4].knots; }

    // Transform
    $: transform = 'rotate('+deg+'deg)';

</script>

<div>
    <img src={backgroundSVG} />
    <img src={dialSVG} style:transform={transform} />
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