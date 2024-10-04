<script lang="ts">
  import CalendarRow from './components/CalendarRow.svelte';
</script>

<table class="calendar">
  <thead>
    <tr>
      <th>S</th>
      <th>M</th>
      <th>T</th>
      <th>W</th>
      <th>T</th>
      <th>F</th>
      <th>S</th>
    </tr>
  </thead>
  <tbody>
    <CalendarRow firstDate={1} />
  </tbody>
</table>

<style>
</style>
