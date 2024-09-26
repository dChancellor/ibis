<script lang="ts">
  import MaterialSymbolsArrowBack from '~icons/material-symbols/arrow-back';
  import PhTrashBold from '~icons/ph/trash-bold';

  import { activeSkill } from '@stores/skillStore';
  import { deleteSkill } from '@helpers/skillHooks';
  import Button from '@components/Button.svelte';

  let dialog: HTMLDialogElement;
</script>

<div class="topbar">
  <Button variant={'secondary'} size={'medium'} onClick={() => activeSkill.set(null)}>
    <MaterialSymbolsArrowBack />
  </Button>
  <p>{$activeSkill.Name}</p>
  <!-- TODO show a modal to confirm deletion -->
  <Button variant={'secondary'} size={'medium'} onClick={() => deleteSkill($activeSkill.Name)}><PhTrashBold /></Button>
</div>

<style>
  .topbar {
    height: 50px;
    background-color: #13213b88;
    border-radius: 1rem;
    display: flex;
    flex-flow: row;
    justify-content: space-between;
    padding: 0.25rem 1rem;
    width: 100%;
    align-items: center;
  }
  p {
    margin-right: auto;
    margin-left: 1rem;
  }
</style>
