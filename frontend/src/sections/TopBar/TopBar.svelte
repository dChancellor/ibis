<script lang='ts'>
  import { DeleteSkill } from "@wails/main/App";

    import { activeSkill, skills } from '@stores/skillStore';

    async function deleteSkill(name: string): Promise<void> {
    try {
      const res = await DeleteSkill(name);
      skills.set(res);
      activeSkill.set(null);
    } catch {
      console.error("The application failed loading");
    }
  }
</script>

<div class='topbar'>
    <!-- TODO make these icon buttons -->
    <button on:click={() => activeSkill.set(null)}>Back</button>
    <p>{$activeSkill.Name}</p>
    <!-- TODO show a modal to confirm deletion -->
    <button on:click={() => deleteSkill($activeSkill.Name)}>Delete</button>
</div>

<style>
.topbar {
    width: 90%;
    height: 50px;
    background-color: #2f2c47;
    border-radius: 1rem;
    display:flex;
    flex-flow:row;
    justify-content: space-between;
    padding:.25rem 1rem;
    margin-bottom:auto;
}
</style>