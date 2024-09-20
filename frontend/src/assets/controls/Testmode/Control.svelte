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

    // Component to handle toggling test mode
    
    // Attributes to be passed into the component
    export let value = false;
    
</script>

<div>

    <!-- Button to toggle test mode -->
    <span class="clickable text-lg mt-2" class:strikethrough={!value} 
    class:text-slate-500={!value} 
    class:text-success={value}
    on:click={() => {value = !value}}>
        Test Mode
    </span>

</div>

<style>
    span {
        display: inline-block;
    }
    .clickable:hover {
        cursor: pointer;
    }
    .strikethrough {
        text-decoration: line-through;
    }
</style>