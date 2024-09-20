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

    // Controls bar along top of window

    import Testmode from './Testmode/Control.svelte';
    import LogModal from './LogModal/Control.svelte';

    // Attributes to be passed into the component
    export let isDev;
    export let testMode;
    export let logModal;

</script>

<div class="fixed flex top-0 left-0 right-0 h-12 bg-base-100 m-0 p-0">

    <div class="flex-none gap-4 ms-4">
    </div>
    <div class="flex-none gap-4 ms-4">
    </div>
    <div class="flex-1 justify-center">
    </div>
    <div class="flex-none gap-4 me-4 h-12">
        <Testmode bind:value={testMode} />
    </div>
    <div class="flex-none gap-4 me-4 h-12">
        <LogModal bind:value={logModal} />
    </div>

</div>
