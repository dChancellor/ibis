<script lang="ts">
  import Button from "./Button.svelte";
  import Ring from "./Ring.svelte";

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
    let hours = String(date.getUTCHours()).padStart(2, "0");
    let minutes = String(date.getUTCMinutes()).padStart(2, "0");
    let seconds = String(date.getUTCSeconds()).padStart(2, "0");
    return `${hours}:${minutes}:${seconds}`;
  }

  $: if (!isRunning) clearInterval(interval);
</script>

<div class="rings" on:click={toggleStopwatch} on:keypress={toggleStopwatch}>
  <!-- {formatTime(time)} -->
  <!-- <Ring
    color="#710014"
    innerDiameter={350}
    outerDiameter={400}
    type={"timed"}
    animationLength={3600}
  />
  <Ring
    color="#2a0041"
    innerDiameter={300}
    outerDiameter={350}
    type={"timed"}
    animationLength={60}
  /> -->
  <Ring
    color="rgb(255, 255, 0)"
    innerDiameter={220}
    outerDiameter={300}
    type={"timed"}
    animationLength={1}
  />
</div>

<style>
    .rings {
    position: relative;
    height: 250px;
    width: 250px;
  }
</style>
