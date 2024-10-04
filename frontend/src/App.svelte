<script lang="ts">
  import type { CounterType } from 'src/typings/global';

  import { GetConstants } from '@wails/application/App.js';
  import { constants } from '@stores/backendConstants';
  import { isSideBarOpen } from '@stores/appStore';
  import { activeSkill } from '@stores/skillStore';
  import { throwError } from '@helpers/errorHandler';
  import Sidebar from '@sections/Sidebar/Sidebar.svelte';
  import TopBar from '@sections/TopBar/TopBar.svelte';
  import Stopwatch from '@sections/Stopwatch/Stopwatch.svelte';

  import SlideSelector from '@components/SlideSelector.svelte';
  import Calendar from '@sections/Calendar/Calendar.svelte';

  let counterType: CounterType = 'Stopwatch';

  async function getConstantsFromBackend(): Promise<void> {
    try {
      let res = await GetConstants();
      constants.set(res);
    } catch (err) {
      throwError('The application failed loading', err);
    }
  }

  const onSelect = (selection: string) => {
    // TODO - this can be better typed I think - I dislike using "as"
    counterType = selection as CounterType;
  };
</script>

<svelte:window on:load={getConstantsFromBackend} />

{#if $isSideBarOpen}
  <Sidebar />
{/if}

<main class:isSideBarOpen>
  {#if $activeSkill}
    <TopBar />
    <SlideSelector buttonLabels={['Stopwatch', 'Timer', 'Manual']} {onSelect} selected={counterType} />
    <div class="time-container">
      {#if counterType === 'Stopwatch'}
        <Stopwatch />
      {:else if counterType === 'Timer'}
        <p>Timer</p>
      {:else if counterType === 'Manual'}
        <p>Manual</p>
      {/if}
    </div>
  {/if}
  <Calendar />
</main>

<style>
  :global(body) {
    --background-color: #0e1a30;
    --sidebar-width: 200px;
    height: 100vh;
    display: flex;
    flex-flow: row;
    background-color: var(--background-color);
    color: #d3d3ff;
  }
  :global(#app) {
    display: flex;
    flex-flow: column;
    width: 100%;
    height: 100%;
  }
  :global(*:not(dialog)) {
    margin: 0;
    font-family: 'Poppins';
    font-weight: 400;
    box-sizing: border-box;
  }
  :global(button) {
    all: unset;
    cursor: pointer;
  }
  :global(button:focus) {
    outline: none;
  }
  /* TODO - Reduce this to three weights */
  @font-face {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 400;
    src:
      local(''),
      url('assets/fonts/Poppins-Regular.ttf') format('truetype');
  }
  @font-face {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 600;
    src:
      local(''),
      url('assets/fonts/Poppins-SemiBold.ttf') format('truetype');
  }
  @font-face {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 700;
    src:
      local(''),
      url('assets/fonts/Poppins-Bold.ttf') format('truetype');
  }
  @font-face {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 800;
    src:
      local(''),
      url('assets/fonts/Poppins-ExtraBold.ttf') format('truetype');
  }
  @font-face {
    font-family: 'Poppins';
    font-style: normal;
    font-weight: 900;
    src:
      local(''),
      url('assets/fonts/Poppins-Black.ttf') format('truetype');
  }
  main {
    display: flex;
    height: 100%;
    width: 100%;
    flex-flow: column;
    padding: 2rem;
    align-items: center;
    gap: 2rem;
  }
  .isSideBarOpen {
    margin-inline: var(--sidebar-width);
    width: calc(100% - var(--sidebar-width));
  }
  .time-container {
    background-color: #0b1527;
    padding: 2rem;
    display: flex;
    justify-content: center;
    width: 100%;
    border-radius: 1rem;
  }
</style>
