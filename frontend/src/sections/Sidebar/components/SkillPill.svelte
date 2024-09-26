<script lang="ts">
  export let name: string;
  export let icon: string = null;
  export let isActive: boolean;
  export let onClick: () => void;
</script>

<button class="pill" class:isActive on:click={onClick}>
  {#if icon}
    <div class="icon">{@html icon}</div>
  {/if}
  {name}
</button>

<style>
  .pill {
    display: flex;
    flex-flow: row;
    align-items: center;
    background-image: linear-gradient(rgba(41, 55, 91, 0.69), #1b253eb0);
    padding: 0.75rem 1rem;
    border-radius: 1rem;
    margin-inline: 0.5rem;
    position: relative;
    box-shadow: 1px 3px 1px 1px #00000024;
    width: 100%;
    width: 100%;
    font-weight: 700;
  }

  .pill:hover {
    background-color: rgb(49, 46, 76);
  }
  .isActive {
    outline: #53587ca0 2px solid;
  }
  .icon {
    width: 24px;
    height: 24px;
    margin-right: 8px;
  }
  /* TODO - this seems problematic to use global here, but need to figure out how to force svelte to apply styles without using global */
  :global(.pill .icon svg) {
    max-width: 100%;
    max-height: 100%;
  }
</style>
