<script lang="ts">
  import Button from "../../../components/Button.svelte";
  export let name: string;
  export let icon: string = null;
  export let isActive: boolean;
</script>

<div class="pill" class:isActive>
  <Button text={name} onClick={() => null} />
</div>

<style>
</style>
