<script lang="ts">
  import Stopwatch from "./components/Stopwatch.svelte";
  import Sidebar from "./sections/Sidebar.svelte";
  import { GetConstants, DeleteSkill } from "../wailsjs/go/main/App.js";
  import { constants } from "./stores/backendConstants";
  import Button from "./components/Button.svelte";
  import { activeSkill, skills } from "./stores/skillStore";
  import Ring from "./components/Ring.svelte";

  async function getConstants(): Promise<void> {
    try {
      let res = await GetConstants();
      constants.set(res);
    } catch {
      console.error("The application failed loading");
    }
  }

  async function deleteSkill(name: string): Promise<void> {
    try {
      const res = await DeleteSkill(name);
      skills.set(res);
    } catch {
      console.error("The application failed loading");
    }
  }

  type CounterType = "timer" | "stopwatch" | "manual";
  let counterType: CounterType = "stopwatch";
</script>

<svelte:window on:load={getConstants} />

<main>
  <Sidebar />
  <Ring
        color="purple"
        innerDiameter={400}
        outerDiameter={450}
        type={"set"}
        animationLength={25}
        percentFilled={25}
      />
  <div class="temp">
    {#if $activeSkill}
      {#if counterType === "stopwatch"}
        <!-- <Stopwatch /> -->
      {/if}
      <Button
        onClick={() => deleteSkill($activeSkill.Name)}
        text={"Delete skill"}
      />
    {/if}
  </div>
</main>

<style>
  main {
    display: flex;
    flex-flow: row;
    height: 100%;
  }
  .temp {
    background-color: rgb(48, 33, 36);
    flex-grow: 2;
    display: flex;
    justify-content: center;
    align-items: center;
  }

</style>
