<script lang="ts">
  import { parseStyles } from '@helpers/styleParsing';

  type ButtonVariant = 'primary' | 'secondary' | 'tertiary';
  type ButtonSize = 'small' | 'medium' | 'large';

  export let onClick: () => void;
  export let iconBefore: string = null;
  export let iconAfter: string = null;
  export let variant: ButtonVariant = 'primary';
  export let size: ButtonSize = 'medium';
  export let full: boolean = false;

  export let styles: any = {};

  let internalStyles: string = '';

  $: internalStyles = parseStyles(styles);
</script>

<button on:click={onClick} on:keypress={onClick} class:full class={`${variant} ${size}`} style={internalStyles}>
  {#if iconBefore}
    <span class="icon">{@html iconBefore}</span>
  {/if}
  <slot />
  {#if iconAfter}
    <span class="icon">{@html iconAfter}</span>
  {/if}
</button>

<style>
  button {
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 0.5rem;
    margin-block: var(--margin-block);
    padding: 0.25rem 0.5rem;
  }
  .primary {
    color: var(--background-color);
    background: linear-gradient(rgb(254, 231, 54), #f2a72d);
  }
  .primary:hover {
    background: linear-gradient(rgb(205, 186, 44), #c78822);
  }
  .secondary {
    color: white;
    background-color: #243351;
  }
  .secondary:hover {
    color: white;
    background-color: rgb(27, 39, 64);
    outline: 1px solid #bba076;
    outline-offset: -1px;
  }
  .tertiary {
    color: transparent;
  }
  .small {
    font-size: 1rem;
  }
  .medium {
    height: 1.5rem;
  }
  .large {
    font-size: 1.25rem;
    height: 2rem;
  }
  .full {
    width: 100%;
  }
</style>
