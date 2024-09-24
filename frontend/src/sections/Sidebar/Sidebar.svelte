<script lang="ts">
  import { activeSkill, skills } from "@stores/skillStore";
  import { openModal } from "@stores/modalContext";
  import { getSkills } from "@helpers/skillHooks";

  import PhPlusFill from '~icons/ph/plus-fill';

  import SkillPill from "@sections/Sidebar/components/SkillPill.svelte";
  import Button from "@components/Button.svelte";

  import AddSkillModal from "./components/AddSkillModal.svelte";
  
  let dialog: HTMLDialogElement;
</script>

<svelte:window on:load={getSkills} />

<section>
  <!-- TODO - Implement a search? -->
  <div class="skills-list">
    {#each $skills as skill}
      <SkillPill
        name={skill.Name}
        isActive={$activeSkill?.Name === skill.Name}
        icon={skill.SVG}
        onClick={() => activeSkill.set($activeSkill===skill ? null : skill)}
      />
    {/each}
  </div>
  <Button onClick={() => openModal(dialog)} size={'large'}><PhPlusFill /></Button>
</section>
<AddSkillModal bind:dialog />

<style>
  section {
    padding-block: 2rem;
    background-color: #2c293e96;
    box-shadow: 1px 0px 1px 0px #dfdeff1c;
  }
  .skills-list {
    display: flex;
    flex-flow: column;
    gap: .25rem;
  }
</style>
