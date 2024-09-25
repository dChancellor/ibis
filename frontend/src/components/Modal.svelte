<script lang="ts">
    import { activeModal } from '@stores/modalContext';
  
    export let dialog: HTMLDialogElement;
</script>

<dialog bind:this={dialog}>
  <slot />
</dialog>

<style>
  dialog::backdrop {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
  }
</style>
