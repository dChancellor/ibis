<script lang="ts">
  import Button from "@components/Button.svelte";
  import Modal from "@components/Modal.svelte";
  import { constants } from "@stores/backendConstants";
  import { closeModal } from "@stores/modalContext";
  import { AddSkill } from "@wails/main/App";

  export let dialog: HTMLDialogElement;
  export let getSkills: () => void;

  let newSkillName: string | null;
  let errorText: string;

  async function addSkill(newSkillName: string): Promise<void> {
    if (!newSkillName) {
      errorText = "Please enter a skill name";
      return;
    }
    try {
      await AddSkill(newSkillName);
      dialog.close();
      newSkillName = null;
      getSkills();
    } catch (error) {
      if ((error = $constants.SkillExistsMessage)) {
        errorText = "This skill name already exists";
      } else {
        errorText = "There was an error saving your skill";
        console.error(errorText + ":", error);
      }
    }
  }
</script>

<Modal bind:dialog>
  {#if errorText}
    <p>{errorText}</p>
  {/if}
  <input
    name="New Skill"
    bind:value={newSkillName}
    placeholder="Name of the new skill"
  />
  <Button text={"Add skill"} onClick={() => addSkill(newSkillName)} />
  <Button
    text={"Close"}
    onClick={() => {
      closeModal(dialog);
    }}
  />
</Modal>

<style>
</style>
