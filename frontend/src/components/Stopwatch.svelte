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
  <Ring
    startingColor="#fa114f"
    endColor="#f93885"
    innerDiameter={470}
    outerDiameter={600}
    type={"timed"}
    animationLength={3600}
  />
  <Ring
    startingColor="#99ff01"
    endColor="#d8ff00"
    innerDiameter={300}
    outerDiameter={470}
    type={"timed"}
    animationLength={60}
  />
  <Ring
    startingColor="#00d8ff"
    endColor="#02ffa9"
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
