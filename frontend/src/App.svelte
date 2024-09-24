<script lang="ts">
  import Stopwatch from "./components/Stopwatch.svelte";
  // import Sidebar from "./sections/Sidebar.svelte";
  import Sidebar from "./sections/Sidebar/Sidebar.svelte";
  import { GetConstants, DeleteSkill } from "../wailsjs/go/main/App.js";
  import { constants } from "./stores/backendConstants";
  import Button from "./components/Button.svelte";
  import { activeSkill, skills } from "./stores/skillStore";
  import Ring from "./components/Ring.svelte";
  import Modal from "@components/Modal.svelte";

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
  <!-- <Ring
        startingColor="#fa114f"
        endColor="#f93885"
        innerDiameter={400}
        outerDiameter={850}
        type={"set"}
        animationLength={25}
        percentFilled={10}
      /> -->
  <div class="temp">
    {#if $activeSkill || true}
      {#if counterType === "stopwatch"}
        <Stopwatch />
      {/if}
      <!-- <Button
        onClick={() => deleteSkill($activeSkill.Name)}
        text={"Delete skill"}
      /> -->
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
    background-color: #111;
    flex-grow: 2;
    display: flex;
    justify-content: center;
    align-items: center;
  }

</style>
