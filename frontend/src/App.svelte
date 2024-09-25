<script lang="ts">
  import type { CounterType } from "@globaltypes/global";

  import { GetConstants } from "@wails/application/App.js";
  import { constants } from "@stores/backendConstants";
  import { activeSkill } from "@stores/skillStore";
  import {throwError} from '@helpers/errorHandler'
  import Sidebar from "@sections/Sidebar/Sidebar.svelte";
  import TopBar from "@sections/TopBar/TopBar.svelte";
  import Stopwatch from "@sections/Stopwatch/Stopwatch.svelte";

  let counterType: CounterType = "stopwatch";

  async function getConstantsFromBackend(): Promise<void> {
    try {
      let res = await GetConstants();
      constants.set(res);
    } catch(err) {
      throwError("The application failed loading", err);
    }
  }
</script>

<svelte:window on:load={getConstantsFromBackend} />

<Sidebar />
<main>
    <!-- {#if $activeSkill}
      <TopBar />
      {#if counterType === "stopwatch"}
        <Stopwatch />
      {/if}
    {/if} -->
</main>

<style>
  :global(body){
    height: 100vh;
    display: flex;
    flex-flow: row;
    background-color: #0e1a30;
  }
  :global(*) {
    margin: 0;
    font-family: "Poppins";
    font-weight: 400;
    box-sizing: border-box; 
  }
  :global(button) {
    all: unset;
  }
  :global(button:focus) {
    outline: none;
  }
  @font-face {
    font-family: "Poppins";
    font-style: normal;
    font-weight: 400;
    src:
      local(""),
      url("assets/fonts/Poppins-Regular.ttf") format("truetype");
  }
  @font-face {
    font-family: "Poppins";
    font-style: normal;
    font-weight: 600;
    src:
      local(""),
      url("assets/fonts/Poppins-SemiBold.ttf") format("truetype");
  }  
  @font-face {
    font-family: "Poppins";
    font-style: normal;
    font-weight: 900;
    src:
      local(""),
      url("assets/fonts/Poppins-Black.ttf") format("truetype");
  }
  main {
    flex-grow: 2;
    display: flex;
    justify-content: center;
    margin: 2rem 0rem;
    flex-flow:column;
    align-items:center;
    gap: 2rem;
  }

</style>
