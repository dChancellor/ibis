<script lang="ts">
  import { addSkill as addSkillToDB } from "@helpers/skillHooks";
  import { closeModal } from "@stores/modalContext";

  import Button from "@components/Button.svelte";
  import Modal from "@components/Modal.svelte";
  import { constants } from "@stores/backendConstants";

  export let dialog: HTMLDialogElement;

  let errorText: string;

  let newSkillName: string | null;
  let svgFileContent = '';

  let fileInput: HTMLInputElement;

  function handleFileUpload(event) {
    const file = event.target.files[0];
    if (file && file.type === 'image/svg+xml') {
      const reader = new FileReader();
      reader.onload = () => {
        svgFileContent = reader.result as string; 
      };
      reader.readAsText(file);  
    } else {
      errorText = 'Please upload an SVG file';
    }
  }

  function addSkill() {
    try {
      addSkillToDB(newSkillName, svgFileContent);
      newSkillName = null;
      svgFileContent = '';
      fileInput.value = '';
      closeModal(dialog);
    } catch(err) {
      if ((err = $constants.SkillExistsMessage)) {
        errorText = "This skill name already exists";
      } else {
        errorText = "There was an error saving your skill";
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
  <input type="file" accept=".svg" on:change={handleFileUpload} bind:this={fileInput}/>
  <Button onClick={addSkill}>Add skill</Button>
  <Button
    onClick={() => {
      closeModal(dialog);
    }}
  >Close</Button>
</Modal>

<style>
</style>
