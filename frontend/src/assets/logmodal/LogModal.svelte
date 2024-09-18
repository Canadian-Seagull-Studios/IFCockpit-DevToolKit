<script>

    export let logModal;
    export let errors;
    export let status;

</script>

<dialog id="logModal" class="modal" class:modal-open={logModal}>
    <div class="modal-box h-3/4">
      <div class="grid grid-cols-2 gap-4 m-4 h-4/5 overflow-y-scroll">
        <pre class="text-left text-red-600">{errors}</pre>
        <pre class="text-left">{status}</pre>
      </div>
      <div class="modal-action">
        <button class="btn" on:click={()=>logModal = false}>Close</button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
</dialog>
