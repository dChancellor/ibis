<script lang="ts">
  /**
   * Time in milliseconds
   */
  let time: number = 0;

  export let isRunning: boolean;

  let interval;

  function toggleStopwatch(): void {
    interval = setInterval(() => {
      time += 1000;
    }, 1000);
  }

  function formatTime(milliseconds): string {
    let date = new Date(milliseconds);
    let hours = String(date.getUTCHours()).padStart(2, '0');
    let minutes = String(date.getUTCMinutes()).padStart(2, '0');
    let seconds = String(date.getUTCSeconds()).padStart(2, '0');
    return `${hours}:${minutes}:${seconds}`;
  }

  $: isRunning ? toggleStopwatch() : clearInterval(interval);
</script>

<div class="clock">
  {formatTime(time)}
</div>

<style>
  .clock {
    text-align: center;
    font-weight: inherit;
  }
</style>
