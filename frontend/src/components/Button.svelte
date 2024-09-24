<script lang="ts">
  type ButtonStyle = 'primary' | 'secondary' | 'tertiary'
  type ButtonSize = 'small' | 'medium' | 'large'

  export let onClick: () => void;
  export let iconBefore: string = null;
  export let iconAfter: string = null;
  export let style: ButtonStyle = 'primary';
  export let size: ButtonSize = 'medium';
</script>

<button on:click={onClick} on:keypress={onClick} class={`${style} ${size}`}>
  {#if iconBefore}
    <span class="icon">{@html iconBefore}</span>
  {/if}
  <slot />
  {#if iconAfter}
    <span class="icon">{@html iconAfter}</span>
  {/if}
</button>

<style>
.primary {
  color: white;
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
</style>
