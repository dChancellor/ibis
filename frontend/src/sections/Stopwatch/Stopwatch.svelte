<script lang="ts">
  import Button from '@components/Button.svelte';
  import Ring from '@components/Ring.svelte';
  import Clock from '@components/Clock.svelte';

  let isRunning = false;

  function toggleStopwatch() {
    isRunning = !isRunning;
  }
</script>

<div class="stopwatch">
  <div class="center">
    <Clock {isRunning} />
    <Button size={'small'} variant={'secondary'} onClick={toggleStopwatch}>{isRunning ? 'Pause' : 'Start'}</Button>
  </div>
  <Ring
    type={'running'}
    {isRunning}
    direction={'up'}
    diameter={205}
    duration={3600}
    strokeWidth={25}
    color={'hsl(344, 96%, 52%)'}
  />
  <Ring
    type={'running'}
    {isRunning}
    direction={'up'}
    diameter={165}
    duration={60}
    strokeWidth={15}
    color={'hsl(184, 100%, 50%)'}
  />
  <Ring
    type={'running'}
    {isRunning}
    direction={'up'}
    diameter={140}
    duration={1}
    strokeWidth={10}
    color={'hsl(189, 100%, 90%)'}
  />
</div>

<style>
  .stopwatch {
    position: relative;
    height: 300px;
    width: 300px;
  }
  .center {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 1;
    font-weight: 700;
    display: flex;
    flex-flow: column;
    justify-content: center;
  }
</style>
