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
  <Ring {isRunning} direction={'up'} diameter={145} duration={1} strokeWidth={10} color={'#FFF'} />
  <Ring {isRunning} direction={'up'} diameter={173} duration={60} strokeWidth={15} color={'hsl(189, 100%, 50%)'} />
  <Ring {isRunning} direction={'up'} diameter={216} duration={3600} strokeWidth={25} color={'hsl(84, 100%, 50%)'} />
  <Ring
    {isRunning}
    direction={'down'}
    oneTime={true}
    diameter={296}
    duration={5}
    strokeWidth={25}
    color={'hsl(84, 100%, 50%)'}
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
