<script lang="ts">
  import Button from '@components/Button.svelte';
  import RingV2 from '@components/RingV2.svelte';
  import Clock from '@components/Clock.svelte';

  let time = 0;
  let interval;
  let isRunning = false;

  function toggleStopwatch() {
    isRunning = !isRunning;
  }
  // function toggleStopwatch(): void {
  //   if (isRunning) {
  //     clearInterval(interval);
  //   } else {
  //     interval = setInterval(() => {
  //       time += 10;
  //     }, 10);
  //   }
  //   isRunning = !isRunning;
  // }
  // $: if (!isRunning) clearInterval(interval);
</script>

<div class="stopwatch">
  <div class="center">
    <Clock {isRunning} />
    <Button size={'small'} variant={'secondary'} onClick={toggleStopwatch}>Toggle</Button>
  </div>
  <RingV2 {isRunning} diameter={145} duration={1} strokeWidth={10} color={'#FFF'} />
  <RingV2 {isRunning} diameter={173} duration={60} strokeWidth={15} color={'#36d28e'} />
  <RingV2 {isRunning} diameter={216} duration={3600} strokeWidth={25} color={'#d036d2'} />
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
  }
  /* .rings {
    z-index: -1;
  } */
</style>
