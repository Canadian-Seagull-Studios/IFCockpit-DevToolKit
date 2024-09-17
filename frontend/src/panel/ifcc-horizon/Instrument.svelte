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

    // GA Horizon instrument
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    // Instrument needs to receive states and previous objects from main App.svelte
    export let states;
    export let previous;

    import backgroundSVG from "./background.svg";
    import outerdialSVG from "./outerdial.svg";
    import foregroundSVG from "./foreground.svg";

    let transition;

    // Return list of required states
    export function getStates() {
        return [
            "aircraft/0/bank",
            "aircraft/0/pitch"
        ]
    }

    dispatch("logmsg","Apply bank and pitch to instrument");

    // Maximum instrument width is used in calculating pitch
    const maxWidth = 400;
    const width = 800; // How to get the instrument width?????

    // Calculate the bank
    $: bankDeg = -Math.round((states["aircraft/0/bank"] * 180) / Math.PI);

    // Outerdial
    $: if (bankDeg < 2 || bankDeg > 358) { transition = "0s";} else { transition = "0.2s"; }
    $: outerdialTransform = 'rotate('+bankDeg+'deg)';

    // Background
    $: pitchSize = Math.round((states['aircraft/0/pitch'] * 180) / Math.PI) * 3 * (width / maxWidth);
    $: backgroundTransform = 'rotate('+Math.floor(bankDeg)+'deg) translateY('+pitchSize+'px)';

</script>

<div>
    <img src={backgroundSVG} style:transform={backgroundTransform} />
    <img src={outerdialSVG} style:transition-duration={transition} style:transform={outerdialTransform} />
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