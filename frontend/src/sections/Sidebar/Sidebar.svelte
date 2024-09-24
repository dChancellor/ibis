<script lang="ts">
  import { GetSkills } from "@wails/main/App";
  import SkillPill from "@sections/Sidebar/components/SkillPill.svelte";
  import Button from "@components/Button.svelte";

  import { activeSkill, skills } from "@stores/skillStore";
  import { openModal } from "@stores/modalContext";

  import AddSkillModal from "./components/AddSkillModal.svelte";
  
  let dialog: HTMLDialogElement;

  async function getSkills(): Promise<void> {
    let res = await GetSkills();
    if (!res) {
      console.warn("Hmm, no skills were fetched...");
    }
    console.log("Skills fetched:", res);
    skills.set(res);
  }
</script>

<svelte:window on:load={getSkills} />

<section>
  <div class="skills-list">
    {#each $skills as skill}
      <SkillPill
        name={skill.Name}
        isActive={$activeSkill?.Name === skill.Name}
        icon={'art'}
      />
    {/each}
  </div>
  <!-- TODO just make this an IconButton -->
  <Button text={"+"} onClick={() => openModal(dialog)} />
</section>
<AddSkillModal bind:dialog getSkills={getSkills} />

<style>
  section {
    padding: 2rem;
  }
  .skills-list {
    display: flex;
    flex-flow: column;
  }
</style>
