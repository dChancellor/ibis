<script lang="ts">
  import { activeSkill, skills } from '@stores/skillStore';
  import { openModal } from '@stores/modalContext';
  import { getSkills } from '@helpers/skillHooks';

  import FoundationPlus from '~icons/foundation/plus';

  import SkillPill from '@sections/Sidebar/components/SkillPill.svelte';
  import Button from '@components/Button.svelte';

  import AddSkillModal from './components/AddSkillModal.svelte';

  let dialog: HTMLDialogElement;
</script>

<svelte:window on:load={getSkills} />

<section>
  <div class="skills-list">
    {#each $skills as skill}
      <SkillPill
        name={skill.Name}
        isActive={$activeSkill?.Name === skill.Name}
        icon={skill.SVG}
        onClick={() => activeSkill.set($activeSkill === skill ? null : skill)}
      />
    {/each}
    <Button full onClick={() => openModal(dialog)} size={'small'} --margin-block="1rem">
      <FoundationPlus />
    </Button>
  </div>
</section>
<AddSkillModal bind:dialog />

<style>
  section {
    width: var(--sidebar-width);
    padding: 2rem;
    background-color: #13213b88;
    position: fixed;
    height: 100vh;
    overflow-y: scroll;
  }
  .skills-list {
    display: flex;
    flex-flow: column;
    gap: 0.25rem;
    align-items: center;
    width: 100%;
  }
</style>
