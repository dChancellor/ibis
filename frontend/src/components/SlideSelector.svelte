<script lang="ts">
  import { onMount, SvelteComponent } from 'svelte';
  import { createEventDispatcher } from 'svelte';

  import Button from '@components/Button.svelte';

  export let buttonLabels: string[];
  export let onSelect: (selection: string) => void;

  export let selected = buttonLabels[0];
</script>

<div class="slideselector">
  {#each buttonLabels as label, i}
    <Button
      onClick={() => (selected = label)}
      variant={selected === label ? 'primary' : 'secondary'}
      styles={{
        'border-radius': i === 0 ? '0.5rem 0 0 0.5rem' : i === buttonLabels.length - 1 ? '0 0.5rem 0.5rem 0' : '0',
        'border-right': i === buttonLabels.length - 1 ? 'none' : '1px solid #847866',
      }}>{label}</Button
    >
  {/each}
</div>

<style>
  .slideselector {
    display: flex;
  }
</style>
