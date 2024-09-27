<script lang="ts">
  import Button from '../../components/Button.svelte';
  import Ring from '../../components/Ring.svelte';
  import RingV2 from '../../components/RingV2.svelte';

  let time = 0;
  let interval;
  let isRunning = false;

  function toggleStopwatch(): void {
    if (isRunning) {
      clearInterval(interval);
    } else {
      interval = setInterval(() => {
        time += 10;
      }, 10);
    }
    isRunning = !isRunning;
  }

  function formatTime(milliseconds): string {
    let date = new Date(milliseconds);
    let hours = String(date.getUTCHours()).padStart(2, '0');
    let minutes = String(date.getUTCMinutes()).padStart(2, '0');
    let seconds = String(date.getUTCSeconds()).padStart(2, '0');
    return `${hours}:${minutes}:${seconds}`;
  }

  $: if (!isRunning) clearInterval(interval);
</script>

<div class="stopwatch" on:click={toggleStopwatch} on:keypress={toggleStopwatch}>
  <!-- {formatTime(time)} -->
  <!-- <Ring
    startingColor="#fa114f"
    endColor="#f93885"
    innerDiameter={250}
    outerDiameter={300}
    type={'timed'}
    animationLength={3600}
  />
  <Ring
    startingColor="#99ff01"
    endColor="#d8ff00"
    innerDiameter={180}
    outerDiameter={250}
    type={'timed'}
    animationLength={60}
  />
  <Ring
    startingColor="#00d8ff"
    endColor="#02ffa9"
    innerDiameter={170}
    outerDiameter={180}
    type={'timed'}
    animationLength={1}
  /> -->
  <!-- <RingV2 diameter={100} duration={60} strokeWidth={20} color={'#00d8ff'} /> -->
  <RingV2 diameter={55} duration={60} strokeWidth={20} color={'#FFF'} />
  <RingV2 diameter={105} duration={60} strokeWidth={10} color={'#FFF'} />
  <RingV2 diameter={155} duration={60} strokeWidth={30} color={'#FFF'} />
</div>

<style>
  .stopwatch {
    position: relative;
    height: 300px;
    width: 300px;
  }
</style>
