<script lang="ts">
  type ButtonStyle = 'primary' | 'secondary' | 'tertiary'
  type ButtonSize = 'small' | 'medium' | 'large'

  export let onClick: () => void;
  export let iconBefore: string = null;
  export let iconAfter: string = null;
  export let style: ButtonStyle = 'primary';
  export let size: ButtonSize = 'medium';
  export let full: boolean = false;
</script>

<button on:click={onClick} on:keypress={onClick} class:full class={`${style} ${size}` }>
  {#if iconBefore}
    <span class="icon">{@html iconBefore}</span>
  {/if}
  <slot />
  {#if iconAfter}
    <span class="icon">{@html iconAfter}</span>
  {/if}
</button>

<style>
  button{
    display:flex;
    align-items: center;
    justify-content: center;
    border-radius: 8px;
    margin-block: var(--margin-block)
  }
.primary {
  color: var(--background-color);
  background-image: linear-gradient(rgb(254, 231, 54), #f2a72d);
}
.primary:hover{
  background-image: linear-gradient(rgb(205, 186, 44), #c78822);

}
.secondary { 
  color: white;
}
.tertiary {
  color: transparent;
}
.small {
  font-size: 0.75rem;
}
.medium {
  font-size: 1rem;
}
.large {
  font-size: 1.25rem;
}
.full {
    width: 100%;
  }
</style>
